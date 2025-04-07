import { useRef, useState, useEffect } from 'react'
import { styled } from '@mui/material/styles'
import {
  TextField,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  TablePagination,
  CircularProgress,
  Paper,
} from '@mui/material'
import Icon from '@/components/icon'
import MyButton from '@/components/button'
import Modal from '@/components/modal'
import MyForm from '@/components/form'
import { postSketchUpload } from '@/api/sketch'
import { getProjectList } from '@/api/project'
import VisuallyHiddenInput from '@/components/visuallyHiddenInput'
import Loading from '@/components/loading'
import { useNavigate } from 'react-router-dom'
import { Project } from '@/types'
import dayjs from 'dayjs'

const Home = () => {
  const fileInputRef = useRef<HTMLInputElement>(null)
  const [searchTerm, setSearchTerm] = useState<string>('')
  const [projects, setProjects] = useState<Project[]>([])
  const [isModalOpen, setIsModalOpen] = useState(false)
  const [formData, setFormData] = useState({
    project: '',
    comment: '',
    address: '',
    attachments: [] as { file: File; preview: string }[],
  })

  const [page, setPage] = useState(0)
  const [perPage, setPerPage] = useState(10)
  const [totalCount, setTotalCount] = useState(0)
  const [loading, setLoading] = useState(true)

  const navigate = useNavigate()

  const loadData = async () => {
    setLoading(true)
    const payload = {
      page_number: page + 1,
      page_size: perPage,
    }
    const { data: resp } = await getProjectList({ params: payload })
    console.log('--------resp-----', resp)

    setTotalCount(resp.data.total)

    resp.data.projects.map((i: Project) => {
      i.updated_at = dayjs(i.updated_at).format('MM-DD-YYYY')
      return i
    })
    setProjects(resp.data.projects)
    setLoading(false)
  }

  useEffect(() => {
    loadData()
  }, [page, perPage])

  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage)
  }

  const handleChangeRowsPerPage = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setPerPage(parseInt(event.target.value, 10))
    setPage(0)
  }

  const handleOpenModal = () => setIsModalOpen(true)
  const handleCloseModal = () => setIsModalOpen(false)
  const handleSubmit = async () => {
    // Add your submission logic here
    const form = new FormData()

    form.append('project', formData.project)
    form.append('comment', formData.comment)
    form.append('address', formData.address)

    // Append files (assuming `attachments` is an array of File objects)
    formData.attachments.forEach((attachment) => {
      form.append('files', attachment.file) // Key must match Gin's FormFile name
    })

    const { data: resp } = await postSketchUpload(form)
    handleCloseModal()
    loadData()
  }

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }))
  }

  const handleFileUpload = (e: React.ChangeEvent<HTMLInputElement>) => {
    console.log('handleFileUpload', e.target.files)
    if (!e.target.files || e.target.files.length === 0) return

    const files = Array.from(e.target.files)

    const newAttachments = files.map((file) => ({
      file,
      preview: URL.createObjectURL(file),
    }))

    setFormData((prev) => ({
      ...prev,
      attachments: [...prev.attachments, ...newAttachments],
    }))
  }

  const handleRemoveFile = (index: number) => {
    setFormData((prev) => {
      const updatedAttachments = [...prev.attachments]
      updatedAttachments.splice(index, 1)
      return {
        ...prev,
        attachments: updatedAttachments,
      }
    })
  }

  const columns: { key: string; label: string }[] = [
    { key: 'project_name', label: 'Project Name' },
    { key: 'action', label: 'Action' },
    { key: 'updated_at', label: 'Updated At' },
    { key: 'comment', label: 'Comment' },
    { key: 'status', label: 'Status' },
  ]

  const handleDetailsButtonClick = (projectId: number) => {
    console.log('Project ID:', projectId)
    navigate(`/project/details?id=${projectId}`)
  }

  return (
    <div className="p-4 space-y-4">
      <div className="flex items-center space-x-2">
        <TextField
          size="small"
          label="Search by project name"
          variant="outlined"
          value={searchTerm}
          className=" flex-1 max-w-xs"
          onChange={(e) => setSearchTerm(e.target.value)}
        />
        <MyButton variant="outlined" className="h-[40px]">
          <Icon name="search" />
        </MyButton>
        <MyButton
          variant="outlined"
          className="h-[40px]"
          onClick={handleOpenModal}
        >
          <Icon name="add" />
        </MyButton>
        <Modal
          open={isModalOpen}
          onClose={handleCloseModal}
          onSubmit={handleSubmit}
          title="Upload New Sketch"
        >
          <MyForm
            formData={formData}
            setFormData={setFormData}
            onSubmit={handleSubmit}
            className="space-y-4"
          >
            <TextField
              fullWidth
              label="project name"
              name="project"
              required
              value={formData.project}
              onChange={handleInputChange}
              variant="outlined"
              className="mb-4"
            />
            <TextField
              fullWidth
              label="project address"
              name="address"
              value={formData.address}
              onChange={handleInputChange}
              variant="outlined"
            />
            <TextField
              fullWidth
              label="comment"
              name="comment"
              value={formData.comment}
              onChange={handleInputChange}
              variant="outlined"
            />
            <input
              type="file"
              ref={fileInputRef}
              onChange={handleFileUpload}
              multiple
              accept="image/*"
              className="hidden"
            />
            <MyButton variant="contained" color="secondary" component="label">
              <Icon name="upload" />
              <VisuallyHiddenInput
                type="file"
                onChange={handleFileUpload}
                multiple
              />
            </MyButton>

            {/* Thumbnail display */}
            {formData.attachments.length > 0 && (
              <div className="mt-3">
                <p className="text-sm font-medium text-gray-700 mb-2">
                  Attachments:
                </p>
                <div className="flex flex-wrap gap-3">
                  {formData.attachments.map((attachment, index) => (
                    <div key={index} className="relative group">
                      <div className="w-24 h-24 rounded-md overflow-hidden border border-gray-200">
                        <img
                          src={attachment.preview}
                          alt={`Preview ${attachment.file.name}`}
                          className="w-full h-full object-cover"
                        />
                      </div>
                      <MyButton
                        size="small"
                        className="text-red-500"
                        onClick={() => handleRemoveFile(index)}
                      >
                        <Icon name="delete" />
                      </MyButton>
                      <p className="text-xs text-gray-500 mt-1 truncate w-24">
                        {attachment.file.name}
                      </p>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </MyForm>
        </Modal>
      </div>

      {/* 列表 */}
      <div className="max-w-full overflow-x-auto border rounded-lg">
        <TableContainer component={Paper} className="max-w-full">
          {loading ? (
            <Loading message="Loading projects..." />
          ) : (
            <>
              <Table aria-label="projects table" className="min-w-max">
                <TableHead>
                  <TableRow className="bg-gray-50">
                    {columns.map((column) => (
                      <TableCell
                        key={column.key}
                        className={`font-semibold text-gray-700 ${
                          column.key === 'project_name'
                            ? 'sticky left-0 z-10 bg-gray-50'
                            : ''
                        }`}
                      >
                        {column.label}
                      </TableCell>
                    ))}
                  </TableRow>
                </TableHead>
                <TableBody>
                  {projects.map((project) => (
                    <TableRow
                      key={project.id}
                      hover
                      className="hover:bg-gray-50"
                    >
                      {columns.map((column) => (
                        <TableCell
                          key={column.key}
                          className={`${
                            column.key === 'project_name'
                              ? 'sticky left-0 z-10 bg-white font-medium text-gray-900'
                              : 'text-gray-600'
                          }`}
                        >
                          {column.key === 'action' ? (
                            <MyButton
                              onClick={() =>
                                handleDetailsButtonClick(project.id)
                              }
                            >
                              Details
                            </MyButton>
                          ) : (
                            project[column.key as keyof Project]
                          )}
                        </TableCell>
                      ))}
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </>
          )}
        </TableContainer>
        <TablePagination
          rowsPerPageOptions={[5, 10, 25]}
          component="div"
          count={totalCount}
          rowsPerPage={perPage}
          page={page}
          onPageChange={handleChangePage}
          onRowsPerPageChange={handleChangeRowsPerPage}
          className="border-t border-gray-200"
        />
      </div>
    </div>
  )
}
export default Home

/*
 * @fileName details.tsx
 * @author Di Sheng
 * @date 2025/04/06 14:05:26
 * @description Project details with sketch images
 */
import MyBreadcrumbs from '@/components/breadcrumb'
import { useSearchParams } from 'react-router-dom'
import { useRef, useState, useEffect } from 'react'
import Loading from '@/components/loading'
import { getProjectDetails } from '@/api/project'
import { ProjectDetails } from '@/types'
import dayjs from 'dayjs'
import { useTheme, useMediaQuery, Dialog } from '@mui/material'
import Icon from '@/components/icon'

import {
  Paper,
  Typography,
  CircularProgress,
  Grid2,
  Grid,
  Card,
  CardMedia,
  CardContent,
} from '@mui/material'

const Details = () => {
  const [searchParams] = useSearchParams()
  const projectId = searchParams.get('id')

  const [previewImage, setPreviewImage] = useState<string | null>(null)
  const theme = useTheme()
  const isMdUp = useMediaQuery(theme.breakpoints.up('md')) // md and up

  const [projectDetail, setProjectDetail] = useState<ProjectDetails | null>(
    null
  )
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchProject = async () => {
      setLoading(true)
      console.log('--------id---------', projectId)
      const { data: response } = await getProjectDetails(projectId)
      console.log('Project details:', response)
      setProjectDetail(response.data)
      setLoading(false)
    }
    fetchProject()
  }, [projectId])
  return (
    <>
      <MyBreadcrumbs />
      <div className="p-4">
        {loading ? (
          <Loading message="Fetching project details..." />
        ) : projectDetail ? (
          <div>
            <div className="p-6 max-w-6xl mx-auto">
              <Typography variant="h4" className="mb-4">
                Project: {projectDetail.project.project_name}
              </Typography>

              <Paper className="p-4 mb-6">
                <Typography variant="body1">
                  <strong>Status:</strong> {projectDetail.project.status}
                </Typography>
                <Typography variant="body1">
                  <strong>Address:</strong>{' '}
                  {projectDetail.project.address || 'No address provided'}
                </Typography>
                <Typography variant="body1">
                  <strong>Comment:</strong>{' '}
                  {projectDetail.project.comment || 'No comment'}
                </Typography>
                <Typography variant="body1">
                  <strong>Updated At:</strong>{' '}
                  {new Date(projectDetail.project.updated_at).toLocaleString()}
                </Typography>
              </Paper>

              <Typography variant="h5" className="mb-4">
                Sketches
              </Typography>

              <Grid2 container spacing={3}>
                {projectDetail.sketches.map((sketch) => (
                  <Grid item xs={12} sm={6} md={4} key={sketch.id}>
                    <Card className="rounded-2xl shadow-md">
                      <CardMedia
                        component="img"
                        height="200"
                        image={sketch.full_image_url}
                        alt={`Sketch ${sketch.id}`}
                        onClick={() =>
                          isMdUp && setPreviewImage(sketch.full_image_url)
                        }
                        className={
                          isMdUp
                            ? 'cursor-pointer hover:opacity-80 transition'
                            : ''
                        }
                      />
                      <CardContent>
                        <Typography variant="body2" className="text-gray-700">
                          Last Update:{' '}
                          {dayjs(sketch.updated_at).format('MM-DD-YYYY HH:mm')}
                        </Typography>
                      </CardContent>
                    </Card>
                  </Grid>
                ))}
              </Grid2>
            </div>
            <Dialog
              open={!!previewImage}
              onClose={() => setPreviewImage(null)}
              maxWidth="md"
            >
              <div className="relative">
                {/* Close Button */}
                <div onClick={() => setPreviewImage(null)}>
                  <Icon
                    name="close"
                    className=" w-6 h-6 absolute top-4 right-4 z-10 text-white bg-black/50 hover:bg-black/70"
                  />
                </div>

                {/* Image Preview */}
                <img
                  src={previewImage ?? ''}
                  alt="Preview"
                  className="w-full h-auto rounded"
                />
              </div>
            </Dialog>
          </div>
        ) : (
          <div>Project not found</div>
        )}
      </div>
    </>
  )
}

export default Details

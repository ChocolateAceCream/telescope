import { useState, useRef } from 'react'
import {
  Button,
  Container,
  Box,
  Typography,
  LinearProgress,
  Snackbar,
  Alert,
} from '@mui/material'
import { PDFDocument, rgb } from 'pdf-lib'
import Icon from '@/components/icon'

type FileState = File | null
type ProcessedPdfState = string | null

export default function PDFTextRemover() {
  const [file, setFile] = useState<FileState>(null)
  const [processedPdf, setProcessedPdf] = useState<ProcessedPdfState>(null)
  const [loading, setLoading] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null)
  const [success, setSuccess] = useState<string | null>(null)
  const fileInputRef = useRef<HTMLInputElement>(null)

  const [isDragging, setIsDragging] = useState(false)

  const handleDrop = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault()
    setIsDragging(false)
    const droppedFile = e.dataTransfer.files?.[0]
    if (droppedFile && droppedFile.type === 'application/pdf') {
      setFile(droppedFile)
      setProcessedPdf(null)
      setError(null)
    } else {
      setError('Please drop a valid PDF file')
    }
  }

  const handleDragOver = (e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault()
    setIsDragging(true)
  }

  const handleDragLeave = () => {
    setIsDragging(false)
  }

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = e.target.files?.[0]
    if (selectedFile && selectedFile.type === 'application/pdf') {
      setFile(selectedFile)
      setProcessedPdf(null)
      setError(null)
      // 👇 reset input so same file can be reselected later
      e.target.value = ''
    } else {
      setError('Please select a PDF file')
    }
  }

  const processPDF = async () => {
    if (!file) {
      setError('Please select a PDF file first')
      return
    }

    setLoading(true)
    setError(null)
    setSuccess(null)

    try {
      // Read the PDF file
      const arrayBuffer = await file.arrayBuffer()
      const pdfDoc = await PDFDocument.load(arrayBuffer)
      const pages = pdfDoc.getPages()

      // Process first page
      let { width, height } = pages[0].getSize()
      console.log('width', width, 'height', height)
      // Draw a white rectangle over the text
      // Adjust these coordinates based on where the text appears in your PDF
      pages[0].drawRectangle({
        x: 225,
        y: height - 62,
        width: 200,
        height: 44,
        color: rgb(1, 1, 1),
        opacity: 1,
        borderWidth: 0,
      })

      // Process last page
      const l = pages.length
      if (l > 1) {
        const { width, height } = pages[l - 1].getSize()
        console.log('width', width, 'height', height)
        // Draw a white rectangle over the text
        // Adjust these coordinates based on where the text appears in your PDF
        pages[l - 1].drawRectangle({
          x: 225,
          y: height - 75,
          width: 200,
          height: 44,
          color: rgb(1, 1, 1),
          opacity: 1,
          borderWidth: 0,
        })
      }

      // Save the modified PDF
      const modifiedPdfBytes = await pdfDoc.save()
      const blob = new Blob([modifiedPdfBytes], { type: 'application/pdf' })

      setProcessedPdf(URL.createObjectURL(blob))
      setSuccess('PDF processed successfully!')
    } catch (err) {
      setError(
        `Failed to process PDF: ${
          err instanceof Error ? err.message : 'Unknown error'
        }`
      )
      console.error(err)
    } finally {
      setLoading(false)
    }
  }

  const handleCloseSnackbar = () => {
    setError(null)
    setSuccess(null)
  }

  return (
    <Container maxWidth="md" className="py-8">
      <Box className="text-center mb-8">
        <Typography variant="h4" component="h1" className="mb-2">
          PDF Text Remover
        </Typography>
        <Typography variant="body1">
          Remove specific text from your PDF files
        </Typography>
      </Box>

      <Box className="bg-white rounded-lg shadow-md p-6">
        <Box
          onDrop={handleDrop}
          onDragOver={handleDragOver}
          onDragLeave={handleDragLeave}
          className={`border-2 border-dashed rounded-lg p-6 w-full text-center transition-colors ${
            isDragging ? 'border-blue-500 bg-blue-50' : 'border-gray-300'
          }`}
        >
          <Typography variant="body2" className="text-gray-600 mb-2">
            Drag and drop your PDF file here
          </Typography>
          <Typography variant="body2" className="text-gray-500">
            or use the button below to select manually
          </Typography>
        </Box>
        <input
          type="file"
          accept=".pdf"
          onChange={handleFileChange}
          ref={fileInputRef}
          className="hidden"
        />

        <Box className="flex flex-col items-center gap-4 mb-6 mt-4">
          <Button
            variant="contained"
            color="primary"
            onClick={() => fileInputRef.current?.click()}
            className="w-full sm:w-auto"
          >
            <Icon name="upload" />
          </Button>

          {file && (
            <Typography variant="body2" className="text-gray-600">
              Selected: {file.name}
            </Typography>
          )}
        </Box>

        {loading && <LinearProgress className="mb-4" />}

        <Box className="flex justify-center gap-4">
          <Button
            variant="contained"
            color="secondary"
            onClick={processPDF}
            disabled={!file || loading}
            className="w-full sm:w-auto"
          >
            Process PDF
          </Button>

          {processedPdf && (
            <Button
              variant="contained"
              color="success"
              href={processedPdf}
              download={
                file
                  ? `${file.name.replace(/\.pdf$/i, '')}.pdf`
                  : 'processed.pdf'
              }
              className="w-full sm:w-auto"
            >
              Download Processed PDF
            </Button>
          )}
        </Box>
      </Box>

      <Snackbar
        open={!!error}
        autoHideDuration={6000}
        onClose={handleCloseSnackbar}
      >
        <Alert severity="error" onClose={handleCloseSnackbar}>
          {error}
        </Alert>
      </Snackbar>

      <Snackbar
        open={!!success}
        autoHideDuration={6000}
        onClose={handleCloseSnackbar}
      >
        <Alert severity="success" onClose={handleCloseSnackbar}>
          {success}
        </Alert>
      </Snackbar>
    </Container>
  )
}

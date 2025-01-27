import React from 'react'
import ReactDOM from 'react-dom/client'
import './assets/styles/application.css'
import router from '@/router'
import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import { createTheme, ThemeProvider } from '@mui/material/styles'

const rootElement = document.getElementById('root') as HTMLElement

const theme = createTheme({
  typography: {
    fontFamily: 'Chewy, Inter, sans-serif', // Set your default font family
  },
  palette: {
    primary: {
      main: '#ffb74d', // Change this to your desired primary color (e.g., Orange 500 from Material UI color palette)
      contrastText: '#ffffff', // Optional: Change the text color to ensure good contrast
    },
    secondary: {
      main: '#ff9800',
    },
    warning: {
      main: '#F0908A',
    },
  },
})

ReactDOM.createRoot(rootElement).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <RouterProvider router={router} />
    </ThemeProvider>
  </React.StrictMode>
)

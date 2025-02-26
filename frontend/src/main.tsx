/*
 * @fileName main.tsx
 * @author Di Sheng
 * @date 2025/02/20 17:00:43
 * @description Description:  App initialization (Mount React, Theme, Zustand, etc.)
 */
import React from 'react'
import ReactDOM from 'react-dom/client'
import './assets/styles/application.css' // Global styles
import App from './App' // ✅ Import App.tsx
import { createTheme, ThemeProvider } from '@mui/material/styles'

const rootElement = document.getElementById('root') as HTMLElement

const theme = createTheme({
  typography: {
    fontFamily: 'Chewy, Inter, sans-serif',
  },
  palette: {
    primary: { main: '#ffb74d', contrastText: '#ffffff' },
    secondary: { main: '#ff9800' },
    warning: { main: '#F0908A' },
  },
})

ReactDOM.createRoot(rootElement).render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <App /> {/* ✅ App.tsx handles UI & routing */}
    </ThemeProvider>
  </React.StrictMode>
)

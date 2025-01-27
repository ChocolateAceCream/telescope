import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Home from '@/views/home'
import Login from '@/views/auth/login'
import OAuthResult from '@/views/auth/oauthResult'
import Layout from '@/layout/baseLayout'
import { SSEProvider } from '@/hooks/useSSE'

const ProtectedRoute = () => {
  return (
    // <SSEProvider url="/backend/api/sse/subscribe">
    <Layout></Layout>
    // </SSEProvider>
  )
}
const router = createBrowserRouter([
  {
    path: '/',
    element: <ProtectedRoute />,
    children: [
      { path: '/', element: <Home /> },
      // { path: 'about', element: <About /> },
      // other routes...
    ],
  },
  {
    path: '/login',
    element: <Login />,
  },
  {
    path: '/oauth-success',
    element: <OAuthResult />,
  },
])

export default router

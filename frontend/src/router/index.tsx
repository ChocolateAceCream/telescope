import {
  createBrowserRouter,
  RouterProvider,
  useNavigate,
} from 'react-router-dom'
import Home from '@/views/home'
import Login from '@/views/auth/login'
import OAuthResult from '@/views/auth/oauthResult'
import Layout from '@/layout/baseLayout'
import { SSEProvider } from '@/hooks/useSSE'
import userStore from '@/store/user'
import { useEffect } from 'react'
import NotFound from '@/views/NotFound'

const ProtectedRoute = () => {
  return (
    // <SSEProvider url="/backend/api/sse/subscribe">
    <Layout></Layout>
    // </SSEProvider>
  )
}

const AuthGuard = ({ children }: { children: React.ReactNode }) => {
  const navigate = useNavigate()
  const isAuthed = userStore.getState().user.isAuthed
  console.log('isAuthed:', isAuthed)
  useEffect(() => {
    if (!isAuthed) {
      navigate('/login', { replace: true }) // Use replace to avoid back navigation
    }
  }, [isAuthed, navigate])
  return children
}

const router = createBrowserRouter([
  {
    path: '/',
    element: (
      <AuthGuard>
        <ProtectedRoute />
      </AuthGuard>
    ),
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
  { path: '*', element: <NotFound /> }, // Catch-all route for 404 pages
])

export default router

import { useState } from 'react'
// import { useHistory } from 'react-router-dom';
import userStore from '@/store/user'
import { useNavigate } from 'react-router-dom'

interface UseAuth {
  isAuthenticated: boolean
  login: () => void
  logout: () => void
}

const useAuth = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  // const history = useHistory();
  const logoutUser = userStore((state) => state.logout)
  // const login = () => {
  //   // Logic for logging in (e.g., API call, setting tokens)
  //   userStore.setIsAuthenticated(true)
  // }
  const navigate = useNavigate()

  const logout = () => {
    // Logic for logging out (e.g., clearing tokens, redirecting)
    logoutUser()
    setIsAuthenticated(false)
    navigate('/login') // Redirect to login page after logout
  }

  return {
    isAuthenticated,
    // login,
    logout,
  }
}

export default useAuth

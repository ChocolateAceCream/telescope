import { useState } from 'react'
// import { useHistory } from 'react-router-dom';

const useAuth = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  // const history = useHistory();

  const login = () => {
    // Logic for logging in (e.g., API call, setting tokens)
    setIsAuthenticated(true)
  }

  const logout = () => {
    // Logic for logging out (e.g., clearing tokens, redirecting)
    setIsAuthenticated(false)
    // history.push('/login'); // Redirect to login page after logout
  }

  return {
    isAuthenticated,
    login,
    logout,
  }
}

export default useAuth

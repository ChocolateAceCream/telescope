/*
* @fileName ase.jsx
* @author Di Sheng
* @date 2024/08/11 21:01:33
* @description: global notification component

example usage:
const handleClick = () => {
  showNotification({
    type: 'success',
    message: 'This is a success notification!',
    duration: 3000, // Duration is in milliseconds
  });
};

*/
import ReactDOM from 'react-dom/client'
import Alert from '@mui/material/Alert'
import Snackbar from '@mui/material/Snackbar'

type MessageType = 'success' | 'info' | 'warning' | 'error'

interface NotificationProps {
  type: MessageType
  message: string
  duration: number
}

const showNotification = ({
  type,
  message,
  duration = 3000,
}: NotificationProps) => {
  const container = document.createElement('div')
  document.body.appendChild(container)

  const root = ReactDOM.createRoot(container)

  root.render(
    <Snackbar
      anchorOrigin={{ vertical: 'top', horizontal: 'right' }}
      open={true}
      autoHideDuration={duration}
      onClose={() => {
        root.unmount()
        document.body.removeChild(container)
      }}
      sx={{ position: 'fixed', zIndex: 1500 }}
    >
      <Alert severity={type} sx={{ width: '100%' }}>
        {message}
      </Alert>
    </Snackbar>,
  )
}

export default showNotification

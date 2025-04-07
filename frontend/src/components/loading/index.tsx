/*
 * @fileName index.tsx
 * @author Di Sheng
 * @date 2025/04/06 16:49:52
 * @description loading component
 */

// components/Loading.tsx
import CircularProgress from '@mui/material/CircularProgress'

const Loading = ({ message = 'Loading...' }: { message?: string }) => {
  return (
    <div className="flex flex-col items-center justify-center p-6 text-gray-500">
      <CircularProgress />
      <span className="mt-2 text-sm">{message}</span>
    </div>
  )
}

export default Loading

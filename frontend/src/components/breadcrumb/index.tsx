/*
 * @fileName index.tsx
 * @author Di Sheng
 * @date 2025/04/06 14:28:06
 * @description Breadcrumb component
 */
// components/Breadcrumb.tsx
import { Breadcrumbs, Link, Typography } from '@mui/material'
import { useLocation, Link as RouterLink } from 'react-router-dom'

const MyBreadcrumbs = () => {
  const location = useLocation()

  const pathnames = location.pathname.split('/').filter((x) => x)

  return (
    <div className="p-2">
      <Breadcrumbs aria-label="breadcrumb">
        {pathnames.map((value, index) => {
          const to = `/${pathnames.slice(0, index + 1).join('/')}`
          const isLast = index === pathnames.length - 1
          return isLast ? (
            <Typography color="text.primary" key={to} className="text-sm">
              {decodeURIComponent(value)}
            </Typography>
          ) : (
            <Link
              underline="hover"
              color="inherit"
              to={to}
              component={RouterLink}
              key={to}
              className="text-sm"
            >
              {decodeURIComponent(value)}
            </Link>
          )
        })}
      </Breadcrumbs>
    </div>
  )
}

export default MyBreadcrumbs

/*
 * @fileName index.tsx
 * @author Di Sheng
 * @date 2024/09/09 08:28:41
 * @description Customized button component
 */
import Button from '@mui/material/Button'
import { styled } from '@mui/material/styles'

const MyButton = styled(Button)(({ theme }) => ({
  backgroundColor: theme.palette.primary.main, // Use the primary color from the theme
  color: theme.palette.primary.contrastText, // Ensure the text color has good contrast
  '&:hover': {
    backgroundColor: theme.palette.secondary.main, // Use the secondary color on hover
  },
})) as typeof Button

export default MyButton

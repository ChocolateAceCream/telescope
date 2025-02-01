/*
 * @fileName index.tsx
 * @author Di Sheng
 * @date 2024/08/14 11:32:56
 * @description Description: base layout
 */
import { Outlet } from 'react-router-dom'
import Header from './header'
const BaseLayout = () => {
  return (
    <>
      <Header />
      <main>
        <Outlet></Outlet>
      </main>
      <div>footer</div>
    </>
  )
}

export default BaseLayout

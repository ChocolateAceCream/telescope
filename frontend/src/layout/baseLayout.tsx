/*
 * @fileName index.tsx
 * @author Di Sheng
 * @date 2024/08/14 11:32:56
 * @description Description: base layout
 */
import { Outlet } from 'react-router-dom'
const BaseLayout = () => {
  return (
    <>
      <div>header</div>
      <div>navbar</div>
      <main>
        <Outlet></Outlet>
      </main>
      <div>footer</div>
    </>
  )
}

export default BaseLayout

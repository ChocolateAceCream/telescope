/*
 * @fileName index.tsx
 * @author Di Sheng
 * @date 2024/08/14 11:32:56
 * @description Base layout with a fixed header and scrollable main content.
 */
import { Outlet } from 'react-router-dom'
import Header from './header'

const BaseLayout = () => {
  return (
    <div className="flex flex-col h-screen">
      {/* 🔹 Fixed Header */}
      <header className="fixed top-0 left-0 w-full z-50 bg-white shadow-md">
        <Header />
      </header>

      {/* 🔹 Scrollable Main Content */}
      <main className="flex-1 mt-[4rem] overflow-auto p-4">
        <Outlet />
      </main>

      {/* 🔹 Footer (Optional) */}
      {/* <footer className="py-4 text-center text-gray-600">Footer Content</footer> */}
    </div>
  )
}

export default BaseLayout

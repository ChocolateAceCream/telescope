/*
 * @fileName header.tsx
 * @author Di Sheng
 * @date 2025/01/31 11:33:37
 * @description header
 */
import React, { useState } from 'react'
import { Link } from 'react-router-dom' // If using React Router
import logo from '@/assets/image/logo.png'
import Icon from '@/components/icon'
const Header = () => {
  const [isOpen, setIsOpen] = useState(false)
  const menus = [
    {
      name: 'Home',
      url: '/home',
    },
    {
      name: 'About',
      url: '/about',
    },
    {
      name: 'Services',
      url: '/services',
    },
    {
      name: 'Contact',
      url: '/contact',
    },
  ]
  return (
    <div className="border-gray-200 bg-gray-50 dark:bg-gray-800 dark:border-gray-700 flex items-center w-full px-4">
      {/* Mobile Menu Button */}
      <button className="md:hidden" onClick={() => setIsOpen(!isOpen)}>
        <Icon name="menu" className="menu" />
      </button>

      {/* Mobile Dropdown Menu */}
      {isOpen && (
        <div
          className={`z-50 fixed inset-0  flex justify-center
          transform ${isOpen ? 'translate-x-0' : 'translate-x-full'}
          transition-transform duration-300 ease-in-out`}
        >
          {/* 🔹 Close Button */}
          <button
            className="absolute top-5 right-5 text-white text-3xl focus:outline-none"
            onClick={() => setIsOpen(false)}
          >
            <Icon name="close" />
          </button>

          {/* 🔹 Menu Items */}
          <nav className=" text-2xl bg-white space-y-6 text-center mt-16 w-full bg-opacity-100 px-8">
            {menus.map((menu) => (
              <Link
                to={menu.url}
                key={menu.name}
                className="font-chewy block hover:text-gray-400 border-b-2 border-shadow h-16 flex items-center justify-center relative px-4"
                onClick={() => setIsOpen(false)}
              >
                {menu.name}
                <Icon name="rightArrow" className="absolute right-0" />
              </Link>
            ))}
          </nav>
        </div>
      )}
      {/* Logo */}
      <h1 className="text-2xl font-bold">
        <Link to="/">
          <img src={logo} alt="logo" className="h-[4rem] w-auto" />
        </Link>
      </h1>

      {/* Desktop Navigation */}
      <nav className="hidden md:flex space-x-12">
        {menus.map((menu) => (
          <Link
            to={menu.url}
            key={menu.name}
            className="font-chewy hover:text-concrete"
          >
            {menu.name}
          </Link>
        ))}
      </nav>
    </div>
  )
}

export default Header

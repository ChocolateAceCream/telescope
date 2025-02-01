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

  return (
    <div className="border-gray-200 bg-gray-50 dark:bg-gray-800 dark:border-gray-700 flex items-center w-full p-4">
      {/* Logo */}
      <h1 className="text-2xl font-bold">
        <Link to="/">
          <img src={logo} alt="logo" className="h-[70px] w-auto" />
        </Link>
      </h1>

      {/* Desktop Navigation */}
      <nav className="md:hidden space-x-6">
        <Link to="/home" className="hover:text-gray-200">
          Home
        </Link>
        <Link to="/about" className="hover:text-gray-200">
          About
        </Link>
        <Link to="/services" className="hover:text-gray-200">
          Services
        </Link>
        <Link to="/contact" className="hover:text-gray-200">
          Contact
        </Link>
      </nav>

      {/* Mobile Menu Button */}
      <button className="md:hidden" onClick={() => setIsOpen(!isOpen)}>
        {isOpen ? (
          <Icon name="close" className="close" />
        ) : (
          <Icon name="menu" className="menu" />
        )}
      </button>

      {/* Mobile Dropdown Menu */}
      {isOpen && (
        <nav className="md:hidden bg-blue-700 mt-2 p-4 space-y-2">
          <Link to="/home" className="block hover:text-gray-200">
            Home
          </Link>
          <Link to="/about" className="block hover:text-gray-200">
            About
          </Link>
          <Link to="/services" className="block hover:text-gray-200">
            Services
          </Link>
          <Link to="/contact" className="block hover:text-gray-200">
            Contact
          </Link>
        </nav>
      )}
    </div>
  )
}

export default Header

/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './src/App.{js,jsx,ts,tsx}',
    './src/**/*.{js,jsx,ts,tsx}',
    './src/**/**/*.{js,jsx,ts,tsx}',
  ],
  theme: {
    extend: {
      colors: {
        'lite-purple': '#f8e9e6',
        'warning': '#F0908A',
        primary: '#ffb74d',
        secondary: '#ff9800',
        'lite-blue': '#C1E1E7',
        shadow: '#919191',
        'lite-pink': '#ea8685',
        'lite-orange': '#FEE5CB',

      },
      height: {
        'screen-dynamic': 'calc(var(--vh, 1vh) * 100)',
      },
      fontFamily: {
        sans: [
          '"Chewy"',
          '"Baloo 2"',
          '"Inter"',
          '"Noto Sans"',
          'system-ui',
          'Segoe UI',
          'Roboto',
          'Helvetica Neue',
          'Arial',
          'sans-serif',
          'Apple Color Emoji',
          'Segoe UI Emoji',
          'Segoe UI Symbol',
        ],
        baloo: ['"Baloo 2"', 'sans-serif'], // Register the font separately for custom use
        chewy: ['"Chewy"', 'sans-serif'], // Register the font separately for custom use
      },
    },
  },
  plugins: [
    function ({ addUtilities, theme }) {
      const newUtilities = {
        '.text-shadow-sm': {
          textShadow: '0 1px 2px rgba(0, 0, 0, 0.05)',
        },
        '.text-shadow-md': {
          textShadow: '0 2px 4px rgba(0, 0, 0, 0.1)',
        },
        '.text-shadow-lg': {
          textShadow: '0 3px 6px rgba(0, 0, 0, 0.15)',
        },
        '.text-shadow-xl': {
          textShadow: `0 4px 8px ${theme('colors.shadow')}`,
        },
        '.text-shadow-2xl': {
          textShadow: '0 5px 10px rgba(0, 0, 0, 0.25)',
        },
        '.text-shadow-none': {
          textShadow: 'none',
        },
      }

      addUtilities(newUtilities, ['responsive', 'hover'])
    },
  ],
}

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{html,js,vue}'],
  theme: {
      extend: {
          boxShadow: {
              mild: 'rgba(149, 157, 165, 0.2) 0px 8px 24px;',
              image: 'rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;',
          },
      },
  },
  plugins: [],
};

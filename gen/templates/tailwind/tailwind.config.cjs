/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      fontSize: {
        title: '2rem'
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms')
  ],
}

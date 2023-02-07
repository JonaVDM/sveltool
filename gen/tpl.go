package gen

func tailwindConfigTemplate() []byte {
	return []byte(`/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {},
  },
  plugins: [],
}`)
}

func tailwindPostcssConfigTemplate() []byte {
	return []byte(`module.exports = {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
}`)
}

func tailwindStylesTemplate() []byte {
	return []byte(`@tailwind base;
@tailwind components;
@tailwind utilities;`)
}

func tailwindLayoutTemplate() []byte {
	return []byte(`<script>
	import '../app.css';
</script>

<slot />`)
}

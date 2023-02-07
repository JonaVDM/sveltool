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

func picoCssLayoutTemplate() []byte {
	return []byte(`<script>
	import '@picocss/pico';
</script>

<slot />`)
}

func pocketbaseLib() []byte {
	return []byte(`import PocketBase from 'pocketbase'
import { writable } from 'svelte/store';

export const pb = new PocketBase('http://127.0.0.1:8090');

export const currentUser = writable(pb.authStore.model);

pb.authStore.onChange(() => {
  currentUser.set(pb.authStore.model);
});`)
}

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

func clientHooksTemplate() []byte {
	return []byte(`import { pb } from '$lib/pocketbase'

pb.authStore.loadFromCookie(document.cookie);`)
}

func serverHooksTemplate() []byte {
	return []byte(`import { pb } from '$lib/pocketbase';
import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
  pb.authStore.loadFromCookie(event.request.headers.get('cookie') || '');
  if (pb.authStore.isValid) {
    try {
      await pb.collection('users').authRefresh();
    } catch (_) {
      pb.authStore.clear();
    }
  }

  event.locals.pb = pb;
  event.locals.user = structuredClone(pb.authStore.model);

  const response = await resolve(event);

  if (event.locals.pb.authStore.isValid) {
    response.headers.set(
      'set-cookie',
      pb.authStore.exportToCookie({ httpOnly: false })
    );
  }

  return response;
};`)
}

func pbAuthLayoutTemplate() []byte {
	return []byte(`import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = ({ locals }) => {
  if (locals.pb.authStore.isValid) {
    throw redirect(303, "/")
  }
}`)
}

func pbLoginServerTemplate() []byte {
	return []byte(`import { fail, redirect } from '@sveltejs/kit';
import type { ClientResponseError } from 'pocketbase';
import type { Actions } from './$types';

export const actions: Actions = {
  default: async ({ locals, request }) => {
    const { email, password } = Object.fromEntries(await request.formData()) as {
      email: string;
      password: string;
    }

    try {
      await locals.pb
        .collection('users')
        .authWithPassword(email, password);
    } catch (e) {
      const err = e as ClientResponseError;
      return fail(401, {
        message: err.data.message,
        data: { email },
      });
    }

    throw redirect(303, '/');
  }
}`)
}

func pbLoginPageTemplate() []byte {
	return []byte(`<script lang="ts">
	import type { ActionData } from './$types';

	export let form: ActionData;
</script>

<form method="post">
	<p>
		Not an account yet?
		<a href="/signup">Sign up</a>
	</p>

	{#if form?.message}
		<p style="color: red">{form.message}</p>
	{/if}

	<label>
		Email
		<input type="text" name="email" value={form?.data.email || ''} />
	</label>

	<label>
		Password
		<input type="password" name="password" />
	</label>

	<button>Login</button>
</form>`)
}

func pbSignupServerTemplate() []byte {
	return []byte(`import { fail, redirect } from '@sveltejs/kit';
import type { ClientResponseError } from 'pocketbase';
import type { Actions } from './$types';

export const actions: Actions = {
  default: async ({ locals, request }) => {
    const { username, email, password, repeat } = Object.fromEntries(await request.formData()) as {
      username: string;
      email: string;
      password: string;
      repeat: string;
    }

    try {
      await locals.pb
        .collection('users')
        .create({ username, email, password, passwordConfirm: repeat });

      await locals.pb
        .collection('users')
        .authWithPassword(email, password);
    } catch (e) {
      const err = e as ClientResponseError;
      console.log(err.data)
      return fail(400, {
        message: err.data.message,
        data: { username, email },
        errors: err.data.data,
      });
    }

    throw redirect(303, '/');
  }
}`)
}

func pbSignupPageTemplate() []byte {
	return []byte(`<script lang="ts">
	import type { ActionData } from './$types';

	export let form: ActionData;
</script>

<form method="post">
	<p>
		Already have an account?
		<a href="/login">Login</a>
	</p>

	{#if form?.message}
		<p style="color: red">{form.message}</p>
	{/if}

	<label>
		Username
		{#if form?.errors.username}
			<span style="color: red">{form?.errors.username.message}</span>
		{/if}
		<input type="text" name="username" value={form?.data.username || ''} />
	</label>

	<label>
		Email
		{#if form?.errors.email}
			<span style="color: red">{form?.errors.email.message}</span>
		{/if}
		<input type="text" name="email" value={form?.data.email || ''} />
	</label>

	<label>
		Password
		{#if form?.errors.password}
			<span style="color: red">{form?.errors.password.message}</span>
		{/if}
		<input type="password" name="password" />
	</label>

	<label>
		Repeat password
		{#if form?.errors.passwordConfirm}
			<span style="color: red">{form?.errors.passwordConfirm.message}</span>
		{/if}
		<input type="password" name="repeat" />
	</label>

	<button>Login</button>
</form>`)
}

func pbLogoutServerTemplate() []byte {
	return []byte(`import { redirect } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ locals, cookies }) => {
    locals.pb.authStore.clear();
    cookies.delete("pb_auth");
    throw redirect(303, "/");
};`)
}

func pbDockerComposeTemplate() []byte {
	return []byte(`version: '3.8'

services:
  mailhog:
    image: mailhog/mailhog:latest
    ports:
      - 1025:1025
      - 8025:8025

  pocketbase:
    image: adrianmusante/pocketbase
    ports:
      - 8090:8090
    volumes:
      - ./pb_data:/pocketbase/data
      - ./pb_migrations:/pocketbase/migrations`)
}

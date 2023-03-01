import { fail, redirect } from '@sveltejs/kit';
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
}

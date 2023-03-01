import { fail, redirect } from '@sveltejs/kit';
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
}

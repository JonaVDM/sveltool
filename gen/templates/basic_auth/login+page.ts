import { fail, redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
  default: async ({ locals, request }) => {
    const { email, password } = Object.fromEntries(await request.formData()) as {
      email: string;
      password: string;
    }

    try {
      // TODO: log user in.
      // const token = await locals.api.login(email, password);
    } catch (e) {
      return fail(401, {
        message: 'Invalid credentials',
        data: { email },
      });
    }

    throw redirect(303, '/');
  }
}

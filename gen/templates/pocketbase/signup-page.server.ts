import { fail, redirect } from '@sveltejs/kit';
import type { ClientResponseError } from 'pocketbase';
import type { Actions, PageServerLoad } from './$types';
import { z } from 'zod';
import { superValidate } from 'sveltekit-superforms/server';

const schema = z.object({
  username: z.string().nonempty('Username cannot be empty'),
  email: z.string().email(),
  password: z.string().nonempty('Password cannot be empty'),
  passwordConfirm: z.string().nonempty(),
  message: z.string(),
}).refine((data) => data.password === data.passwordConfirm, {
  message: "Passwords don't match",
  path: ['passwordConfirm'],
});

export const load: PageServerLoad = async (event) => {
  const form = await superValidate(event, schema);
  return { form };
}

export const actions: Actions = {
  default: async (event) => {
    const form = await superValidate(event, schema);

    if (!form.valid) {
      return fail(400, { form });
    }

    try {
      await event.locals.pb
        .collection('users')
        .create(form.data);

      await event.locals.pb
        .collection('users')
        .authWithPassword(form.data.email, form.data.password);
    } catch (e) {
      const err = e as ClientResponseError;
      form.errors.message = [err.data.message]

      if (err.data.data.email) {
        form.errors.email = [err.data.data.email.message];
      }
      if (err.data.data.username) {
        form.errors.username = [err.data.data.username.message];
      }
      if (err.data.data.password) {
        form.errors.password = [err.data.data.password.message];
      }

      return fail(err.data.code, { form });
    }

    throw redirect(303, '/');
  }
}

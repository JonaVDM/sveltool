import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms/server';
import { z } from 'zod';
import type { Actions, PageServerLoad } from './$types';
import type { ClientResponseError } from 'pocketbase';

const schema = z.object({
  email: z.string().email(),
  password: z.string().nonempty(),
  message: z.string(),
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
        .authWithPassword(form.data.email, form.data.password);
    } catch (e) {
      const err = e as ClientResponseError;
      form.errors.message = [err.data.message]
      return fail(401, { form });
    }

    throw redirect(303, '/');
  }
}

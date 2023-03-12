import { redirect } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ cookies, locals }) => {
  // locals.api.setToken(''); 
  cookies.delete("auth_token");
  throw redirect(303, "/");
};

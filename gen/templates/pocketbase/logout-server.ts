import { redirect } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const GET: RequestHandler = async ({ locals, cookies }) => {
  locals.pb.authStore.clear();
  cookies.delete("pb_auth");
  throw redirect(303, "/");
};

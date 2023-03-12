import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
  const cookies = event.request.headers.get('cookie')?.split('; ') ?? [];
  const session = cookies.filter(c => c.startsWith('session='))[0]?.split('=')[1];

  // event.locals.api = new Api();
  // event.locals.api.setToken(session);

  const response = await resolve(event);

  // if (event.locals.api.isLoggedIn()) {
  //   response.headers.set(
  //     'set-cookie',
  //     `session=${session};`
  //   );
  // }

  return response;
};


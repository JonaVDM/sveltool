import { fail, redirect } from '@sveltejs/kit';
import type { Actions } from './$types';

export const actions: Actions = {
  default: async ({ request }) => {
    const { {{ range $index, $element:= .Fields }}{{ if ne $index 0}}, {{ end }}{{ .}}{{ end }} } = Object.fromEntries(await request.formData()) as { {{ range .Fields }}
      {{ . }}: string,{{ end }}
    }

    try {
      // TODO: Do something with the data
    } catch (e) {
      return fail(500, {
        // TODO: Add actual error checking
        message: "Something went terribly wrong",
        errors: { {{ range .Fields }}
          {{ . }}: "Oh no! Anyway", {{ end }}
        },
        data: { {{ range $index, $element:= .Fields }}{{ if ne $index 0}}, {{ end }}{{ .}}{{ end }} },
      });
    }

    // TODO: Redirect url
    throw redirect(303, '/');
  }
}

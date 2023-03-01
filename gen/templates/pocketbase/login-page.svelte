<script lang="ts">
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

  <div>
    <label>
      Email
      <input type="text" name="email" value={form?.data.email || ''} />
    </label>
  </div>

  <div>
    <label>
      Password
      <input type="password" name="password" />
    </label>
  </div>

  <button>Login</button>
</form>

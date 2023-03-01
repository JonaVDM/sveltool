<script lang="ts">
  import type { ActionData } from './$types';

  export let form: ActionData;
</script>

<form method="post">
  <p>
    Already have an account?
    <a href="/login">Login</a>
  </p>

  {#if form?.message}
    <p style="color: red">{form.message}</p>
  {/if}

  <div>
    <label>
      Username
      {#if form?.errors.username}
        <span style="color: red">{form?.errors.username.message}</span>
      {/if}
      <input type="text" name="username" value={form?.data.username || ''} />
    </label>
  </div>

  <div>
    <label>
      Email
      {#if form?.errors.email}
        <span style="color: red">{form?.errors.email.message}</span>
      {/if}
      <input type="text" name="email" value={form?.data.email || ''} />
    </label>
  </div>

  <div>
    <label>
      Password
      {#if form?.errors.password}
        <span style="color: red">{form?.errors.password.message}</span>
      {/if}
      <input type="password" name="password" />
    </label>
  </div>

  <div>
    <label>
      Repeat password
      {#if form?.errors.passwordConfirm}
        <span style="color: red">{form?.errors.passwordConfirm.message}</span>
      {/if}
      <input type="password" name="repeat" />
    </label>
  </div>

  <button>Login</button>
</form>

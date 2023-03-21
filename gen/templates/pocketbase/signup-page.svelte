<script lang="ts">
  import { superForm } from 'sveltekit-superforms/client';
  import Input from '$lib/Input.svelte';
  import type { PageData } from './$types';

  export let data: PageData;

  const { fields, errors } = superForm(data.form);
</script>

<form method="post">
  <p>
    Already have an account?
    <a href="/login">Login</a>
  </p>

  {#if $errors.message}
    <p class="error">{$errors.message}</p>
  {/if}

  <Input label="Username" field={fields.username} />
  <Input label="Email" field={fields.email} />
  <Input type="password" label="Password" field={fields.password} />
  <Input
    type="password"
    label="Repeat password"
    field={fields.passwordConfirm}
  />

  <button>Login</button>
</form>

<style>
  .error {
    color: red;
  }
</style>

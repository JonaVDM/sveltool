<script lang="ts">
  import Input from "$lib/Input.svelte";
  import { superForm } from "sveltekit-superforms/client";
  import type { PageData } from "./$types";

  export let data: PageData;

  const { fields, errors } = superForm(data.form);
</script>

<form method="post">
  <p>
    Not an account yet?
    <a href="/signup">Sign up</a>
  </p>

  {#if $errors.message}
    <p class="error">{$errors.message}</p>
  {/if}

  <Input label="Email" field={fields.email} />
  <Input type="password" label="Password" field={fields.password} />

  <button>Login</button>
</form>

<style>
  .error {
    color: red;
  }
</style>

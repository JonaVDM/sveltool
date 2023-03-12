<script lang="ts">
  import type { ActionData } from "./$types";

  export let form: ActionData;
</script>

<svelte:head>
  <title>My cool app - Login</title>

  <!--
    Don't actually put this here, this belongs in some sorts of global styles
  -->
  <style>
    * {
      margin: 0;
    }
  </style>
</svelte:head>

<div class="container">
  <form method="post">
    <h1>Login</h1>

    {#if form?.message}
      <p class="error">{form.message}</p>
    {/if}

    <div>
      <label>
        Email
        <input type="text" name="email" value={form?.data.email || ""} />
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
</div>

<style>
  .error {
    color: red;
  }

  .container {
    width: 100vw;
    height: 100vh;

    display: flex;
    justify-content: center;
    align-items: center;
  }

  form {
    border: 1px solid black;
    padding: 2rem;
  }

  input {
    width: 100%;
    margin-bottom: 1rem;
  }
</style>

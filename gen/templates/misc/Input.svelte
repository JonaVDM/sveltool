<script lang="ts">
  import type { FormField } from "sveltekit-superforms";
  import type { z, AnyZodObject } from "zod";

  type T = $$Generic<AnyZodObject>;
  type P = $$Generic<keyof z.infer<T>>;

  export let type = "text";
  export let label: string;
  export let field: FormField<T, P>;
  export let placeholder = "";

  function setType(el: HTMLInputElement) {
    el.type = type;
  }

  $: value = field.value;
  $: errors = field.errors;
</script>

<div>
  <label>
    {label}
    {#if $errors}
      - <span class="error">{$errors}</span>
    {/if}

    {#if type == "password"}
      <input
        type="password"
        {placeholder}
        bind:value={$value}
        name={field.name}
      />
    {:else}
      <input use:setType {placeholder} bind:value={$value} name={field.name} />
    {/if}
  </label>
</div>

<style>
  .error {
    color: red;
  }
</style>

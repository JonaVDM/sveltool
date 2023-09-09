<script lang="ts">
  import type { Writable } from "svelte/store";
  import type { z, AnyZodObject } from "zod";
  import type { ZodValidation, FormPathLeaves } from "sveltekit-superforms";
  import { formFieldProxy, type SuperForm } from "sveltekit-superforms/client";

  type T = $$Generic<AnyZodObject>;

  export let form: SuperForm<ZodValidation<T>, unknown>;
  export let field: FormPathLeaves<z.infer<T>>;
  export let label: string | null;

  $: computedLabel = label ? label : field;
  $: boolValue = value as Writable<boolean>;

  const { value, errors, constraints } = formFieldProxy(form, field);
</script>

<label class="py-2 flex gap-2 cursor-pointer">
  <input
    class="checkbox checkbox-primary"
    name={field}
    type="checkbox"
    aria-invalid={$errors ? "true" : undefined}
    bind:checked={$boolValue}
    {...$constraints}
    {...$$restProps}
  />
  {computedLabel}
  {#if $errors} - <span class="text-red-600">{$errors}</span>{/if}
</label>

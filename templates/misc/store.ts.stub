import { writable } from 'svelte/store';

function create{{.Name}}() {
  // For a basic store with some added functionality
  const store = writable(0);
  return {
    ...store,
    increment: () => store.update(n => n + 1),
  }

  // For a complete custom store uncomment the following block.
  //
  // const { subscribe, set, update } = writable(0);
  // return {
  //   subscribe,
  //   increment: () => update(n => n + 1),
  //   decrement: () => update(n => n - 1),
  //   reset: () => set(0)
  // };
}

export const {{.ExportName}} = create{{.Name}}();

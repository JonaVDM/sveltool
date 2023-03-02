<script lang="ts">
  import { page } from "$app/stores";

  interface Item {
    url: string;
    title: string;
  }

  {{if .Auth}}
  export let authed = false;

  const loggedIn: Item[] = [
    {
      url: "/",
      title: "Home",
    },
    {
      url: "/private",
      title: "Some private route",
    },
    {
      url: "/logout",
      title: "Logout",
    },
  ];

  const loggedOut: Item[] = [
    {
      url: "/",
      title: "Home",
    },
    {
      url: "/login",
      title: "Login",
    },
  ];
  $: items = authed ? loggedIn : loggedOut;
  {{else}}
  let items: Item[] = [
    {
      url: '/',
      title: 'Home'
    },
    {
      url: '/about',
      title: 'About Us'
    },
    {
      url: '/contact',
      title: 'Contact'
    },
  ]
  {{end}}
</script>

<nav>
  <ul>
    {#each items as { url, title }}
      <li>
        <a class:current={$page.url.pathname == url} href={url}>{title}</a>
      </li>
    {/each}
  </ul>
</nav>

<style>
  ul {
    display: flex;
    margin: 0;
    padding: 0;
  }

  li {
    list-style: none;
    padding: 0.5rem 1.2rem;
  }

  a {
    color: black;
  }

  .current {
    color: green;
  }
</style>

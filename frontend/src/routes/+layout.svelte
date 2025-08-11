<script lang="ts">
	import '$lib/styles/_theme.scss';
	import Header from '$lib/components/Header.svelte';
	import Sidebar from '$lib/components/Sidebar.svelte';
  import favicon from '$lib/assets/favicon.svg';

	let { children } = $props();
  let sidebarOpen = false;

  function toggleSidebar() {
    sidebarOpen = !sidebarOpen;
  }
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="app-container">
  <Header on:menu={toggleSidebar} />
  <Sidebar bind:open={sidebarOpen} />

  <main class="app-content">
    {@render children?.()}
  </main>
</div>

<style>
  :global(body) {
    margin: 0;
  }

  .app-container {
    display: flex;
    flex-direction: column;
    height: 100vh;
  }

  .app-content {
    flex-grow: 1;
    padding: 1rem;
    overflow: auto;
  }
</style>

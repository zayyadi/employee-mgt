<script lang="ts">
  import Card, { Content } from '@smui/card';
  import Button from '@smui/button';
  import Textfield from '@smui/textfield';
  import authStore from '$lib/stores/authStore';
  import { goto } from '$app/navigation';

  let email = '';
  let password = '';
  let loading = false;
  let error: string | null = null;

  async function handleSubmit() {
    loading = true;
    error = null;
    try {
      await authStore.login(email, password);
      goto('/'); // Redirect to dashboard on success
    } catch (e: any) {
      error = e.message;
    } finally {
      loading = false;
    }
  }
</script>

<Card style="width: 400px; margin: 2rem auto;">
  <form on:submit|preventDefault={handleSubmit}>
    <Content class="mdc-typography--headline6" style="padding: 1rem;">
      Login
    </Content>
    <div style="padding: 0 1rem 1rem 1rem;">
      <Textfield type="email" label="Email" bind:value={email} style="width: 100%;" required />
      <div style="height: 1rem;"></div>
      <Textfield type="password" label="Password" bind:value={password} style="width: 100%;" required />

      {#if error}
        <p style="color: red; margin-top: 1rem;">{error}</p>
      {/if}
    </div>
    <div style="padding: 1rem; display: flex; justify-content: flex-end;">
      <Button type="submit" variant="unelevated" disabled={loading}>
        {#if loading}
          Loading...
        {:else}
          Login
        {/if}
      </Button>
    </div>
  </form>
</Card>

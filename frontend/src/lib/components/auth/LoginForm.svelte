<script lang="ts">
  import Card, { Content, Media, MediaContent, PrimaryAction } from '@smui/card';
  import Button from '@smui/button';
  import Textfield from '@smui/textfield';

  let email = '';
  let password = '';
</script>

<Card style="width: 400px; margin: 2rem auto;">
  <Content class="mdc-typography--headline6" style="padding: 1rem;">
    Login
  </Content>
  <div style="padding: 0 1rem 1rem 1rem;">
    <Textfield type="email" label="Email" bind:value={email} style="width: 100%;" />
    <div style="height: 1rem;"></div>
    <Textfield type="password" label="Password" bind:value={password} style="width: 100%;" />
  </div>
  <div style="padding: 1rem; display: flex; justify-content: flex-end;">
    <Button variant="unelevated">Login</Button>
  </div>
</Card>

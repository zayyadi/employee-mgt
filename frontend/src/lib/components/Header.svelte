<script lang="ts">
  import TopAppBar, { Row, Section, Title } from '@smui/top-app-bar';
  import IconButton from '@smui/icon-button';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher();

  let mainTitle = 'Employee Management';

  function handleMenuClick() {
    dispatch('menu');
  }
</script>

<TopAppBar variant="static">
  <Row>
    <Section>
      <IconButton class="material-icons" on:click={handleMenuClick}>menu</IconButton>
      <Title>{mainTitle}</Title>
    </Section>
  </Row>
</TopAppBar>

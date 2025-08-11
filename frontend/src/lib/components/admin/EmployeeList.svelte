<script lang="ts">
  import { onMount } from 'svelte';
  import DataTable, { Head, Body, Row, Cell } from '@smui/data-table';
  import employeeStore from '$lib/stores/employeeStore';

  onMount(() => {
    employeeStore.fetchAll();
  });
</script>

<h1>Employees</h1>

{#if $employeeStore.loading}
  <p>Loading...</p>
{:else if $employeeStore.error}
  <p style="color: red;">Error: {$employeeStore.error}</p>
{:else}
  <DataTable>
    <Head>
      <Row>
        <Cell>ID</Cell>
        <Cell>Name</Cell>
        <Cell>Email</Cell>
        <Cell>Position</Cell>
      </Row>
    </Head>
    <Body>
      {#each $employeeStore.employees as employee (employee.id)}
        <Row>
          <Cell>{employee.id}</Cell>
          <Cell>{employee.first_name} {employee.last_name}</Cell>
          <Cell>{employee.email}</Cell>
          <Cell>{employee.position}</Cell>
        </Row>
      {/each}
    </Body>
  </DataTable>
{/if}

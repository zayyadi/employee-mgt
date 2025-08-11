<script lang="ts">
  import Drawer, { Content, Header, Title, Subtitle } from '@smui/drawer';
  import List, { Item, Text } from '@smui/list';

  export let open = false;
</script>

<Drawer variant="modal" bind:open>
  <Header>
    <Title>Menu</Title>
    <Subtitle>Navigation</Subtitle>
  </Header>
  <Content>
    <List>
      <Item href="/">
        <Text>Home</Text>
      </Item>
      <Item href="/admin/employees">
        <Text>Employees</Text>
      </Item>
      <Item href="/auth/login">
        <Text>Login</Text>
      </Item>
    </List>
  </Content>
</Drawer>

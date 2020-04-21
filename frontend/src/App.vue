<template>
  <div class="bg-black">
    <q-layout>
      <!-- Header Bar -->
      <q-header>
        <q-page-sticky position="top" expand>
          <header-bar />
        </q-page-sticky>
      </q-header>
      
      <!-- Page Content -->
      <q-page-container>
        <q-page style="padding-top: 50px">
          <router-view />
        </q-page>
      </q-page-container>
    </q-layout>
  </div>
</template>

<script>
import HeaderBar from './components/HeaderBar';

export default {
  components: {
    headerBar: HeaderBar
  }
}
</script>

<style>
a {
  color: #9e9e9e;
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}
</style>

<template>
  <div class="bg-black">
    <q-layout>
      <!-- Header Bar -->
      <q-header>
        <q-page-sticky position="top" expand>
          <header-bar></header-bar>
        </q-page-sticky>
      </q-header>
      
      <!-- Page Content -->
      <q-page-container>
        <q-page style="padding-top: 50px">
          <router-view></router-view>
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

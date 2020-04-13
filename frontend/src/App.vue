<template>
  <div class="bg-black">
    <header-bar></header-bar>
    <router-view></router-view>
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
</style>

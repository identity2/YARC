<template>
  <q-toolbar class="bg-grey-10 text-white">
    <q-avatar @click="linkToHome" class="cursor-pointer">
      <img src="../assets/logo.jpg" alt="YARC" />
    </q-avatar>

    <q-toolbar-title @click="linkToHome" class="cursor-pointer">
        YARC
    </q-toolbar-title>

    <q-input class="q-ml-xl" style="width: 400px" dark clearable outlined dense standout v-model="searchText" label="Search" />
    <q-space />
    <q-btn stretch flat label="Log in" />
    <q-btn stretch flat label="Register" />
  </q-toolbar>
</template>

<script>
export default {
  data() {
    return {
      searchText: ''
    };
  },
  methods: {
    linkToHome() {
      if (this.$route.path !== '/') {
        // Get back to home.
        this.$router.push('/');
      } else {
        // Refresh the page.
        location.href = this.$route.fullPath;
      }
    }
  }
}
</script>

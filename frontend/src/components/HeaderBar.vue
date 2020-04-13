<template>
  <div class>
    <q-toolbar class="bg-grey-10 text-white">
      <q-avatar>
        <img src="../assets/logo.jpg" alt="YARC" />
      </q-avatar>

      <q-toolbar-title>
        YARC
      </q-toolbar-title>

      <q-input class="q-ml-xl" style="width: 400px" dark clearable outlined dense standout v-model="searchText" label="Search" />
      <q-space />
      <q-btn stretch flat label="Log in" />
      <q-btn stretch flat label="Register" />
    </q-toolbar>
  </div>
</template>

<script>
export default {
  data() {
    return {
      searchText: ''
    };
  }
}
</script>

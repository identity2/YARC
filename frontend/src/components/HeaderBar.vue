<template>
  <q-toolbar class="bg-grey-10 text-white">
    <q-avatar @click="linkToHome" class="cursor-pointer">
      <img src="../assets/logo.jpg" alt="YARC" />
    </q-avatar>

    <q-toolbar-title @click="linkToHome" class="cursor-pointer">
        YARC
    </q-toolbar-title>

    <q-input @keydown.enter.prevent="submitSearch" class="q-ml-xl" style="width: 400px" dark clearable outlined dense standout v-model="searchText" label="Search" :maxlength="searchTermMaxLen" />
    <q-space />

    <q-btn v-if="showLoginRegister" @click="loginRegisterClicked('login')" stretch flat label="Log in" />
    <q-btn v-if="showLoginRegister" @click="loginRegisterClicked('register')" stretch flat label="Register" />
  </q-toolbar>
</template>

<script>
import Limits from '../limits';

export default {
  data() {
    return {
      searchText: '',
      searchTermMaxLen: Limits.searchTermMaxLen
    };
  },
  computed: {
    showLoginRegister() {
      if (this.$route.path === '/account') {
        return false;
      }

      // TODO: Check log in.

      return true;
    }
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
    },
    loginRegisterClicked(type) {
      this.$router.push({
        name: 'account',
        query: {
          type: type
        }
      });
    },
    submitSearch() {
      this.$router.push({
        name: 'search',
        query: {
          q: this.searchText
        }
      });
    }
  }
}
</script>

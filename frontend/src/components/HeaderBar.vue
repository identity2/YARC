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

    <!-- Log in, Register -->
    <q-btn v-if="showLoginRegister" @click="loginRegisterClicked('login')" stretch flat label="Log in" />
    <q-btn v-if="showLoginRegister" @click="loginRegisterClicked('register')" stretch flat label="Register" />
    
    <!-- Profile, Log out -->
    <q-btn no-caps v-if="!showLoginRegister" @click="showProfileClicked()" stretch flat :label="$store.state.auth.user.username" />
    <q-btn v-if="!showLoginRegister" @click="logoutClicked()" stretch flat label="Log out" />
  </q-toolbar>
</template>

<script>
import Limits from '../limits';
import AuthService from '../services/authorization';

export default {
  data() {
    return {
      searchText: '',
      searchTermMaxLen: Limits.searchTermMaxLen,

    };
  },
  computed: {
    showLoginRegister() {
      // Check if the user is logged in.
      if (this.$store.state.auth.loggedIn) {
        return false;
      }

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
    showProfileClicked() {
      this.$router.push({
        path: '/u/' + this.$store.state.auth.user.username
      });
    },
    logoutClicked() {
      AuthService.logout();
      this.$store.commit('logout');

      // Refresh the page.
      location.href = this.$route.fullPath;
    },
    submitSearch() {
      if (this.searchText.trim() === '') {
        return
      }

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

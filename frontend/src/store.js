import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

// user should contain an authHeader and a username.
const user = JSON.parse(localStorage.getItem("user"));
const initialAuthState = user ? {loggedIn: true, user: user} : {loggedIn: false, user: null};

export const store = new Vuex.Store({
  state: {
    auth: initialAuthState
  },
  mutations: {
    logout(state) {
      state.auth.loggedIn = false;
      state.auth.user = null;
    },
    loginSuccess(state, user) {
      state.auth.loggedIn = true;
      state.auth.user = user;
    }
  }
});
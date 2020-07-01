import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

// user should contain an authHeader and a username.
const expireTime = 1000 * 60 * 60 * 24 * 7; // 7 days.
const user = JSON.parse(localStorage.getItem("user"));
const currTime = new Date().getTime();
const initialAuthState = user && currTime - user.timestamp < expireTime ? user : null;

export const store = new Vuex.Store({
  state: {
    auth: initialAuthState
  },
  mutations: {
    logout(state) {
      state.auth = null;
    },
    loginSuccess(state, user) {
      state.auth = user;
    }
  }
});
<template>
  <div class="row">
    <div class="col q-pa-lg">
      <!-- Log In / Register Text -->
      <div class="text-h6 text-white">Log In / Register</div>
      <q-separator color="grey" />

      <!-- Log In / Register Box -->
      <q-card dark class="q-mt-lg">
        <q-tabs v-model="formType">
          <q-tab name="login" label="Log In" />
          <q-tab name="register" label="Register" />
        </q-tabs>

        <q-separator color="grey-9" />

        <q-tab-panels class="bg-grey-10 q-pt-md q-pb-md q-pl-xl q-pr-xl" v-model="formType" animated>
          <q-tab-panel name="login">
            <div class="q-mb-sm text-subtitle">Username:</div>
            <q-input class="q-mb-lg" dark outlined dense standout v-model="username" />

            <div class="q-mb-sm text-subtitle">Password:</div>
            <q-input class="q-mb-md" type="password" dark outlined dense standout v-model="password" />

            <q-card-actions align="right">
              <q-btn @click="loginClicked" style="background: white; color: black; width: 100px" label="Log In" />
            </q-card-actions>
          </q-tab-panel>

          <q-tab-panel name="register">
            <div class="q-mb-sm text-subtitle">Username:</div>
            <q-input class="q-mb-md" dark outlined dense standout v-model="username" counter :maxlength="usernameMaxLen" />
            
            <div class="q-mb-sm text-subtitle">Email Address:</div>
            <q-input class="q-mb-xl" dark outlined dense standout v-model="emailAddress" maxlength="256" />
            
            <div class="q-mb-sm text-subtitle">Password:</div>
            <q-input class="q-mb-md" type="password" dark outlined dense standout v-model="password" counter :maxlength="passwordMaxLen" />
            
            <div class="q-mb-sm text-subtitle">Reenter Password:</div>
            <q-input class="q-mb-md" type="password" dark outlined dense standout v-model="passwordConfirm" :maxlength="passwordMaxLen" />
            
            <q-card-actions align="right">
              <q-btn @click="registerClicked" style="background: white; color: black; width: 100px" label="Register" />
            </q-card-actions>
          </q-tab-panel>
        </q-tab-panels>
      </q-card>
    </div>

    <!-- Right Panel -->
    <div class="col-3 q-pr-md q-pt-md gt-sm">
      <tips
        title="Account Rules"
        :tips="[
          '1. There are actually no rules.',
          '2. Except this one.',
          '3. There\'s three actually.'
        ]"
      />
      <advertisement />
    </div>
  </div>
</template>

<script>
import Tips from '../rightPanel/Tips';
import Advertisement from '../rightPanel/Advertisement';
import Limits from '../../limits';

export default {
  mounted() {
    document.title = 'Log In / Register - YARC';
  },
  computed: {
    validUsername() {
      return true;
    },
    validPassword() {
      return true;
    },
    validEmail() {
      return true;
    }
  },
  methods: {
    loginClicked() {

    },
    registerClicked() {

    }
  },
  data() {
    return {
      formType: 'login',
      username: '',
      usernameMaxLen: Limits.usernameMaxLen,
      password: '',
      passwordMaxLen: Limits.passwordMaxLen,
      emailAddress: '',
      passwordConfirm: ''
    };
  },
  watch: {
    formType(newVal) {
      this.$router.replace({
        query: {
          type: newVal
        }
      }).catch(() => {});
    }
  },
  created() {
    this.formType = this.$route.query.type;
    if (!this.formType) {
      this.formType = 'login';
    }
  },
  components: {
    advertisement: Advertisement,
    tips: Tips
  }
}
</script>
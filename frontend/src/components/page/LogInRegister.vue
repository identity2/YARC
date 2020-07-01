<template>
  <div class="row">
    <div class="col q-pa-lg">
      <!-- Log In / Register Text -->
      <div class="text-h6 text-white">Log In / Register</div>
      <q-separator color="grey" />

      <!-- Log In / Register Box -->
      <q-card dark class="q-mt-lg">
        <q-tabs v-model="formType">
          <q-tab name="login" label="Log In" :disabled="loading" />
          <q-tab name="register" label="Register" :disabled="loading" />
        </q-tabs>

        <q-separator color="grey-9" />

        <q-tab-panels class="bg-grey-10 q-pt-md q-pb-md q-pl-xl q-pr-xl" v-model="formType" animated>
          <q-tab-panel name="login">
            <q-banner v-if="loginErrMessage != ''" inline-actions class="text-white bg-red q-mb-md">
              <span class="text-h6">Error: {{loginErrMessage}}.</span>
            </q-banner>
            
            <div class="q-mb-sm text-subtitle">Username:</div>
            <q-input class="q-mb-lg" dark outlined dense standout v-model="username" :disabled="loading" />

            <div class="q-mb-sm text-subtitle">Password:</div>
            <q-input class="q-mb-md" type="password" dark outlined dense standout v-model="password" :disabled="loading" />

            <q-card-actions align="right">
              <q-btn @click="loginClicked" style="background: white; color: black; width: 100px" label="Log In" :loading="loading" />
            </q-card-actions>
          </q-tab-panel>

          <q-tab-panel name="register">
            <q-banner v-if="registerErrMessage != ''" inline-actions class="text-white bg-red q-mb-md">
              <span class="text-h6">Error: {{registerErrMessage}}.</span>
            </q-banner>

            <div class="q-mb-sm text-subtitle">
              Username:
              <span class="text-grey">({{usernameMinLen}} to {{usernameMaxLen}} alpha-numeric or underscore characters)</span>
            </div>
            <q-input class="q-mb-md" dark outlined dense standout v-model="username" counter :maxlength="usernameMaxLen" :disabled="loading" />
            
            <div class="q-mb-sm text-subtitle">Email Address:</div>
            <q-input class="q-mb-xl" dark outlined dense standout v-model="emailAddress" maxlength="256" :disabled="loading" />
            
            <div class="q-mb-sm text-subtitle">
              Password:
              <span class="text-grey">({{passwordMinLen}} to {{passwordMaxLen}} alpha-numeric or underscore characters)</span>
            </div>
            <q-input class="q-mb-md" type="password" dark outlined dense standout v-model="password" counter :maxlength="passwordMaxLen" :disabled="loading" />
            
            <div class="q-mb-sm text-subtitle">Reenter Password:</div>
            <q-input class="q-mb-md" type="password" dark outlined dense standout v-model="passwordConfirm" :maxlength="passwordMaxLen" :disabled="loading" />
            
            <q-card-actions align="right">
              <q-btn @click="registerClicked" style="background: white; color: black; width: 100px" label="Register" :loading="loading" />
            </q-card-actions>
          </q-tab-panel>
        </q-tab-panels>
      </q-card>

      <!-- Register successful diolog -->
      <q-dialog v-model="showRegisterSuccessful" persistent>
        <q-card class="bg-blue text-white" style="width: 300px">
          <q-card-section>
            <div class="text-h5">Congratulations, {{username}}!</div>
          </q-card-section>

          <q-card-section class="q-pt-none">
            You have been registered successfully! You can now log in.
          </q-card-section>

          <q-card-actions align="right" class="bg-white text-teal">
            <q-btn flat label="Log In" @click="registerSuccessLogInClicked"/>
          </q-card-actions>
        </q-card>
      </q-dialog>
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
import AuthService from '../../services/authorization';

export default {
  created() {
    // If already logged in, redirect to the home page.
    if (this.$store.state.auth) {
      this.$router.push('/');
    }

    // Check if the action is login or register.
    this.formType = this.$route.query.type;
    if (!this.formType) {
      this.formType = 'login';
    }
  },
  mounted() {
    document.title = 'Log In / Register - YARC';
  },
  computed: {
    validUsername() {
      const re = /^[a-zA-Z0-9_]*$/;
      if (!re.test(this.username) || this.username > this.usernameMaxLen || this.username < this.usernameMinLen) {
        return false;
      }
      return true;
    },
    validPassword() {
      const re = /^[a-zA-Z0-9_]*$/;
      if (!re.test(this.password) || this.password > this.passwordMaxLen || this.password < this.passwordMinLen) {
        return false;
      }
      return true;
    },
    validPasswordConfirm() {
      if (this.password !== this.passwordConfirm) {
        return false;
      }
      return true;
    },
    validEmail() {
      const re = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/;
      return re.test(this.emailAddress);
    }
  },
  methods: {
    loginClicked() {
      // Client side validation.
      if (!this.validUsername || !this.validPassword) {
        this.loginErrMessage = "the username or the password is invalid";
        return;
      }

      // Send login request to the server.
      this.loading = true;
      AuthService.login(this.username, this.password).then(user => {
        // Store user info to the global state.
        this.$store.commit("loginSuccess", user);
        
        // Return to the previous page.
        this.$router.go(-1);
      }).catch(error => {
        // Report the error to the user.
        this.loginErrMessage = "failed to login (unknown reasons)";
        if (error.response) {
          this.loginErrMessage = error.response.data.error;
        }
        
        this.loading = false;
      });
    },
    registerClicked() {
      // Client side validation.
      if (!this.validUsername) {
        this.registerErrMessage = 'the username is invalid';
        return;
      }
      if (!this.validEmail) {
        this.registerErrMessage = 'the email address is invalid';
        return;
      }
      if (!this.validPassword) {
        this.registerErrMessage = 'the password is invalid';
        return;
      }
      if (!this.validPasswordConfirm) {
        this.registerErrMessage = 'the reentered password does not match'
        return;
      }

      // Send request.
      this.loading = true;
      AuthService.register(this.username, this.password, this.emailAddress).then(() => {
        this.showRegisterSuccessful = true;
      }).catch(error => {
        // Report the error to the user.
        this.registerErrMessage = "failed to register (unknown reasons)";
        if (error.response) {
          this.registerErrMessage = error.response.data.error;
        }
        
        this.loading = false;
      });
    },
    registerSuccessLogInClicked() {
      location.href = this.$route.path + '?type=login';
    }
  },
  data() {
    return {
      formType: 'login',
      username: '',
      usernameMinLen: Limits.usernameMinLen,
      usernameMaxLen: Limits.usernameMaxLen,
      password: '',
      passwordMaxLen: Limits.passwordMaxLen,
      passwordMinLen: Limits.passwordMinLen,
      emailAddress: '',
      emailAddressMaxLen: Limits.emailAddressMaxLen,
      passwordConfirm: '',
      loading: false,
      loginErrMessage: '',
      registerErrMessage: '',
      showRegisterSuccessful: false
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
  components: {
    advertisement: Advertisement,
    tips: Tips
  }
}
</script>
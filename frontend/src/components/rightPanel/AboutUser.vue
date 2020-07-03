<template>
  <q-card dark flat bordered class="bg-grey-10 q-mb-md">
    <q-card-section>
      <div class="text-h6 q-mb-md">About u/{{userInfo.username}}</div>

      <!-- Short Bio -->
      <div v-if="!editMode">{{userInfo.bio}}</div>
      <div v-if="editMode">
        <q-input dark outlined standout v-model="currBio" type="textarea" counter :maxlength="shortBioMaxLen" :disable="loading" />
      </div>
    
    </q-card-section>

    <q-card-actions align="center">
      <q-btn v-if="!editMode && canEdit" @click="editMode = !editMode" class="q-mb-md" flat style="background: white; color: black; width: 50%" label="Edit" />
      <q-btn v-if="editMode" @click="editMode = !editMode" class="q-mb-md" outline style="color: red; width: 30%" label="Cancel" :disabled="loading" />
      <q-btn v-if="editMode" @click="saveClicked" class="q-mb-md" flat style="background: white; color: black; width: 30%" label="Save" :loading="loading" />
    </q-card-actions>

    <q-separator dark inset />

    <q-card-section>
      <div class="text-center text-weight-bold">Karma: {{userInfo.karma}}</div>
      <div class="text-center text-grey q-mt-md">Joined at {{userInfo.joinTime | formatDate}}.</div>
    </q-card-section>

    <div v-if="errMsg != ''" class="text-red text-subtitle q-mt-md">{{errMsg}}</div>
  </q-card>
</template>

<script>
import Limits from '../../limits';
import AccountService from '../../services/account';

export default {
  props: {
    userInfo: Object
  },
  data() {
    return {
      editMode: false,
      shortBioMaxLen: Limits.shortBioMaxLen,
      currBio: '',
      errMsg: '',
      loading: false
    };
  },
  computed: {
    canEdit() {
      let auth = this.$store.state.auth;
      return auth && auth.username === this.userInfo.username;
    }
  },
  watch: {
    editMode() {
      this.currBio = this.userInfo.bio;
      this.errMsg = '';
      this.loading = false;
    }
  },
  methods: {
    saveClicked() {
      this.loading = true;

      AccountService.modify(this.currBio, this.$store.state.auth.authHeader).then(() => {
        // Successful, inform the parent.
        this.$emit('bioSaved', this.currBio);
        this.editMode = false;
      }).catch(error => {
        if (error.response && error.response.status === 401) {
          this.errMsg = 'login credentials expired, please re-login';
          localStorage.removeItem('user');
          this.$store.commit('logout');
        } else if (error.response && error.response.body) {
          this.errMsg = error.response.body;
        } else {
          this.errMsg = 'an unknown error occurred, try again';
        }
        this.loading = false;
      });
    }
  }
}
</script>
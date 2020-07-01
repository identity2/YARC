<template>
  <div>
    <q-card-section>
      <div class="text-h5">{{article.title}}</div>
      
      <!-- Text Post -->
      <div v-if="!editMode && article.type === 'text'" class="q-mt-lg">{{article.body}}</div>
      <div v-if="editMode" class="q-mt-lg">
        <div class="q-ml-xs q-mb-sm">Editing post:</div>
        <q-input dark outlined standout v-model="currBody" type="textarea" counter :maxlength="textPostMaxLen" :disable="loading" />
        <div v-if="editMode && errMsg != ''" class="q-mb-md text-red">Error: {{errMsg}}.</div>
        <q-btn class="q-mr-md" @click="saveClicked" style="background: white; color: black" label="Save" :loading="loading" />
        <q-btn outline @click="cancelClicked" style="color: red" label="Cancel" :disabled="loading" />
      </div>

      <!-- Image Post -->
      <div v-if="article.type === 'image'" class="q-mt-lg">
        <img :src="article.body" style="max-width: 100%; max-height: 100%" />
      </div>

      <!-- Link Post -->
      <div v-if="article.type === 'link'" class="q-mt-lg">
        <a :href="article.body" class="text-blue" target="_blank">
          {{article.body}}
          <q-icon name="open_in_new" size="xs" />
        </a>
      </div>
    </q-card-section>
  </div>
</template>

<script>
import Limits from '../../limits';
import ArticleService from '../../services/article';

export default {
  props: {
    article: Object,
    editMode: Boolean
  },
  data() {
    return {
      textPostMaxLen: Limits.textPostMaxLen,
      currBody: this.article.body,
      errMsg: '',
      loading: false
    };
  },
  watch: {
    editMode() {
      this.errOccurred = false;
      this.loading = false;
    }
  },
  methods: {
    saveClicked() {
      this.loading = true;

      // Send request to the server.
      ArticleService.modify(this.article.articleID, this.currBody, this.$store.state.auth.authHeader).then(() => {
        // Inform the parent about the change.
        this.$emit('saveEdit', this.currBody);
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
    },
    cancelClicked() {
      this.currBody = this.article.body;  // Reset to the original content.
      this.$emit('cancelEdit');
    }
  }
}
</script>

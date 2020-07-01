<template>
  <div class="q-mt-md">
    <q-separator dark />
    <q-item-section>
      <div v-if="loggedInUser" class="text-grey text-subtitle2 q-mt-md q-mb-sm q-ml-xs">
        <span>Add a comment as </span>
        <router-link :to="'/u/' + loggedInUser.username">
          <span class="text-blue">r/{{loggedInUser.username}}</span>
        </router-link>
      </div>

      <div v-if="!loggedInUser" class="text-red text-subtitle2 q-mt-md q-mb-sm q-ml-xs">
        <router-link class="text-blue" :to="'/account?login'">Login</router-link> to add a comment!
      </div>

      <q-input dark outlined standout v-model="commentText" type="textarea" placeholder="What are your thoughts?" counter :maxlength="commentMaxLen" :disable="!loggedInUser || loading" />
    </q-item-section>

    <q-btn class="q-mr-md q-mb-lg" @click="commentClicked" style="background: white; color: black" label="Comment" :disabled="!loggedInUser" :loading="loading" />

    <!-- Error Indicator -->
    <div v-if="errMsg != ''" class="text-red q-mb-md">
      Error: {{errMsg}}.
    </div>

    <q-separator dark />
  </div>
</template>

<script>
import Limits from '../../limits';
import CommentService from '../../services/comment';

export default {
  props: {
    articleID: String,
    loggedInUser: Object
  },
  data() {
    return {
      commentText: '',
      commentMaxLen: Limits.commentMaxLen,
      loading: false,
      errMsg: ''
    };
  },
  methods: {
    commentClicked() {
      this.errMsg = '';

      let body = this.commentText.trim();
      if (body === '') {
        this.errMsg = 'the comment body cannot be empty';
        return;
      }

      // Send request to the server.
      this.loading = true;
      this.commentText = '';
      CommentService.create(this.articleID, body, this.loggedInUser.authHeader).then(commentID => {
        // Get the response so that it can be appended to the article list.
        CommentService.get(commentID).then(comment => {
          this.$emit('commentCreated', comment);
          this.loading = false;
        }).catch(() => {
          this.errMsg = 'comment created, but failed to show. Please refresh the page';
          this.loading = false;
        });
      }).catch(error => {
        if (error.response && error.response.status === 401) {
          this.$store.commit('logout');
          localStorage.removeItem('user');
        } else {
          this.errMsg = 'failed to post the comment, please try again';
        }
        this.loading = false;
      });
    }
  }
}
</script>
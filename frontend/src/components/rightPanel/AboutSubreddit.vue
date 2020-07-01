<template>
  <q-card dark flat bordered class="bg-grey-10 q-mb-md">
    <q-card-section>
      <div class="text-h6 q-mb-md">r/{{subreddit}}</div>
      <div>{{description}}</div>
    </q-card-section>

    <q-separator dark inset />

    <q-card-section>
      <div class="text-center">{{members}} Members</div>
    </q-card-section>


    <q-card-section class="column q-ml-md q-mr-md q-gutter-sm">
      <q-btn @click="createPostClicked" color="white" text-color="black" class="row" label="Create Post" />
      <q-btn
        class="row"
        v-if="loggedInUser"
        @click="joinClicked"
        color="white"
        text-color="black"
        :label="joined ? 'Leave' : 'Join'"
        :loading="loading"
      />
      <div v-if="errMsg != ''" class="text-red text-subtitle q-mt-md">{{errMsg}}</div>
    </q-card-section>
  </q-card>
</template>

<script>
import AccountService from '../../services/account';
import SubredditService from '../../services/subreddit';

export default {
  props: {
    subreddit: String
  },
  data() {
    return {
      loggedInUser: null,
      members: 0,
      description: '',
      joined: false,
      loading: false,
      errMsg: ''
    }
  },
  mounted() {
    this.loggedInUser = this.$store.state.auth;
  },
  watch: {
    // We need to fetch the join state using watch because the parent component
    // fetches the subreddit name asynchronizedly, making this.subreddit become ''
    // when mounted.
    async subreddit() {
      this.loading = true;
      this.errMsg = '';
      // Fetch the subreddit information.
      try {
        let subInfo = await SubredditService.get(this.subreddit);
        this.members = subInfo.members;
        this.description = subInfo.description;
      } catch (error) {
        this.errMsg = 'Failed to fetch subreddit information.';
      }

      // If a user is logged in, check if the user is joined.
      if (this.loggedInUser) {
        try {
          this.joined = await AccountService.getJoinState(this.subreddit, this.loggedInUser.authHeader);
        } catch (error) {
          if (error.response && error.response.status === 401) {
            // Invalid token.
            this.$store.commit('logout');
            localStorage.removeItem('user');
          }
          this.loggedInUser = null;
          this.errMsg = 'Failed to fetch the state of joining.';
        }
        this.loading = false;
      }
    }
  },
  methods: {
    createPostClicked() {
      this.$router.push({
        name: 'submit',
        query: {
          subreddit: this.subreddit
        }
      });
    },
    async joinClicked() {
      this.loading = true;
      try {
        if (this.joined) {
          await AccountService.leaveSubreddit(this.subreddit, this.loggedInUser.authHeader);
          this.members--;
        } else {
          await AccountService.joinSubreddit(this.subreddit, this.loggedInUser.authHeader);
          this.members++;
        }
        this.joined = !this.joined;
      } catch (error) {
        if (error.response && error.response.status === 401) {
          // Invalid token.
          this.$store.commit('logout');
          localStorage.removeItem('user');

          this.errMsg = 'Please log in and try again.';
        } else {
          this.errMsg = 'Failed to join. Try again later.';
        }
      }
      this.loading = false;
    }
  }
}
</script>

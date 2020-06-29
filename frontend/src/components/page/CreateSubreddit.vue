<template>
  <div class="row">
    <div class="col q-pa-lg">
      <!-- Create Subreddit Text -->
      <div class="text-h6 text-white">Create a subreddit</div>
      <q-separator color="grey" />

      <!-- Error Indicator -->
      <q-banner v-if="errMsg != ''" class="text-white bg-red q-mt-xs">
          Error: {{errMsg}}.
      </q-banner>

      <!-- Create Subreddit Box -->
      <q-card dark class="q-mt-lg">
        <q-card-section>
          <q-input class="q-mb-md" dark outlined dense standout v-model="subreddit" counter :maxlength="subredditMaxLen" label="Subreddit Name" :disabled="loading" />
          <q-input dark outlined standout v-model="description" type="textarea" placeholder="What is the subreddit about?" label="Description" counter :maxlength="subredditDescriptionMaxLen" :disabled="loading" />
        </q-card-section>

        <q-card-actions class="bg-grey-10 q-pr-lg q-pb-lg" align="right">
          <q-btn @click="cancelClicked" outline class="q-mr-sm" style="color: white; width: 100px" label="Cancel" :disabled="loading" />
          <q-btn @click="createClicked" :disabled="!validSubreddit" style="background: white; color: black; width: 100px" label="Create" :loading="loading" />
        </q-card-actions>
      </q-card>
    </div>

    <!-- Right Panel -->
    <div class="col-3 q-pr-md q-pt-md gt-sm">
      <tips
        title="Subreddit Rules"
        :tips="[
          '1. Do not create offensive or NSFW subreddits, or I will simply reset the database.',
          '2. Subreddit name should contain 1-16 alphanumerical or underscore characters.',
          '3. Rules are rules.',
        ]"
      />
      <advertisement />
    </div>
  </div>
</template>

<script>
import Advertisement from '../rightPanel/Advertisement';
import Tips from '../rightPanel/Tips';
import Limits from '../../limits';
import SubredditService from '../../services/subreddit';

export default {
  mounted() {
    document.title = 'Create Subreddit - YARC';

    // Redirect to the log in page if the user isn't logged in.
    if (!this.$store.state.auth.loggedIn) {
      this.$router.push('/account?login');
    }
  },
  data() {
    return {
      subreddit: '',
      description: '',
      subredditMaxLen: Limits.subredditMaxLen,
      subredditDescriptionMaxLen: Limits.subredditDescriptionMaxLen,
      errMsg: '',
      loading: false
    };
  },
  computed: {
    validSubreddit() {
      if (this.subreddit === '' || this.description === '') {
        return false;
      }

      return true;
    }
  },
  methods: {
    cancelClicked() {
      this.$router.push('/');
    },
    createClicked() {
      this.loading = true;
      
      // Send request.
      SubredditService.create(this.subreddit, this.description, this.$store.state.auth.user.authHeader).then(() => {
        // Successful, redirect to the newly created subreddit.
        this.$router.push('/r/'+this.subreddit);
      }).catch(error => {
        if (error.response.status == 401) {
          this.errMsg = 'your login credentials have expired, please log in again';
          this.$store.commit('logout');
          localStorage.removeItem('user');
        } else {
          this.errMsg = error.response.data.error;
        }
        this.loading = false;
      });
    }
  },
  components: {
    advertisement: Advertisement,
    tips: Tips
  }
}
</script>
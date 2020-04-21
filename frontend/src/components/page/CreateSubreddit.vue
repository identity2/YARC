<template>
  <div class="row">
    <div class="col q-pa-lg">
      <!-- Create Subreddit Text -->
      <div class="text-h6 text-white">Create a subreddit</div>
      <q-separator color="grey" />

      <!-- Create Subreddit Box -->
      <q-card dark class="q-mt-lg">
        <q-card-section>
          <q-input class="q-mb-md" dark outlined dense standout v-model="subreddit" counter :maxlength="subredditMaxLen" label="Subreddit Name" />
          <q-input dark outlined standout v-model="description" type="textarea" placeholder="What is the subreddit about?" label="Description" counter :maxlength="subredditDescriptionMaxLen" />
        </q-card-section>

        <q-card-actions class="bg-grey-10 q-pr-lg q-pb-lg" align="right">
          <q-btn @click="cancelClicked" outline class="q-mr-sm" style="color: white; width: 100px" label="Cancel" />
          <q-btn @click="createClicked" :disabled="!validSubreddit" style="background: white; color: black; width: 100px" label="Create" />
        </q-card-actions>
      </q-card>
    </div>

    <!-- Right Panel -->
    <div class="col-3 q-pr-md q-pt-md gt-sm">
      <tips
        title="Subreddit Rules"
        :tips="[
          '1. Do not create offensive or NSFW subreddits, or I will simply reset the database.',
          '2. The space characters in the subreddit name will be replaced by underscores.',
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

export default {
  mounted() {
    document.title = 'Create Subreddit - YARC';
  },
  data() {
    return {
      subreddit: '',
      description: '',
      subredditMaxLen: Limits.subredditMaxLen,
      subredditDescriptionMaxLen: Limits.subredditDescriptionMaxLen
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

    }
  },
  components: {
    advertisement: Advertisement,
    tips: Tips
  }
}
</script>
<template>
  <div class="column">
    <div class="col">
      <q-btn @click.stop="upvoteClicked" round flat size="xs" icon="arrow_upward" :style="voteStatus === 'up' ? 'color: red' : ''" :disabled="voteStatus === null" />
    </div>
    <div class="col text-center">
      {{currVotes}}
    </div>
    <div class="col">
      <q-btn @click.stop="downvoteClicked" round flat size="xs" icon="arrow_downward" :style="voteStatus === 'down' ? 'color: cyan' : ''" :disabled="voteStatus === null" />
    </div>
  </div>
</template>

<script>
import KarmaService from '../../services/karma';

export default {
  props: {
    votes: Number,
    articleID: String
  },
  data() {
    return {
      currVotes: 0,
      voteStatus: null
    };
  },
  async mounted() {
    this.currVotes = this.votes;

    // Get the voting status.
    try {
      this.voteStatus = await KarmaService.getArticleVote(this.articleID, this.$store.state.auth.authHeader);
    } catch (error) {
      if (error.response && error.response.status === 401) {
        this.$store.commit('logout');
        localStorage.removeItem('user');
      }
    }
  },
  methods: {
    upvoteClicked() {
      let args = {
        up: {action: 'cancel', delta: -1},
        down: {action: 'up', delta: 2},
        cancel: {action: 'up', delta: 1}
      };
      
      this.sendVoteRequest(args[this.voteStatus]);
    },
    downvoteClicked() {
      let args = {
        up: {action: 'down', delta: -2},
        down: {action: 'cancel', delta: 1},
        cancel: {action: 'down', delta: -1}
      };

      this.sendVoteRequest(args[this.voteStatus]);
    },
    async sendVoteRequest({action, delta}) {
      this.voteStatus = null;  // Indicating that it is loading.

      try {
        await KarmaService.voteArticle(this.articleID, action, this.$store.state.auth.authHeader);
        this.currVotes += delta;
        this.voteStatus = action;
      } catch (error) {
        if (error.response && error.response.status === 401) {
          this.$store.commit('logout');
          localStorage.removeItem('user');
        }
      }
    }
  }
}
</script>
<template>
  <div class="q-mt-lg q-mb-md">
    <q-item-section>
      <div class="row">
        <!-- Upvote, Downvote -->
        <div class="column q-mr-md q-mt-sm">
          <q-btn
            @click.stop="upvoteClicked"
            round
            flat
            size="xs"
            icon="arrow_upward"
            :style="voteStatus === 'up' ? 'color: red' : ''"
            :disabled="voteStatus === null"
          />
          <q-btn
            @click.stop="downvoteClicked"
            round
            flat
            size="xs"
            icon="arrow_downward"
            :style="voteStatus === 'down' ? 'color: cyan' : ''"
            :disabled="voteStatus === null"
          />
        </div>
        
        <div class="col">
          <div class="text-caption text-grey">
            <router-link :to="'/u/' + comment.postedBy">
              <span class="text-grey-5">u/{{comment.postedBy}}</span>
            </router-link>
            <span>．{{comment.points}} points．{{comment.postedTime | formatDate}}</span>
            <q-btn v-if="!deleting && !editMode && postedByCurrUser" class="q-ml-md" @click="editMode = true; deleteConfirming = false" dense flat size="xs" icon="create" label="Edit" />
            <q-btn v-if="!deleting && !editMode && postedByCurrUser" class="q-ml-md" @click="deleteConfirming = !deleteConfirming" dense falt size="xs" icon="delete" label="Delete" />
          </div>

          <!-- Comment Text -->
          <div v-if="!editMode" class="text-white q-mt-xs q-mr-sm" style="white-space: pre-wrap">
            {{comment.body}}
          </div>

          <div v-if="editMode" class="q-mt-sm q-mr-sm q-mb-sm">
            <div class="q-ml-xs q-mb-sm">Editing comment:</div>
            <q-input dark outlined standout v-model="currBody" type="textarea" counter :maxlength="commentMaxLen" :disable="saving" />
            <q-btn class="q-mr-md" @click="saveClicked" style="background: white; color: black" label="Save" :loading="saving" />
            <q-btn outline @click="cancelClicked" style="color: red" label="Cancel" :disabled="saving" />
          </div>

          <div v-if="deleteConfirming" class="q-mt-sm q-mr-sm q-mb-sm">
            <q-btn class="q-mr-md" @click="deleteConfirmed" style="background: red; color: white" label="Confirm Delete" :loading="deleting" />
            <q-btn outline @click="deleteConfirming = false" style="color: red" label="Cancel" :disabled="deleting" />
          </div>
        </div>

        <!-- Error Message  -->
        <div v-if="errMsg != ''" class="text-red q-mr-lg">Error: {{errMsg}}.</div>
      </div>
    </q-item-section>
  </div>
</template>

<script>
import Limits from '../../limits';
import KarmaService from '../../services/karma';
import CommentService from '../../services/comment';

export default {
  props: {
    comment: Object
  },
  data() {
    return {
      editMode: false,
      commentMaxLen: Limits.commentMaxLen,
      currBody: this.comment.body,
      voteStatus: null,  // Will be either 'up', 'down' or 'cancel' when finish loading.
      errMsg: '',
      saving: false,
      deleteConfirming: false,
      deleting: false
    };
  },
  computed: {
    postedByCurrUser() {
      return this.$store.state.auth && this.$store.state.auth.username === this.comment.postedBy;
    }
  },
  async mounted() {
    // Get the voting status.
    try {
      this.voteStatus = await KarmaService.getCommentVote(this.comment.commentID, this.$store.state.auth.authHeader);
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
        await KarmaService.voteComment(this.comment.commentID, action, this.$store.state.auth.authHeader);
        this.comment.points += delta;
        this.voteStatus = action;
      } catch (error) {
        if (error.response && error.response.status === 401) {
          this.$store.commit('logout');
          localStorage.removeItem('user');
          this.errMsg = 'login credentials expired, please relogin to vote';
        } else {
          this.errMsg = 'failed to vote the comment, please refresh page and try again';
        }
      }
    },
    async saveClicked() {
      this.saving = true;

      try {
        await CommentService.modify(this.comment.commentID, this.currBody, this.$store.state.auth.authHeader);
        
        // Inform the parent.
        this.$emit('commentEdited', this.comment.commentID, this.currBody);

        this.saving = false;
        this.editMode = false;
      } catch (error) {
        if (error.response && error.response.status === 401) {
          this.$store.commit('logout');
          localStorage.removeItem('user');
          this.errMsg = 'login credentials expired, please relogin to edit';
        } else {
          this.errMsg = 'failed to save the comment, please try again';
        }
      }
    },
    cancelClicked() {
      this.editMode = false;
      this.currBody = this.comment.body; // Change to the original comment text.
    },
    async deleteConfirmed() {
      this.deleting = true;
      try {
        await CommentService.remove(this.comment.commentID, this.$store.state.auth.authHeader);
        
        // Notify the parent to remove the comment rendered.
        this.$emit('commentDeleted');
      } catch (error) {
        if (error.response && error.response.status === 401) {
          this.$store.commit('logout');
          localStorage.removeItem('user');
          this.errMsg = 'login credentials expired, please relogin to delete';
        } else {
          this.errMsg = 'failed to delete the comment, please try again';
        }
        this.deleting = false;
      }
    }
  }
}
</script>

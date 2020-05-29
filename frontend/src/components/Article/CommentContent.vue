<template>
  <div class="q-mt-lg q-mb-md">
    <q-item-section>
      <div class="row">
        <!-- Upvote, Downvote -->
        <div class="column q-mr-md q-mt-sm">
          <q-btn @click.stop="upvoteClicked" round flat size="xs" icon="arrow_upward" />
          <q-btn @click.stop="downvoteClicked" round flat size="xs" icon="arrow_downward" />
        </div>
        
        <div class="col">
          <div class="text-caption text-grey">
            <router-link :to="'/u/' + commentedBy">
              <span class="text-grey-5">u/{{commentedBy}}</span>
            </router-link>
            <span>．{{points}} points．{{commentedDate}}</span>
          </div>

          <!-- Comment Text -->
          <div v-if="!editMode" class="text-white q-mt-xs q-mr-sm">
            {{commentBody}}
          </div>

          <div v-if="editMode" class="q-mt-sm q-mr-sm q-mb-md">
            <div class="q-ml-xs q-mb-sm">Editing comment:</div>
            <q-input dark outlined standout v-model="commentBody" type="textarea" counter :maxlength="commentMaxLen" />
            <q-btn class="q-mr-md" @click="saveClicked" style="background: white; color: black" label="Save" />
            <q-btn outline @click="cancelClicked" style="color: red" label="Cancel" />
          </div>

          <div class="text-caption text-grey q-mt-xs">
            <q-btn @click="editMode = true" dense flat size="xs" icon="create" label="Edit" />
            <q-btn @click="deleteClicked" class="q-ml-sm" dense flat size="xs" icon="delete" label="Delete" />
          </div>
        </div>

      </div>
    </q-item-section>
  </div>
</template>

<script>
import Limits from '../../limits';

export default {
  props: {
    commentedBy: String,
    points: Number,
    commentedDate: String,
    commentText: String
  },
  data() {
    return {
      editMode: false,
      commentMaxLen: Limits.commentMaxLen,
      commentBody: this.commentText
    };
  },
  methods: {
    upvoteClicked() {

    },
    downvoteClicked() {
      
    },
    deleteClicked() {

    },
    saveClicked() {

    },
    cancelClicked() {
      this.editMode = false;
      this.commentBody = this.commentText; // Store to the original comment text.
    }
  }
}
</script>

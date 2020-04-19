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
            <router-link :to="'/user/' + commentedBy">
              <span class="text-grey-5">u/{{commentedBy}}</span>
            </router-link>
            <span>．{{points}} points．{{commentedDate}}</span>
          </div>

          <div class="text-white q-mt-xs q-mr-sm">
            {{commentText}}
          </div>
        </div>

      </div>
    </q-item-section>
  </div>
</template>

<script>
export default {
  props: {
    commentedBy: String,
    points: Number,
    commentedDate: String,
    commentText: String
  },
  data() {
    return {
      
    };
  },
  methods: {
    upvoteClicked() {

    },
    downvoteClicked() {
      
    }
  }
}
</script>

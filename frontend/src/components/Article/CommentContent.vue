<template>
  <div class="q-mt-lg q-mb-md">
    <q-item-section>
      <div class="row">
        <div class="column q-mr-md q-mt-sm">
          <q-btn @click.stop="upvoteClicked" round flat size="xs" icon="arrow_upward" />
          <q-btn @click.stop="downvoteClicked" round flat size="xs" icon="arrow_downward" />
        </div>
        
        <div class="col">
          <div class="text-caption text-grey">
            <!-- Upvote, Downvote -->

            <a class="text-grey-5" :href="'/user/' + commentedBy">u/{{commentedBy}}</a>
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

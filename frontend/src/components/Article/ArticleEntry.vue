<template>
  <q-item clickable>
    <!-- Upvote, Downvote -->
    <q-item-section thumbnail>
      <voter :votes="votes"></voter>
    </q-item-section>

    <!-- Thumbnail -->
    <q-item-section thumbnail>
      <q-avatar color="grey-9" rounded size="65px">
        <q-icon v-if="postType != 'image'" :name="thumbnailIcon"></q-icon>
        <img v-if="postType == 'image'" :src="thumbnailUrl" />
      </q-avatar>
    </q-item-section>

    <!-- Article Title and info -->
    <q-item-section class="col">
      <q-item-label lines="1">
        <span class="text-h6">{{title}}</span>
      </q-item-label>

      <q-item-label lines="1" class="text-grey">
        <span>r/{{subreddit}}</span>
        <span>．Posted by u/{{postedBy}} on {{postedDate}}</span>
      </q-item-label>

      <q-item-label lines="1" class="text-grey text-weight-bold">
        <q-btn flat round @click="expandClicked" :icon="expanded ? 'expand_less' : 'expand_more'" size='xs'></q-btn>
        | {{comments}} Comments 
        | Share 
        | (Save) 
        | (Edit) 
        | (Delete)
      </q-item-label>
    </q-item-section>
  </q-item>
</template>

<script>
import Voter from './Voter';

export default {
  props: {
    title: String,
    postType: String, // text, image, link
    thumbnailUrl: String,
    votes: Number,
    postedBy: String,
    comments: Number,
    subreddit: String,
    postedDate: String
  },
  data() {
    return {
      expanded: false
    };
  },
  computed: {
    thumbnailIcon() {
      if (this.postType == 'text') {
        return 'chat';
      } else if (this.postType == 'link') {
        return 'link';
      }

      // No icon for image.
      return '';
    }
  },
  methods: {
    expandClicked() {
      this.expanded = !this.expanded;

      // TODO: Expand the article.
    }
  },
  components: {
    voter: Voter
  }
}
</script>

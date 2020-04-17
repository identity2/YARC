<template>
  <q-item class="column" clickable>
    <div class="row">
      <!-- Upvote, Downvote -->
      <q-item-section thumbnail>
        <voter :votes="votes"></voter>
      </q-item-section>

      <!-- Thumbnail -->
      <q-item-section thumbnail>
        <q-avatar color="grey-9" rounded size="65px">
          <q-icon v-if="postType !== 'image'" :name="thumbnailIcon"></q-icon>
          <img v-if="postType === 'image'" :src="imageUrl" />
        </q-avatar>
      </q-item-section>

      <!-- Article Title and info -->
      <q-item-section class="col">
        <q-item-label lines="1">
          <span class="text-h6">{{title}}</span>
        </q-item-label>

        <q-item-label lines="1" class="text-grey">
          <a :href="'r/' + subreddit">r/{{subreddit}}</a>
          <span>．Posted by </span>
          <a :href="'user/' + postedBy">u/{{postedBy}}</a>
          <span> on {{postedDate}}</span>
        </q-item-label>

        <!-- Expand, Comments, Share, Save, Edit, Delete -->
        <q-item-label lines="1" class="text-grey text-weight-bold">
          <q-btn v-if="postType !== 'link'" flat round @click.stop="expandClicked" :icon="expanded ? 'expand_less' : 'expand_more'" size="xs" />
          <q-btn v-if="postType === 'link'" flat round @click.stop="openLinkInNewTab" size="xs" icon="open_in_new" />
          |
          <q-btn class="q-ml-sm" dense flat size="xs" icon="chat_bubble" :label="'' + comments + ' Comments'" />

          <q-btn @click.stop="" class="q-ml-sm" dense flat size="xs" icon="share" label="Share" />

          <q-btn @click.stop="" class="q-ml-sm" dense flat size="xs" icon="save_alt" label="(Save)" />

          <q-btn @click.stop="" class="q-ml-sm" dense flat size="xs" icon="create" label="(Edit)" />

          <q-btn @click.stop="" class="q-ml-sm" dense flat size="xs" icon="delete" label="(Delete)" />
        </q-item-label>
      </q-item-section>
    </div>

    <!-- Expanded article content -->
    <div v-if="expanded" class="row q-pt-lg q-pl-lg q-pr-lg q-pb-sm">
      <img v-if="postType === 'image'" :src="imageUrl" />
      <p v-if="postType === 'text'">{{textBody}}</p>
    </div>
  </q-item>
</template>

<script>
import Voter from './Voter';

export default {
  props: {
    title: String,
    postType: String, // text, image, link
    textBody: String,
    imageUrl: String,
    linkUrl: String,
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
      if (this.postType === 'text') {
        return 'chat';
      } else if (this.postType === 'link') {
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
    },
    openLinkInNewTab() {
      window.open(this.linkUrl, '_blank');
    }
  },
  components: {
    voter: Voter
  }
}
</script>
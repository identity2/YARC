<template>
  <div class="col">
    <!-- Sort, Create Post -->
    <div v-if="canCreatePost" class="q-pt-md q-pl-md q-pr-md text-white">
      <q-list dark bordered class="bg-grey-10 rounded-borders">
        <q-item class="q-mt-sm q-mb-sm">          
            <q-btn-dropdown color="blue" label="Sort">
              <q-list>
                <q-item clickable v-close-popup>Hot</q-item>
                <q-item clickable v-close-popup>New</q-item>
                <q-item clickable v-close-popup>Old</q-item>
              </q-list>
            </q-btn-dropdown>

            <template v-if="canCreatePost">
              <q-input @click="createPost('text')" class="col q-ml-lg q-mr-sm" dark clearable outlined dense standout label="Create Post" />
              <q-btn @click="createPost('image')" flat round color="grey" icon="insert_photo">
                <q-tooltip>Create Image Post</q-tooltip>
              </q-btn>
              <q-btn @click="createPost('link')" flat round color="grey" icon="insert_link">
                <q-tooltip>Create Link Post</q-tooltip>
              </q-btn>
            </template>
        </q-item>
      </q-list>
    </div>

    <!-- Article Lists -->
    <div class="q-pa-md text-white">
      <q-list dark bordered separator class="bg-grey-10 rounded-borders">
        <article-entry
          v-for="article in articles"
          :key="article.id"
          :title="article.title"
          :postType="article.postType"
          :imageUrl="article.imageUrl"
          :linkUrl="article.linkUrl"
          :textBody="article.textBody"
          :points="article.points"
          :postedBy="article.postedBy"
          :comments="article.comments"
          :subreddit="article.subreddit"
          :postedDate="article.postedDate"
          @click.native="articleClicked(article)"
        ></article-entry>
      </q-list>
    </div>
  </div>
</template>

<script>
import ArticleEntry from '../article/ArticleEntry';

export default {
  props: {
    articles: Array,
    subreddit: {
      type: String,
      default: ''
    },
    canCreatePost: {
      type: Boolean,
      default: true
    }
  },
  methods: {
    articleClicked(article) {
      this.$router.push({
        path: '/r/' + article.subreddit + '/p/' + article.id
      });
    },
    createPost(type) {
      this.$router.push({
        name: 'submit',
        query: {
          subreddit: this.subreddit,
          postType: type
        }
      });
    }
  },
  components: {
    articleEntry: ArticleEntry
  }
}
</script>
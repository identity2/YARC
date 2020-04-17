<template>
  <div class="row">
    <div class="col">
      <!-- Sort, Create Post -->
      <div class="q-pt-md q-pl-md q-pr-md text-white">
        <q-list dark bordered class="bg-grey-10 rounded-borders">
          <q-item class="q-mt-sm q-mb-sm">          
              <q-btn-dropdown color="blue" label="Sort">
                <q-list>
                  <q-item clickable v-close-popup>Hot</q-item>
                  <q-item clickable v-close-popup>New</q-item>
                  <q-item clickable v-close-popup>Old</q-item>
                </q-list>
              </q-btn-dropdown>

              <q-input @click="createPost('text')" class="col q-ml-lg q-mr-sm" dark clearable outlined dense standout label="Create Post" />
              <q-btn @click="createPost('image')" flat round color="grey" icon="insert_photo">
                <q-tooltip>Create Image Post</q-tooltip>
              </q-btn>
              <q-btn @click="createPost('link')" flat round color="grey" icon="insert_link">
                <q-tooltip>Create Link Post</q-tooltip>
              </q-btn>
          </q-item>
        </q-list>
      </div>

      <!-- Article Lists -->
      <div class="q-pa-md text-white">
        <q-list dark bordered separator class="bg-grey-10 rounded-borders">
          <article-entry
            v-for="article in loadedArticles"
            :key="article.id"
            :title="article.title"
            :postType="article.postType"
            :imageUrl="article.imageUrl"
            :linkUrl="article.linkUrl"
            :textBody="article.textBody"
            :votes="article.votes"
            :postedBy="article.postedBy"
            :comments="article.comments"
            :subreddit="article.subreddit"
            :postedDate="article.postedDate"
            @click.native="articleClicked(article)"
          ></article-entry>
        </q-list>
      </div>
    </div>

    <!-- Right Panel -->
    <div class="col-3 q-pr-md q-pt-md gt-sm">
      <about-subreddit v-if="currSubreddit != ''"></about-subreddit>
      <trending-subreddits></trending-subreddits>
      <advertisement></advertisement>
    </div>

    <!-- Article Overlay -->
    <q-dialog v-model="articleOverlay" full-height full-width>
      <q-card class="bg-black text-white full-height">
        <q-bar class="bg-grey-10 q-pa-md">
          <q-icon :name="postTypeToIcon(overlayArticle.postType)" />
          <div class="q-ma-sm">{{overlayArticle.title}}</div>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>CLOSE</q-btn>
        </q-bar>
        <q-card-section>
          <div class="text-h6">
            Article Component Here!!
          </div>
        </q-card-section>
      </q-card>
    </q-dialog>
  </div>
</template>

<script>
import ArticleEntry from '../article/ArticleEntry';
import AboutSubreddit from '../rightPanel/AboutSubreddit';
import Advertisement from '../rightPanel/Advertisement';
import TrendingSubreddits from '../rightPanel/TrendingSubreddits';
import mockArticles from '../../mock_data/mock_articles';

export default {
  components: {
    articleEntry: ArticleEntry,
    aboutSubreddit: AboutSubreddit,
    advertisement: Advertisement,
    trendingSubreddits: TrendingSubreddits
  },
  data() {
    return {
      currSubreddit: '',
      articleOverlay: false,
      overlayArticle: {},
      loadedArticles: mockArticles
    };
  },
  methods: {
    articleClicked(article) {
      this.overlayArticle = article;
      this.articleOverlay = true;
    },
    postTypeToIcon(postType) {
      if (postType === 'text') return 'chat';
      if (postType === 'image') return 'image';
      if (postType === 'link') return 'link';
      return 'error';
    },
    createPost(type) {
      this.$router.push({
        name: 'submit',
        query: {
          subreddit: this.currSubreddit,
          postType: type
        }
      });
    }
  }
}
</script>
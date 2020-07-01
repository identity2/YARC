<template>
  <div class="col">
    <!-- Sort, Create Post -->
    <div class="q-pt-md q-pl-md q-pr-md text-white">
      <q-list dark bordered class="bg-grey-10 rounded-borders">
        <q-item class="q-mt-sm q-mb-sm">
          <div class="q-gutter-sm">
            <q-radio dark v-model="sortBy" val="hot" label="Hot" />
            <q-radio dark v-model="sortBy" val="new" label="New" />
            <q-radio dark v-model="sortBy" val="old" label="Old" />
          </div>

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
        <q-infinite-scroll @load="loadMoreArticles" :offset="250" ref="infiniteScroll">
          <article-entry
            v-for="article in articles"
            :key="article.articleID"
            :article="article"
            @click.native="articleClicked(article)"
          />

          <template v-slot:loading>
            <div class="row justify-center q-my-md">
              <q-spinner-dots color="white" size="40px" />
            </div>
          </template> 
        </q-infinite-scroll>

        <!-- Loading Indicator -->
        <div v-if="loading">
          <q-item v-for="i in 15" :key="i">
              <q-item-section avatar>
                <q-skeleton type="QAvatar" />
              </q-item-section>
              <q-item-section>
                <q-item-label><q-skeleton type="text" /></q-item-label>
              </q-item-section>
          </q-item>
        </div>
        
        <!-- Error Indicator -->
        <q-banner v-if="errOccurred" class="text-white bg-red">
            Error fetching the article list, please try again.
        </q-banner>
        
        <!-- Empty Indicator -->
        <q-item v-if="articles.length == 0 && !errOccurred" class="column">
          <div class="q-ma-lg">
            <div class="text-h3 q-mb-md">Wow, such empty!</div>
            <div class="text-h6">{{emptyText}}</div>
          </div>
        </q-item>

      </q-list>
    </div>
  </div>
</template>

<script>
import ArticleEntry from '../article/ArticleEntry';
import ArticleService from '../../services/article';

export default {
  props: {
    emptyText: {
      type: String,
      default: ''
    },
    criterion: {
      type: String,
      default: ''
    },
    criterionKey: {
      type: String,
      default: ''
    },
    subreddit: {
      type: String,
      default: ''
    },
    canCreatePost: {
      type: Boolean,
      default: true
    }
  },
  data() {
    return {
      articles: [],
      sortBy: "hot",
      articlesPerRequest: 10,
      errOccurred: false,
      loading: true
    };
  },
  watch: {
    sortBy() {
      this.reloadArticles();
    },
    criterion() {
      this.reloadArticles();
    },
    criterionKey() {
      this.reloadArticles();
    }
  },
  mounted() {
    this.reloadArticles();
  },
  methods: {
    articleClicked(article) {
      this.$router.push({
        path: '/r/' + article.subreddit + '/p/' + article.articleID
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
    },
    async fetchArticleLists(after) {
      let auth = this.$store.state.auth;
      let fetchedArticles = await ArticleService.getList(this.sortBy, after, this.articlesPerRequest, this.criterion, this.criterionKey, auth ? auth.authHeader : null);
      
      // Check if it is the end of the article list. If it is, stop the scroll.
      if (fetchedArticles.length < this.articlesPerRequest) {
        this.$refs.infiniteScroll.stop();
      }

      return fetchedArticles;
    },
    async reloadArticles() {
      // Refetch the articles since the sort criterion has changed.
      this.loading = true;
      this.articles = [];
      try {
        this.$refs.infiniteScroll.resume(); // Make the infinite scroller work again.
        this.articles = await this.fetchArticleLists("");
        this.errOccurred = false;
      } catch {
        this.errOccurred = true;
      }
      this.loading = false;
    },
    loadMoreArticles(_, done) {
      // Fetch more articles.
      let after = this.articles.length == 0 ? "" : this.articles[this.articles.length-1].articleID;
      this.fetchArticleLists(after).then(fetchedArticles => {
          for (const art of fetchedArticles) {
            this.articles.push(art);
          }
          done();
      }).catch(() => {
        this.errOccurred = true;
        done(true);
      });
    }
  },
  components: {
    articleEntry: ArticleEntry
  }
}
</script>
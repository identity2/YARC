<template>
  <div>
    <!-- Title -->
    <div class="text-h6 text-white q-mt-md q-ml-lg">
      Search Results of <span class="text-blue">{{query}}</span>
    </div>

    <!-- Error Indicator -->
    <q-banner v-if="errOccurred" class="text-white bg-red">
        Error fetching the search result, please try again.
    </q-banner>

    <div class="row">
      <div class="col">
        <!-- Subreddits -->
        <div class="text-h6 text-white q-mt-lg q-ml-lg">Subreddits</div>
        <div class="q-pa-md">
          <q-list dark bordered class="bg-grey-10 rounded-borders q-pt-md q-pb-md q-pl-sm q-pr-sm">
            
            <!-- Empty Indicator -->
            <q-item v-if="result.subreddits.length === 0">
              <div class="q-ma-lg">
                <div class="text-h6">No result. Try a different search keyword!</div>
              </div>
            </q-item>

            <!-- Loop through subreddits -->
            <q-item v-for="sub in result.subreddits" :key="sub.name">
              <q-item-section thumbnail top>
                <q-avatar class="q-ml-md" size="md" color="blue" text-color="white">
                  {{sub.name[0].toUpperCase()}}
                </q-avatar>
              </q-item-section>

              <q-item-section>
                <div class="row">
                  <div class="col-2 q-mr-xl">
                    <span class="text-white">r/{{sub.name}}</span>
                    <q-item-label caption lines="1">Members: {{sub.members}}</q-item-label>
                  </div>

                  <div class="col gt-sm q-mr-xl">
                    <span class="text-grey">
                      {{sub.description}}
                    </span>
                  </div>
                </div>
              </q-item-section>

              <q-item-section side>
                <q-btn @click="$router.push('/r/' + sub.name)" class="q-ml-md" dense style="background: white; color: black; width: 100px" label="Visit" />
              </q-item-section>
            </q-item>
          </q-list>
        </div>

        <!-- Articles -->
        <div class="text-h6 text-white q-mt-sm q-ml-lg">Posts</div>
        <div class="q-pa-md text-white">
          <q-list dark bordered separator class="bg-grey-10 rounded-borders">
            <!-- Empty Indicator -->
            <q-item v-if="result.articles.length === 0">
              <div class="q-ma-lg">
                <div class="text-h6">No result. Try a different search keyword!</div>
              </div>
            </q-item>

            <article-entry
              v-for="article in result.articles"
              :key="article.articleID"
              :title="article.title"
              :postType="article.type"
              :imageUrl="article.type == 'image' ? article.body : ''"
              :linkUrl="article.type == 'link' ? article.body : ''"
              :textBody="article.type == 'text' ? article.body : ''"
              :points="article.points"
              :postedBy="article.postedBy"
              :comments="article.comments"
              :subreddit="article.subreddit"
              :postedDate="article.postedTime"
              @click.native="articleClicked(article)"
            />
          </q-list>
        </div>
      </div>

      <!-- Right Panel -->
      <div class="col-3 q-pr-md q-pt-md gt-sm">
        <!-- This div is for aligning -->
        <div class="text-h6 q-mt-lg q-ml-lg">hIdDEn tExT</div>

        <div class="">
          <new-subreddit />
          <trending-subreddits />
          <advertisement />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NewSubreddit from '../rightPanel/NewSubreddit';
import TrendingSubreddits from '../rightPanel/TrendingSubreddits';
import Advertisement from '../rightPanel/Advertisement';
import ArticleEntry from '../article/ArticleEntry';
import SearchService from '../../services/search';

export default {
  data() {
    return {
      query: '',
      result: {
        subreddits: [],
        articles: []
      },
      errOccurred: false
    };
  },
  mounted() {
    this.reloadPage(this.$route.query.q);
  },
  watch: {
    $route(newVal) {
      this.reloadPage(newVal.query.q);
    }
  },
  methods: {
    async reloadPage(query) {
      window.scrollTo(0, 0);
      this.query = query;
      document.title = 'YARC search: ' + this.query;

      // Fetch the search results.
      try {
        this.result = await SearchService.search(this.query);
      } catch {
        this.errOccurred = true;
      }
    },
    articleClicked(article) {
      this.$router.push({
        path: '/r/' + article.subreddit + '/p/' + article.id
      });
    }
  },
  components: {
    newSubreddit: NewSubreddit,
    trendingSubreddits: TrendingSubreddits,
    advertisement: Advertisement,
    articleEntry: ArticleEntry
  }
}
</script>
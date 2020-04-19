<template>
  <div>
    <!-- Title -->
    <div class="text-h6 text-white q-mt-md q-ml-lg">
      Welcome to <span class="text-blue">r/{{subreddit}}</span>!
    </div>

    <div class="row">
      <list-of-articles :articles="articles" :subreddit="subreddit" />

      <!-- Right Panel -->
      <div class="col-3 q-pr-md q-pt-md gt-sm">
        <about-subreddit :subreddit="subreddit" :members="members" :description="description" />
        <trending-subreddits />
        <advertisement />
      </div>
    </div>
  </div>
</template>

<script>
import AboutSubreddit from '../rightPanel/AboutSubreddit';
import TrendingSubreddits from '../rightPanel/TrendingSubreddits';
import Advertisement from '../rightPanel/Advertisement';
import ListOfArticles from '../article/ListOfArticles';
import mock_articles from '../../mock_data/mock_articles';

export default {
  data() {
    return {
      subreddit: '',
      members: 20,
      description: 'This is the description of this subreddit!',
      articles: mock_articles.slice(1, -1)
    };
  },
  watch: {
    $route(newVal) {
      this.reloadPage(newVal.params.subreddit);
    }
  },
  mounted() {
    this.reloadPage(this.$route.params.subreddit);
  },
  methods: {
    reloadPage(subreddit) {
      window.scrollTo(0, 0);
      this.subreddit = subreddit;

      // Load the articles.
    }
  },
  components: {
    listOfArticles: ListOfArticles,
    aboutSubreddit: AboutSubreddit,
    trendingSubreddits: TrendingSubreddits,
    advertisement: Advertisement
  }
}
</script>
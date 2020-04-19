<template>
  <div class="row">
    <list-of-articles :articles="articles" />

    <!-- Right Panel -->
    <div class="col-3 q-pr-md q-pt-md gt-sm">
      <new-subreddit />
      <trending-subreddits />
      <advertisement />
    </div>
  </div>
</template>

<script>
import NewSubreddit from '../rightPanel/NewSubreddit';
import TrendingSubreddits from '../rightPanel/TrendingSubreddits';
import Advertisement from '../rightPanel/Advertisement';
import ListOfArticles from '../article/ListOfArticles';
import mock_articles from '../../mock_data/mock_articles';

export default {
  data() {
    return {
      articles: mock_articles
    };
  },
  mounted() {
    window.scrollTo(0, 0);
  },
  components: {
    listOfArticles: ListOfArticles,
    newSubreddit: NewSubreddit,
    trendingSubreddits: TrendingSubreddits,
    advertisement: Advertisement
  }
}
</script>
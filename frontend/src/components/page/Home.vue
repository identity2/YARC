<template>
  <div class="row">
    <list-of-articles :criterion="''" :criterionKey="''" :emptyText="'There are currently no articles on YARC :('" />

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

export default {
  mounted() {
    window.scrollTo(0, 0);
    document.title = 'YARC';
  },
  components: {
    listOfArticles: ListOfArticles,
    newSubreddit: NewSubreddit,
    trendingSubreddits: TrendingSubreddits,
    advertisement: Advertisement
  }
}
</script>
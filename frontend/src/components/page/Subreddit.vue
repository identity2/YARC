<template>
  <div>
    <!-- Title -->
    <div class="text-h6 text-white q-mt-md q-ml-lg">
      Welcome to <span class="text-blue">r/{{subInfo.name}}</span>!
    </div>

    <!-- Error Indicator -->
    <q-banner v-if="errOccurred" class="text-white bg-red">
        Error: failed to load the subreddit, please try again.
    </q-banner>

    <div class="row">
      <list-of-articles :criterion="'sub'" :criterionKey="subInfo.name" :subreddit="subInfo.name" />

      <!-- Right Panel -->
      <div class="col-3 q-pr-md q-pt-md gt-sm">
        <about-subreddit :subreddit="subInfo.name" :members="subInfo.members" :description="subInfo.description" />
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
import SubredditService from '../../services/subreddit';

export default {
  data() {
    return {
      subInfo: {name: '', members: 0, description: ''},
      errOccurred: false
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
    async reloadPage(subreddit) {
      window.scrollTo(0, 0);
      this.subInfo.name = subreddit;
      document.title = 'r/' + subreddit + ' - YARC';
      
      // Load the subreddit information.
      try {
        this.subInfo = await SubredditService.get(subreddit);
      } catch (error) {
        if (error.response.status === 404) {
          this.$router.replace({name: 'page not found'});
        } else {
          this.errOccurred = true;
        }
      }
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
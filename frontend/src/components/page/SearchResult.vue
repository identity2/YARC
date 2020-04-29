<template>
  <div>
    <!-- Title -->
    <div class="text-h6 text-white q-mt-md q-ml-lg">
      Search Results of <span class="text-blue">{{query}}</span>
    </div>

    <div class="row">
      <div class="col">
        <!-- Subreddits -->
        <div class="text-h6 text-white q-mt-lg q-ml-lg">Subreddits</div>
        <div class="q-pa-md">
          <q-list dark bordered class="bg-grey-10 rounded-borders q-pt-md q-pb-md q-pl-sm q-pr-sm">
            
            <!-- Loop through subreddits -->
            <q-item v-for="sub in subreddits" :key="sub.name">
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

              <q-item-actions align="right">
                <q-btn @click="$router.push('/r/' + sub.name)" class="q-ml-md" dense style="background: white; color: black; width: 100px" label="Visit" />
              </q-item-actions>
            </q-item>
          </q-list>
        </div>

        <!-- Articles -->
        <div class="text-h6 text-white q-mt-sm q-ml-lg">Posts</div>
        <list-of-articles :articles="articles" :canCreatePost="false" />
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
import ListOfArticles from '../article/ListOfArticles';
import mock_articles from '../../mock_data/mock_articles';

export default {
  data() {
    return {
      articles: mock_articles,
      subreddits: [
        {
          name: 'theclashcirclejerk',
          members: 1300,
          description: 'This is a subreddit for the kind-hearted, the smart, and the enlightened!'
        },
        {
          name: 'programmerhumor',
          members: 520,
          description: 'A place for nerdy programmers.'
        },
        {
          name: 'golang',
          members: 39,
          description: 'Is good language!'
        },
        {
          name: 'test',
          members: 20,
          description: ''
        }
      ],
      query: ''
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
    reloadPage(query) {
      window.scrollTo(0, 0);
      this.query = query;
      document.title = 'YARC search: ' + this.query;
    }
  },
  components: {
    listOfArticles: ListOfArticles,
    newSubreddit: NewSubreddit,
    trendingSubreddits: TrendingSubreddits,
    advertisement: Advertisement
  }
}
</script>
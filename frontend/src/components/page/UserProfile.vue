<template>
  <div>
    <!-- Title -->
    <div class="text-h6 text-white q-mt-md q-ml-lg">
      <span class="text-blue">r/{{username}}</span>'s profile.
    </div>

    <q-card dark class="q-mt-lg">
      <q-tabs v-model="viewList">
        <q-tab name="post" label="Posts" />
        <q-tab name="comment" label="Comments" />
        <q-tab name="saved" label="Saved Posts" />
      </q-tabs>
    </q-card>

    <div class="row">
      <list-of-articles
        v-if="viewList === 'post'"
        :articles="posts"
        :canCreatePost="false"
      />

      <list-of-comments
        v-if="viewList === 'comment'"
        :comments="comments"
      />

      <list-of-articles
        v-if="viewList === 'saved'"
        :articles="savedPosts"
        :canCreatePost="false"
      />

      <!-- Right Panel -->
      <div class="col-3 q-pr-md q-pt-md gt-sm">
        <about-user
          :username="username"
          :karma="karma"
          :joined="joined"
        />
        <advertisement />
      </div>
    </div>
  </div>
</template>

<script>
import ListOfArticles from '../article/ListOfArticles';
import ListOfComments from '../article/ListOfComments';
import Advertisement from '../rightPanel/Advertisement';
import AboutUser from '../rightPanel/AboutUser';
import mock_articles from '../../mock_data/mock_articles';
import mock_comments from '../../mock_data/mock_comments';

export default {
  data() {
    return {
      posts: mock_articles,
      comments: mock_comments,
      savedPosts: mock_articles.slice(1, -1),
      username: 'carbon_cop',
      karma: 420,
      joined: Date(),
      viewList: 'post'
    };
  },
  mounted() {
    this.reloadPage(this.$route.params.username);
  },
  watch: {
    $route(newVal) {
      this.reloadPage(newVal.params.username);
    }
  },
  methods: {
    reloadPage(username) {
      this.username = username;
    // TODO: Load user data.
    }
  },
  components: {
    advertisement: Advertisement,
    listOfArticles: ListOfArticles,
    listOfComments: ListOfComments,
    aboutUser: AboutUser
  }
}
</script>
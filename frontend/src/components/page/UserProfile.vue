<template>
  <div>
    <!-- Title -->
    <div class="text-h6 text-white q-mt-md q-ml-lg">
      <span class="text-blue">r/{{userInfo.username}}</span>'s profile.
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
        :canCreatePost="false"
        :criterion="'by'"
        :criterionKey="this.$route.params.username"
        :emptyText="'The user has not posted any articles.'"
      />

      <list-of-comments
        v-if="viewList === 'comment'"
        :username="this.$route.params.username"
      />

      <list-of-articles
        v-if="viewList === 'saved'"
        :canCreatePost="false"
        :criterion="'savedBy'"
        :criterionKey="this.$route.params.username"
        :emptyText="'The user has not saved any articles.'"
      />

      <!-- Right Panel -->
      <div class="col-3 q-pr-md q-pt-md gt-sm">
        <about-user :userInfo="userInfo" @bioSaved="bioSaved"/>
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
import AccountService from '../../services/account';

export default {
  data() {
    return {
      userInfo: {
        username: '',
        karma: 0,
        bio: '',
        joinTime: '1995-11-24T00:00:00Z'
      },
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
    async reloadPage(username) {
      this.userInfo.username = username;

      window.scrollTo(0, 0);
      document.title = 'u/' + username + ' - YARC';
      
      // Check if the user exists.
      try {
        this.userInfo = await AccountService.getUser(username);
      } catch (error) {
        this.$router.replace({name: 'page not found'});
      }
    },
    bioSaved(bio) {
      this.userInfo.bio = bio;
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
<template>
  <q-item class="column" clickable>
    <div class="row">
      <!-- Upvote, Downvote -->
      <q-item-section thumbnail>
        <voter :votes="article.points" :articleID="article.articleID" />
      </q-item-section>

      <!-- Thumbnail -->
      <q-item-section thumbnail>
        <q-avatar color="grey-9" rounded size="65px">
          <q-icon v-if="article.type !== 'image'" :name="thumbnailIcon" />
          <img v-if="article.type === 'image'" :src="article.body" />
        </q-avatar>
      </q-item-section>

      <!-- Article Title and info -->
      <q-item-section class="col">
        <q-item-label lines="1">
          <span class="text-h6">{{article.title}}</span>
        </q-item-label>

        <q-item-label lines="1" class="text-grey">
          <router-link @click.native.stop :to="'/r/' + article.subreddit">
            r/{{article.subreddit}}
          </router-link>
          <span>．Posted by </span>
          <router-link @click.native.stop :to="'/u/' + article.postedBy">
            u/{{article.postedBy}}
          </router-link>
          <span> on {{article.postedTime | formatDate}}</span>
        </q-item-label>

        <!-- Expand, Comments, Share, Save, Delete -->
        <q-item-label lines="1" class="text-grey text-weight-bold">
          <q-btn v-if="article.type !== 'link'" flat round @click.stop="expandClicked" :icon="expanded ? 'expand_less' : 'expand_more'" size="xs" />
          <q-btn v-if="article.type === 'link'" flat round @click.stop="openLinkInNewTab" size="xs" icon="open_in_new" />
          |
          <q-btn class="q-ml-sm" dense flat size="xs" icon="chat_bubble" :label="'' + article.comments + ' Comments'" />
          <q-btn @click.stop="shareClicked" class="q-ml-sm" dense flat size="xs" icon="share" label="Share" />
          <q-btn v-if="$store.state.auth" @click.stop="saveClicked" class="q-ml-sm" dense flat size="xs" :icon="articleSaved ? 'cancel_presentation' : 'save_alt'" :label="articleSaved ? 'Unsave' : 'Save'" :disabled="saveLoading" />
        </q-item-label>
      </q-item-section>
    </div>

    <!-- Expanded article content -->
    <div v-if="expanded" class="row q-pt-lg q-pl-lg q-pr-lg q-pb-sm">
      <img v-if="article.type === 'image'" :src="article.body" />
      <p v-if="article.type === 'text'" style="white-space: pre-wrap">{{article.body}}</p>
    </div>
  </q-item>
</template>

<script>
import Voter from './Voter';
import AccountService from '../../services/account';

export default {
  props: {
    article: Object
  },
  data() {
    return {
      expanded: false,
      saveLoading: false,
      articleSaved: false
    };
  },
  async mounted() {
    this.saveLoading = true;
    
    let auth = this.$store.state.auth;

    if (auth) {
      try {
        this.articleSaved = await AccountService.getSaveState(this.article.articleID, this.$store.state.auth.authHeader);
        this.saveLoading = false;
      } catch (error) {
        if (error.response && error.response.status === 401) {
          this.$store.commit('logout');
          localStorage.removeItem('user');
        }
      }
    }

    this.saveLoading = false;
  },
  computed: {
    thumbnailIcon() {
      if (this.article.type === 'text') {
        return 'chat';
      } else if (this.article.type === 'link') {
        return 'link';
      }

      // No icon for image.
      return '';
    }
  },
  methods: {
    expandClicked() {
      this.expanded = !this.expanded;
    },
    openLinkInNewTab() {
      window.open(this.article.body, '_blank');
    },
    shareClicked() {
      window.open('https://twitter.com/intent/tweet?text=' + this.article.title + '%0D' + location.protocol + '//' + location.host + '/r/' + this.article.subreddit + '/p/' + this.article.articleID, '_blank');
    },
    async saveClicked() {
      this.saveLoading = true;
      try {
        if (this.articleSaved) {
          await AccountService.unsaveArticle(this.article.articleID, this.$store.state.auth.authHeader);
        } else {
          await AccountService.saveArticle(this.article.articleID, this.$store.state.auth.authHeader);
        }
        this.articleSaved = !this.articleSaved;
      } catch (error) {
        if (error.response && error.response.status === 401) {
          this.$store.commit('logout');
          localStorage.removeItem('user');
        }
      }
      this.saveLoading = false;
    }
  },
  components: {
    voter: Voter
  }
}
</script>
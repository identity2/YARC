<template>
  <div class="row">
    <div class="col q-pa-lg">
      <!-- Create post text -->
      <div class="text-h6 text-white">Create a post</div>
      <q-separator color="grey" />

      <!-- Error Indicator -->
      <q-banner v-if="errMsg != ''" class="text-white bg-red q-mt-xs">
          Error: {{errMsg}}.
      </q-banner>

      <!-- Subreddit Selector -->
      <q-select
        class="q-mt-lg"
        dense
        dark
        outlined
        standout
        use-input
        hide-selected
        fill-input
        input-debounce="0"
        :options="subredditOptions"
        v-model="subreddit"
        @filter="filterFn"
        label="Choose Subreddit"
        style="width: 250px"
      >
        <template v-slot:no-option>
          <q-item>
            <q-item-section class="text-grey">No results</q-item-section>
          </q-item>
        </template>
      </q-select>

      <!-- Create Article Box -->
      <q-card dark class="q-mt-lg">
        <q-tabs v-model="postType">
          <q-tab name="text" label="Text" :disabled="loading" />
          <q-tab name="image" label="Image" :disabled="loading" />
          <q-tab name="link" label="Link" :disabled="loading" />
        </q-tabs>

        <q-separator color="grey-9" />

        <q-tab-panels class="bg-grey-10" v-model="postType" animated>
          <q-tab-panel name="text">
            <q-input class="q-mb-md" dark outlined dense standout v-model="title" label="title" counter :maxlength="titleMaxLen" :disabled="loading" />
            <q-input dark outlined standout v-model="textPostBody" type="textarea" placeholder="Text (optional)" counter :maxlength="textPostMaxLen" :disabled="loading" />
          </q-tab-panel>

          <q-tab-panel name="image">
            <q-input class="q-mb-md" dark outlined dense standout v-model="title" label="title" counter :maxlength="titleMaxLen" :disabled="loading" />
            <q-uploader
              v-if="imagePostUrl == ''"
              bordered
              flat
              dark
              color="grey-9"
              :max-files="1"
              :max-file-size="10*1024*1024"
              label="Upload Image (< 10 MB)"
              accept=".jpg, image/*"
              style="width: 100%"
              :factory="imageUploadFactory"
              @rejected="onImageRejected"
            />
            <img v-if="imagePostUrl != ''" :src="imagePostUrl" style="max-width: 600px; max-height: 600px; border:2px white solid" />
          </q-tab-panel>

          <q-tab-panel name="link">
            <q-input class="q-mb-md" dark outlined dense standout v-model="title" label="title" counter :maxlength="titleMaxLen" :disabled="loading" />
            <q-input class="q-mb-md" dark outlined dense standout v-model="linkPostUrl" label="Url" counter :maxlength="linkPostUrlMaxLen" :disabled="loading" />
          </q-tab-panel>
        </q-tab-panels>

        <q-card-actions class="bg-grey-10 q-pr-lg q-pb-lg" align="right">
          <q-btn :disabled="loading" @click="cancelClicked" outline class="q-mr-sm" style="color: white; width: 100px" label="Cancel" />
          <q-btn :loading="loading" @click="postClicked" :disabled="!validPost" style="background: white; color: black; width: 100px" label="Post" />
        </q-card-actions>

      </q-card>
    </div>

    <!-- Discard Article Dialog -->
    <q-dialog v-model="showDiscardDialog">
      <q-card class="bg-grey-10">
        <q-card-section class="row items-center">
          <span class="q-ml-sm text-white text-h6">Sure you want discard the post and go back to the home page?</span>
        </q-card-section>
        <q-card-actions align="right">
          <q-btn @click="discardClicked" outline style="color: red" label="Discard Post" />
          <q-btn style="background: white; color: black" label="Continue Editing" v-close-popup />
        </q-card-actions>
      </q-card>
    </q-dialog>

    <!-- Right Panel -->
    <div class="col-3 q-pr-md q-pt-md gt-sm">
      <tips
        title="Posting Rules"
        :tips="[
          '1. You do not talk about YARC.',
          '2. Do not post something important here, this site is not resposible of keeping your contents alive.',
          '3. Do not post offensive or NSFW contents, or I will simply reset the database.',
          '4. The fourth rule doesn\'t exist.'
        ]"
      />
      <advertisement />
    </div>
  </div>
</template>

<script>
import Tips from '../rightPanel/Tips';
import Advertisement from '../rightPanel/Advertisement';
import Limits from '../../limits';
import SubredditService from '../../services/subreddit';
import ArticleService from '../../services/article';
import ImgurService from '../../services/imgur';

export default {
  async mounted() {
    document.title = 'Submit Post - YARC';

    // Redirect to the log in page if the user isn't logged in.
    if (!this.$store.state.auth) {
      this.$router.push('/account?login');
    }

    // Load the subreddit list.
    try {
      this.allSubreddits = await SubredditService.getList();
      this.subredditOptions = this.allSubreddits;
    } catch {
      this.errMsg = 'Failed to load the subreddit list. Refresh the page and try again.'
    }
  },
  computed: {
    validPost() {
      if (this.subreddit === '' || this.title.trim() === '') {
        return false;
      }

      if (this.postType === 'image' && this.imagePostUrl === '') {
        return false;
      }

      if (this.postType === 'link' && this.linkPostUrl === '') {
        return false;
      }

      return true;
    }
  },
  data() {
    return {
      postType: 'text',
      subreddit: '',
      title: '',
      titleMaxLen: Limits.titleMaxLen,
      textPostMaxLen: Limits.textPostMaxLen,
      linkPostUrlMaxLen: Limits.linkPostUrlMaxLen,
      textPostBody: '',
      linkPostUrl: '',
      imagePostUrl: '',
      showDiscardDialog: false,
      allSubreddits: [],
      subredditOptions: [],
      loading: false,
      errMsg: ''
    };
  },
  watch: {
    postType(newVal) {
      this.$router.replace({
        query: {
          subreddit: this.subreddit,
          postType: newVal
        }
      }).catch(() => {});
    },
    subreddit(newVal) {
      this.$router.replace({
        query: {
          subreddit: newVal,
          postType: this.postType
        }
      }).catch(() => {});
    }
  },
  created() {
    this.postType = this.$route.query.postType;
    if (!this.postType) {
      this.postType = 'text';
    }
    this.subreddit = this.$route.query.subreddit;
  },
  methods: {
    cancelClicked() {
      this.showDiscardDialog = true;
    },
    postClicked() {
      this.loading = true;
      
      let body = this.textPostBody;
      if (this.postType === 'link') {
        body = this.linkPostUrl;
      } else if (this.postType === 'image') {
        body = this.imagePostUrl;
      }

      // Send request.
      ArticleService.create(this.subreddit, this.postType, body, this.title, this.$store.state.auth.authHeader).then(articleID => {
        // Successful, redirect to the newly created article.
        this.$router.push('/r/'+this.subreddit+'/p/'+articleID);
      }).catch(error => {
        if (error.response.status == 401) {
          this.errMsg = 'your login credentials have expired, please log in again';
          this.$store.commit('logout');
          localStorage.removeItem('user');
        } else if (this.postType === 'link') {
          this.errMsg = 'the link URL is invalid (make sure to include http(s)://)'
        } else {
          this.errMsg = 'failed to create the post, try again later';
        }
        this.loading = false;
      });
    },
    discardClicked() {
      this.$router.push('/');
    },
    filterFn(val, update) {
      update(() => {
        const needle = val.toLowerCase();
        this.subredditOptions = this.allSubreddits.filter(v => v.toLowerCase().indexOf(needle) > -1);
      });
    },
    onImageRejected() {
      this.errMsg = 'only one image is allowed, and it should be below 10 MB';
    },
    async imageUploadFactory(files) {
      try {
        this.imagePostUrl = await ImgurService.upload(this.title, files[0]);
      } catch {
        this.errMsg = 'failed to upload image, try again later';
      }
    },
  },
  components: {
    tips: Tips,
    advertisement: Advertisement
  }
}
</script>
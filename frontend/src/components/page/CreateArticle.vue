<template>
  <div class="row">
    <div class="col q-pa-lg">
      <!-- Create post text -->
      <div class="text-h6 text-white">Create a post</div>
      <q-separator color="grey" />

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
          <q-tab name="text" label="Text" />
          <q-tab name="image" label="Image" />
          <q-tab name="link" label="Link" />
        </q-tabs>

        <q-separator color="grey-9" />

        <q-tab-panels class="bg-grey-10" v-model="postType" animated>
          <q-tab-panel name="text">
            <q-input class="q-mb-md" dark outlined dense standout v-model="title" label="title" counter :maxlength="titleMaxLen" />
            <q-input dark outlined standout v-model="textPostBody" type="textarea" placeholder="Text (optional)" counter :maxlength="textPostMaxLen" />
          </q-tab-panel>

          <q-tab-panel name="image">
            <q-input class="q-mb-md" dark outlined dense standout v-model="title" label="title" counter :maxlength="titleMaxLen" />
            <q-uploader bordered flat dark color="grey-9" :url="imageUploadUrl" label="Upload Image" style="width: 100%" />
          </q-tab-panel>

          <q-tab-panel name="link">
            <q-input class="q-mb-md" dark outlined dense standout v-model="title" label="title" counter :maxlength="titleMaxLen" />
            <q-input class="q-mb-md" dark outlined dense standout v-model="linkPostUrl" label="Url" counter :maxlength="linkPostUrlMaxLen" />
          </q-tab-panel>
        </q-tab-panels>

        <q-card-actions class="bg-grey-10 q-pr-lg q-pb-lg" align="right">
          <q-btn @click="cancelClicked" outline class="q-mr-sm" style="color: white; width: 100px" label="Cancel" />
          <q-btn @click="postClicked" :disabled="!validPost" style="background: white; color: black; width: 100px" label="Post" />
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
          '2. Do not post something important here, this site will not be persistent.',
          '3. Do not post offensive or NSFW contents, or I will simply reset the database.',
          '4. The forth rule doesn\'t exist.'
        ]"
      />
      <advertisement />
    </div>
  </div>
</template>

<script>
const allSubreddits = ['radiohead', 'news', 'vuejs']; // Mock data.

import Tips from '../rightPanel/Tips';
import Advertisement from '../rightPanel/Advertisement';
import Limits from '../../limits';

export default {
  mounted() {
    document.title = 'Submit Post - YARC';
  },
  computed: {
    validPost() {
      if (this.subreddit === '' || this.title.trim() === '') {
        return false;
      }

      if (this.postType === 'image') {
        // TODO: Check if image is uploaded successfully.
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
      imagePostFile: '',
      imageUploadUrl: '',
      showDiscardDialog: false,
      subredditOptions: allSubreddits
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
      // TODO: Post the article.
    },
    discardClicked() {
      this.$router.push('/');
    },
    filterFn(val, update) {
      update(() => {
        const needle = val.toLowerCase();
        this.subredditOptions = allSubreddits.filter(v => v.toLowerCase().indexOf(needle) > -1);
      });
    }
  },
  components: {
    tips: Tips,
    advertisement: Advertisement
  }
}
</script>
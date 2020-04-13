<template>
<div>
  <!-- Sort, Create Post -->
  <div class="q-pt-md q-pl-md q-pr-md text-white">
    <q-list dark bordered class="bg-grey-10 rounded-borders">
      <q-item class="q-mt-sm q-mb-sm">          
          <q-btn-dropdown color="blue" label="Sort">
            <q-list>
              <q-item clickable v-close-popup>Hot</q-item>
              <q-item clickable v-close-popup>New</q-item>
              <q-item clickable v-close-popup>Old</q-item>
            </q-list>
          </q-btn-dropdown>

          <q-input class="col q-ml-lg q-mr-sm" dark clearable outlined dense standout label="Create Post" />
          <q-btn flat round color="grey" icon="insert_photo"></q-btn>
          <q-btn flat round color="grey" icon="insert_link"></q-btn>
      </q-item>
    </q-list>
  </div>

  <!-- Article Lists -->
  <div class="q-pa-md text-white">
    <q-list dark bordered separator class="bg-grey-10 rounded-borders">
      <article-entry></article-entry>
      <article-entry></article-entry>
      <article-entry></article-entry>
    
    </q-list>
  </div>
</div>
</template>

<script>
import ArticleEntry from '../article/ArticleEntry';

export default {
  components: {
    articleEntry: ArticleEntry
  }
}
</script>
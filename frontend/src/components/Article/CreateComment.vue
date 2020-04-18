<template>
  <div class="q-mt-md">
    <q-separator dark />
    <q-item-section>
      <div class="text-grey text-subtitle2 q-mt-md q-mb-sm q-ml-xs">
        <span>Add a comment as </span>
        <a class="text-blue" :href="'/user/' + currUser">u/{{currUser}}</a>
      </div>
      <q-input dark outlined standout v-model="commentText" type="textarea" placeholder="What are your thoughts?" counter :maxlength="commentMaxLen" />
    </q-item-section>

    <q-btn class="q-mr-md q-mb-lg" @click="commentClicked" style="background: white; color: black" label="Comment" />

    <q-separator dark />
  </div>
</template>

<script>
import Limits from '../../limits';

export default {
  data() {
    return {
      commentText: '',
      commentMaxLen: Limits.commentMaxLen,
      currUser: 'wes_ftw'
    };
  },
  methods: {
    commentClicked() {
      
    }
  }
}
</script>
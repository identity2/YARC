<!-- This is for the list of comments on a user's profile, not for the article. -->
<template>
  <div class="col">
    <div class="q-pa-md text-white">
      <q-list dark bordered separator class="bg-grey-10 rounded-borders">
        <q-item v-for="comment in comments" :key="comment.id" class="column">
          <q-item-label lines="1" class="q-mt-sm">
            Commented on post: 
            <router-link class="text-blue" :to="'/r/' + comment.subreddit + '/p/' + comment.postID">
              {{comment.postTitle}}
            </router-link>
          </q-item-label>
          <comment-content
            :commentedBy="comment.commentedBy"
            :points="comment.points"
            :commentedDate="comment.commentedDate"
            :commentText="comment.commentText"
          />
        </q-item>
      </q-list>
    </div>
  </div>
</template>

<script>
import CommentContent from '../article/CommentContent';

export default {
  props: {
    comments: Array
  },
  components: {
    commentContent: CommentContent
  }
}
</script>
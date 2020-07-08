<!-- This is for the list of comments on a user's profile, not for the article. -->
<template>
  <div class="col">
    <div class="q-pa-md text-white">
      <q-list dark bordered separator class="bg-grey-10 rounded-borders">
        <q-infinite-scroll @load="loadMoreComments" :offset="250" ref="infiniteScroll">
          <q-item v-for="(comment, index) in comments" :key="comment.commentID" class="column">
            <q-item-label lines="1" class="q-mt-sm">
              Commented on post: 
              <router-link class="text-blue" :to="'/r/' + comment.subreddit + '/p/' + comment.articleID">
                {{comment.articleTitle}}
              </router-link>
            </q-item-label>
            <comment-content
              :comment="comment"
              @commentDeleted="commentDeleted(comment.commentID)"
              @commentEdited="commentEdited"
            />
            <q-separator v-if="index !== comments.length-1" dark inset />
          </q-item>
        </q-infinite-scroll>
      </q-list>

      <!-- Error Indicator -->
      <q-banner v-if="errOccurred" class="text-white bg-red">
          Error fetching the comment list, please try again.
      </q-banner>
    </div>
  </div>
</template>

<script>
import CommentContent from '../article/CommentContent';
import ArticleService from '../../services/article';
import CommentService from '../../services/comment';

export default {
  props: {
    username: String
  },
  data() {
    return {
      comments: [],
      commentsPerRequest: 8,
      errOccurred: false,
      commentsLoading: false
    };
  },
  mounted() {
    this.commentsLoading = true;
    this.fetchCommentLists("").then(fetchedComments => {
      this.comments = fetchedComments;
      this.$refs.infiniteScroll.resume();
      this.errOccurred = false;
      this.commentsLoading = false;
    }).catch(() => {
      this.errOccurred = true;
    });
  },
  methods: {
    async fetchCommentLists(after) {
      let fetchedComments = await CommentService.getList(after, this.commentsPerRequest, "", this.username);

      for (let com of fetchedComments) {
        let article = await ArticleService.get(com.articleID);  // Get the title of the comments.
        com.articleTitle = article.title;
      }

      if (fetchedComments.length < this.commentsPerRequest) {
        this.$refs.infiniteScroll.stop();
      }
      return fetchedComments;
    },
    loadMoreComments(_, done) {
      if (this.commentsLoading) {
        done();
        return;
      }

      let after = this.comments.length === 0 ? "" : this.comments[this.comments.length-1].commentID;
      this.fetchCommentLists(after).then(fetchedComments => {
        for (const com of fetchedComments) {
          this.comments.push(com);
        }
        done();
      }).catch(() => {
        this.errOccurred = true;
        done(true);
      });
    },
    commentEdited(commentID, body) {
      for (let com of this.comments) {
        if (com.commentID === commentID) {
          com.body = body;
          break;
        }
      }
    },
    commentDeleted(commentID) {
      this.comments = this.comments.filter(comment => comment.commentID != commentID);
    }
  },
  components: {
    commentContent: CommentContent
  }
}
</script>
<template>
  <div class="row">
    <div class="col">
      <!-- Main Content -->
      <div class="q-pa-md">
        <!-- Error Indicator -->
        <q-banner v-if="errOccurred" class="text-white bg-red q-mb-md">
            Error: failed to load the article.
        </q-banner>

        <q-list dark bordered class="bg-grey-10 rounded-borders">
          <q-item class="column">
            <div class="row">
              <!-- Upvote, Downvote -->
              <q-item-section class="q-mt-lg" thumbnail top>
                <voter v-if="article" :votes="article.points" :articleID="article.articleID" />
              </q-item-section>

              <q-item-section class="col">
                <!-- Placeholder -->
                <div v-if="!article" class="col q-ma-md">
                  <q-item-label v-for="i in 3" :key="i">
                    <q-skeleton type="text" />
                  </q-item-label>
                </div>

                <div v-if="article">
                  <!-- Posted by, Date -->
                  <div class="text-caption text-grey q-ml-md q-mt-md">
                    <span>Posted by </span>
                    <router-link :to="'/u/' + article.postedBy">
                      u/{{article.postedBy}}
                    </router-link>
                    <span> on {{article.postedTime | formatDate}}</span>
                  </div>

                  <!-- Article -->
                  <article-content
                    :article="article"
                    :editMode="editMode"
                    @saveEdit="editSaved"
                    @cancelEdit="editMode = false"
                  />

                  <!-- Comments, Share, Save, Edit, Delete -->
                  <q-item-label v-if="!editMode" lines="1" class="text-grey q-ml-sm q-mt-sm q-mb-sm">
                    <q-btn class="q-ml-sm" dense flat size="xs" icon="chat_bubble" :label="'' + comments.length + ' Comments'" />
                    <q-btn @click="shareClicked" class="q-ml-sm" dense flat size="xs" icon="share" label="Share" />
                    <q-btn v-if="loggedInUser" @click="saveClicked" class="q-ml-sm" dense flat size="xs" :icon="articleSaved ? 'cancel_presentation' : 'save_alt'" :label="articleSaved ? 'Unsave' : 'Save'" :disable="saveBtnLoading" />
                    <q-btn v-if="canEdit" @click="editMode = true" class="q-ml-sm" dense flat size="xs" icon="create" label="Edit" />
                    <q-btn v-if="canDelete" @click="deleteClicked" class="q-ml-sm" dense flat size="xs" icon="delete" label="Delete" />
                  </q-item-label>

                  <div v-if="!editMode && bottomErrMsg != ''" class="q-ml-md text-red">Error: {{bottomErrMsg}}.</div>

                </div>
                
                <!-- New Comment -->
                <create-comment v-if="article" :articleID="article.articleID" :loggedInUser="loggedInUser" @commentCreated="commentCreated" />

                <!-- Comment Contents -->
                <q-infinite-scroll @load="loadMoreComments" :offset="250" ref="infiniteScroll">
                  <comment-content
                    v-for="comment in comments"
                    :key="comment.commentID"
                    :comment="comment"
                    @commentEdited="commentEdited"
                    @commentDeleted="commentDeleted(comment.commentID)"
                  />

                  <template v-slot:loading>
                    <div class="row justify-center q-my-md">
                      <q-spinner-dots color="white" size="40px" />
                    </div>
                  </template> 
                </q-infinite-scroll>

                <!-- No Comment -->
                <div v-if="!commentErrOccurred && comments.length === 0" class="q-mt-lg q-mb-md">
                  <span class="text-subtitle text-white">No comments yet...</span>
                </div>

                <!-- Comment Error Indicator -->
                <q-banner v-if="commentErrOccurred" class="text-white bg-red q-mt-md q-mb-md">
                    Error: failed to load the comments.
                </q-banner>
                
              </q-item-section>
            </div>
          </q-item>
        </q-list>
      </div>

      <!-- Delete confirm dialog -->
      <q-dialog v-model="showDeleteConfirm" :persistent="deleteLoading" >
        <q-card class="bg-white">
          <div class="q-ma-md text-subtitle">Sure you want to delete the article?</div>

          <!-- Error -->
          <div v-if="deleteErrMsg != ''" class="q-ml-md q-mr-md text-red text-subtitle">Error: {{deleteErrMsg}}.</div>
          
          <q-card-actions align="right">
            <q-btn flat label="Cancel" color="blue" v-close-popup :disabled="deleteLoading" />
            <q-btn flat label="Delete" color="red" @click="deleteConfirmClicked" :loading="deleteLoading" />
          </q-card-actions>
        </q-card>
      </q-dialog>

      <!-- Delete successful dialog -->
      <q-dialog v-model="showDeleteSuccessful" persistent>
        <q-card class="bg-orange text-white q-pt-md" style="width: 300px">
          <q-card-section class="text-white q-pt-none">
            The article has been deleted successfully!
          </q-card-section>
          <q-card-actions align="right" class="bg-white text-teal">
            <q-btn flat label="Alright" @click="deleteSuccessClicked"/>
          </q-card-actions>
        </q-card>
      </q-dialog>

    </div>

    <!-- Right Panel -->
    <div class="col-3 q-pr-md q-pt-md gt-sm">
      <about-subreddit :subreddit="article ? article.subreddit : ''" />
      <advertisement />
    </div>
  </div>
</template>

<script>
import AboutSubreddit from '../rightPanel/AboutSubreddit';
import Voter from '../article/Voter';
import ArticleContent from '../article/ArticleContent';
import CommentContent from '../article/CommentContent';
import CreateComment from '../article/CreateComment';
import Advertisement from '../rightPanel/Advertisement';
import ArticleService from '../../services/article';
import CommentService from '../../services/comment';
import AccountService from '../../services/account';

export default {
  data() {
    return {
      article: null,
      comments: [],
      editMode: false,
      errOccurred: false,
      bottomErrMsg: '',
      commentErrOccurred: false,
      deleteErrMsg: '',
      commentsPerRequest: 5,
      loggedInUser: null,
      articleSaved: false,
      saveBtnLoading: false,
      deleteLoading: false,
      showDeleteSuccessful: false,
      showDeleteConfirm: false,
      initialCommentLoaded: false
    };
  },
  computed: {
    canEdit() {
      return this.article.type === 'text' && this.loggedInUser && this.loggedInUser.username === this.article.postedBy;
    },
    canDelete() {
      return this.loggedInUser && this.loggedInUser.username === this.article.postedBy;
    }
  },
  watch: {
    $route(newVal) {
      this.reload(newVal.params.pid);
    }
  },
  mounted() {
    this.reload(this.$route.params.pid);
  },
  methods: {
    async reload(articleID) {
      // Check the currently logged in user.
      this.loggedInUser = this.$store.state.auth;

      this.initialCommentLoaded = false;

      // Load the article content.
      this.article = null;
      try {
        this.article = await ArticleService.get(articleID);
        this.errOccurred = false;
        document.title = this.article.title + ' - YARC';
      } catch (error) {
        if (error.response && error.response.status === 404) {
          this.$router.replace({name: 'page not found'});
        }

        this.errOccurred = true;
        return;
      }

      // Check if the article is saved by the logged in user.
      this.articleSaved = false;
      this.saveBtnLoading = true;
      this.bottomErrMsg = '';
      if (this.loggedInUser) {
        try {
          this.articleSaved = await AccountService.getSaveState(articleID, this.loggedInUser.authHeader);
          this.saveBtnLoading = false;
        } catch (error) {
          if (error.response && error.response.status === 401) {
            this.$store.commit('logout');
            localStorage.removeItem('user');
          }
          this.bottomErrMsg = 'Failed to fetch the state of saving.';
        }
      }

      // Load the comments.
      this.comments = [];
      try {
        this.comments = await this.fetchCommentLists("");
        this.$refs.infiniteScroll.resume();
        this.commentErrOccurred = false;
        this.initialCommentLoaded = true;
      } catch (error) {
        this.commentErrOccurred = true;
      }
    },
    async fetchCommentLists(after) {
      let fetchedComments = await CommentService.getList(after, this.commentsPerRequest, this.article.articleID, "");

      // Check if it is the end of the comment list. If it is, stop the scroll.
      if (fetchedComments.length < this.commentsPerRequest) {
        this.$refs.infiniteScroll.stop();
      }

      return fetchedComments;
    },
    loadMoreComments(_, done) {
      // Article not loaded yet.
      if (!this.article || !this.initialCommentLoaded) {
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
        this.commentErrOccurred = true;
        done(true);
      });
    },
    shareClicked() {
      window.open('https://twitter.com/intent/tweet?text=' + this.article.title + '%0D' + location.href, '_blank');
    },
    async saveClicked() {
      this.saveBtnLoading = true;
      try {
        if (this.articleSaved) {
          await AccountService.unsaveArticle(this.article.articleID, this.loggedInUser.authHeader);
        } else {
          await AccountService.saveArticle(this.article.articleID, this.loggedInUser.authHeader);
        }
        this.articleSaved = !this.articleSaved;
      } catch (error) {
        if (error.response && error.response.status === 401) {
          this.$store.commit('logout');
          localStorage.removeItem('user');
          this.bottomErrMsg = 'login credentials expired, please re-login';
        } else {
          this.bottomErrMsg = 'failed to save the article, please try again';
        }
      }
      this.saveBtnLoading = false;
    },
    deleteClicked() {
      this.showDeleteConfirm = true;
      this.deleteErrOccurred = false;
    },
    deleteConfirmClicked() {
      // Send the delete request.
      this.deleteLoading = true;
      
      ArticleService.remove(this.article.articleID, this.loggedInUser.authHeader).then(() => {
        this.deleteLoading = false;
        this.showDeleteConfirm = false;
        this.showDeleteSuccessful = true;
      }).catch(error => {
        if (error.response && error.response.status === 401) {
          this.$store.commit('logout');
          localStorage.removeItem('user');
          this.deleteErrMsg = 'login credentials expired, please re-login';
        }
        this.deleteErrMsg = 'failed to delete the article, please try again';
        this.deleteLoading = false;
      });

    },
    deleteSuccessClicked() {
      this.$router.push('/r/'+this.article.subreddit);
    },
    editSaved(newBody) {
      this.article.body = newBody;
      this.editMode = false;
    },
    commentCreated(comment) {
      this.comments.unshift(comment);
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
    aboutSubreddit: AboutSubreddit,
    voter: Voter,
    articleContent: ArticleContent,
    commentContent: CommentContent,
    createComment: CreateComment,
    advertisement: Advertisement
  }
}
</script>
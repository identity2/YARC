<template>
  <div class="row">
    <div class="col">
      <!-- Main Content -->
      <div class="q-pa-md">
        <q-list dark bordered class="bg-grey-10 rounded-borders">
          <q-item class="column">
            <div class="row">
              <!-- Upvote, Downvote -->
              <q-item-section class="q-mt-lg" thumbnail top>
                <voter :votes="points" />
              </q-item-section>

              <!-- Posted by, Date -->
              <q-item-section class="col">
                <div class="text-caption text-grey q-ml-md q-mt-md">
                  <span>Posted by </span>
                  <router-link :to="'/u/' + postedBy">
                    u/{{postedBy}}
                  </router-link>
                  <span> on {{postedDate}}</span>
                </div>
                
                <!-- Article -->
                <article-content
                  :title="title"
                  :postType="postType"
                  :textBody="textBody"
                  :imageUrl="imageUrl"
                  :linkUrl="linkUrl"
                  :editMode="editMode"
                  @cancelEdit="editMode = false"
                />

                <!-- Comments, Share, Save, Edit, Delete -->
                <q-item-label v-if="!editMode" lines="1" class="text-grey q-ml-sm q-mt-sm q-mb-sm">
                  <q-btn class="q-ml-sm" dense flat size="xs" icon="chat_bubble" :label="'' + comments.length + ' Comments'" />
                  <q-btn @click="shareClicked" class="q-ml-sm" dense flat size="xs" icon="share" label="Share" />
                  <q-btn @click="saveClicked" class="q-ml-sm" dense flat size="xs" icon="save_alt" label="Save" />
                  <q-btn v-if="postType === 'text'" @click="editMode = true" class="q-ml-sm" dense flat size="xs" icon="create" label="Edit" />
                  <q-btn @click="deleteClicked" class="q-ml-sm" dense flat size="xs" icon="delete" label="Delete" />
                </q-item-label>

                <!-- New Comment -->
                <create-comment />

                <!-- Comment Contents -->
                <comment-content
                  v-for="comment in comments"
                  :key="comment.id"
                  :commentedBy="comment.commentedBy"
                  :points="comment.points"
                  :commentedDate="comment.commentedDate"
                  :commentText="comment.commentText"
                />
                
              </q-item-section>
            </div>
          </q-item>
        </q-list>
      </div>
    </div>

    <!-- Right Panel -->
    <div class="col-3 q-pr-md q-pt-md gt-sm">
      <about-subreddit :subreddit="subreddit" :members="subredditMembers" />
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
import mockComments from '../../mock_data/mock_comments';

export default {
  data() {
    return {
      title: 'This is the title of the article.',
      postType: 'text',
      textBody: 'This is the body of the text post.',
      imageUrl: '',
      linkUrl: '',
      points: 39,
      postedBy: 'boneless_kandra',
      comments: mockComments,
      subreddit: 'mistborn',
      subredditMembers: 32,
      postedDate: Date(),
      editMode: false
    };
  },
  mounted() {
    this.subreddit = this.$route.params.subreddit;
    // TODO: Load the articles, change title.
  },
  methods: {
    shareClicked() {

    },
    saveClicked() {

    },
    deleteClicked() {

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
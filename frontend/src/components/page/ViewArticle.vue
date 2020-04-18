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
                <voter :votes="votes"></voter>
              </q-item-section>

              <!-- Posted by, Date -->
              <q-item-section class="col">
                <div class="text-caption text-grey q-ml-md q-mt-md">
                  <span>Posted by </span>
                  <a :href="'/user/' + postedBy">u/{{postedBy}}</a>
                  <span> on {{postedDate}}</span>
                </div>
                
                <!-- Article -->
                <article-content
                  :title="title"
                  :textBody="textBody"
                  :editMode="editMode"
                  @cancelEdit="editMode = false"
                ></article-content>

                <!-- Comments, Share, Save, Edit, Delete -->
                <q-item-label v-if="!editMode" lines="1" class="text-grey q-ml-sm q-mt-sm q-mb-sm">
                  <q-btn class="q-ml-sm" dense flat size="xs" icon="chat_bubble" :label="'' + comments.length + ' Comments'" />
                  <q-btn @click="shareClicked" class="q-ml-sm" dense flat size="xs" icon="share" label="Share" />
                  <q-btn @click="saveClicked" class="q-ml-sm" dense flat size="xs" icon="save_alt" label="(Save)" />
                  <q-btn @click="editMode = true" class="q-ml-sm" dense flat size="xs" icon="create" label="(Edit)" />
                  <q-btn @click="deleteClicked" class="q-ml-sm" dense flat size="xs" icon="delete" label="(Delete)" />
                </q-item-label>

                <!-- New Comment -->
                <create-comment></create-comment>

                <!-- Comment Contents -->
                <comment-content
                  v-for="comment in comments"
                  :key="comment.id"
                  :commentedBy="comment.commentedBy"
                  :points="comment.points"
                  :commentedDate="comment.commentedDate"
                  :commentText="comment.commentText"
                ></comment-content>
                
              </q-item-section>
            </div>
          </q-item>
        </q-list>
      </div>
    </div>

    <!-- Right Panel -->
    <div class="col-3 q-pr-md q-pt-md gt-sm">
      <about-subreddit></about-subreddit>
    </div>
  </div>
</template>

<script>
import AboutSubreddit from '../rightPanel/AboutSubreddit';
import Voter from '../article/Voter';
import ArticleContent from '../article/ArticleContent';
import CommentContent from '../article/CommentContent';
import CreateComment from '../article/CreateComment';
import mockComments from '../../mock_data/mock_comments';

export default {
  data() {
    return {
      votes: 42,
      postedBy: 'wigfrid',
      postedDate: Date(1587340801),
      title: 'This is the title of the article.',
      textBody: "Please could you stop the noise I'm trying to get some rest! For all the unborn chikens voices in my head! What was that? When I am king I will be first against the wall. With your opinion which is of no consequence at all. What was that, what was that?",
      comments: mockComments,
      editMode: false
    };
  },
  components: {
    aboutSubreddit: AboutSubreddit,
    voter: Voter,
    articleContent: ArticleContent,
    commentContent: CommentContent,
    createComment: CreateComment,
  },
  methods: {
    shareClicked() {

    },
    saveClicked() {

    },
    deleteClicked() {

    }
  }
}
</script>
<template>
  <q-card dark flat bordered class="bg-grey-10 q-mb-md">
    <q-card-section>
      <div class="text-h6 q-mb-md">{{subreddit}}</div>
      <div>{{description}}</div>
    </q-card-section>

    <q-separator dark inset />

    <q-card-section>
      <div class="text-center">{{members}} Members</div>
    </q-card-section>


    <q-card-section class="column q-ml-md q-mr-md q-gutter-sm">
      <q-btn @click="createPostClicked" color="white" text-color="black" class="row" label="Create Post" />
      <q-btn @click="joinClicked" color="white" text-color="black" class="row" label="Join" />
    </q-card-section>
  </q-card>
</template>

<script>
export default {
  data() {
    return {
      subreddit: 'r/subreddit',
      members: 0,
      description: 'The description of the subreddit goes here!'
    };
  },
  methods: {
    createPostClicked() {

    },
    joinClicked() {
      
    }
  }
}
</script>

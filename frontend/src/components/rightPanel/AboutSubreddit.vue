<template>
  <q-card dark flat bordered class="bg-grey-10 q-mb-md">
    <q-card-section>
      <div class="text-h6 q-mb-md">r/{{subreddit}}</div>
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
  props: {
    subreddit: String,
    members: Number,
    description: String
  },
  methods: {
    createPostClicked() {

    },
    joinClicked() {
      
    }
  }
}
</script>

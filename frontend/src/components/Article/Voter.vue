<template>
  <div class="column">
    <div class="col">
      <q-btn @click.stop="upvoteClicked" round flat size="xs" icon="arrow_upward" />
    </div>
    <div class="col text-center">
      {{votes}}
    </div>
    <div class="col">
      <q-btn @click.stop="downvoteClicked" round flat size="xs" icon="arrow_downward" />
    </div>
  </div>
</template>

<script>
export default {
  props: {
    votes: Number
  },
  data() {
    return {
    };
  },
  methods: {
    upvoteClicked() {
      // TODO
    },
    downvoteClicked() {
      // TODO
    }
  }
}
</script>
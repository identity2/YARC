<template>
  <q-item clickable>
    <q-item-section avatar>
      <voter></voter>
    </q-item-section>

    <q-item-section class="col-1 gt-sm">
      <q-item-label>(Thumbnail)</q-item-label>
    </q-item-section>

    <q-item-section>
      <q-item-label lines="1">
        <span class="text-weight-medium">Article Title Here</span>
      </q-item-label>

      <q-item-label lines="1">
        <span>r/subreddit</span>
        <span>．Posted by u/user time ago</span>
      </q-item-label>

      <q-item-label lines="1">
        (Expand) | 10 Comments | Share | (Save) | (Edit) | (Delete)
      </q-item-label>
    </q-item-section>
  </q-item>
</template>

<script>
import Voter from './Voter';
export default {

  components: {
    voter: Voter
  }
}
</script>

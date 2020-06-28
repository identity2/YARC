<template>
  <q-card dark flat bordered class="bg-grey-10 q-mb-md q-pl-md">
    <q-card-section>
      <div class="text-h6 q-mb-md">Trending Subreddits</div>
      <q-list dark>

        <q-item v-if="subreddits.length == 0">
          Wow, such empty! Browse some subreddits to fill this list!
        </q-item>

        <q-item v-for="sub in subreddits" :key="sub.name">
          <q-item-section thumbnail>
            <q-avatar size="md" color="blue" text-color="white">
              {{sub.name[0].toUpperCase()}}
            </q-avatar>
          </q-item-section>
          <q-item-section>
            <q-item-label>
              <router-link :to="'/r/' + sub.name">
                <span class="text-white">r/{{sub.name}}</span>
              </router-link>
            </q-item-label>
            <q-item-label caption lines="1">Members: {{sub.members}}</q-item-label>
          </q-item-section>
        </q-item>

      </q-list>
    </q-card-section>
  </q-card>
</template>

<script>
import SubredditService from '../../services/subreddit';

const LIMIT = 5;

export default {
  async mounted() {
    try {
      this.subreddits = await SubredditService.getTrending(LIMIT);
    } catch {
      this.subreddits = [];
    }
  },
  data() {
    return {
      subreddits: []
    };
  }
}
</script>
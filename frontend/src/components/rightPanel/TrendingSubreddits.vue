<template>
  <q-card dark flat bordered class="bg-grey-10 q-mb-md q-pl-md">
    <q-card-section>
      <div class="text-h6 q-mb-md">Trending Subreddits</div>
      <q-list dark>
        <q-item v-for="sub in subreddits" :key="sub.name">
          <q-item-section thumbnail>
            <q-avatar size="md" color="blue" text-color="white">
              {{sub.name[0].toUpperCase()}}
            </q-avatar>
          </q-item-section>
          <q-item-section>
            <q-item-label>r/{{sub.name}}</q-item-label>
            <q-item-label caption lines="1">Members: {{sub.members}}</q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-card-section>
  </q-card>
</template>

<script>
export default {
  data() {
    return {
      subreddits: [
        {name: 'radiohead', members: 1311},
        {name: 'covid19', members: 420},
        {name: 'bikinibottom_twitter', members: 190},
        {name: 'cuhk', members: 42},
        {name: 'dontstarvetogether', members: 13}
      ]
    };
  }
}
</script>
<template>
  <q-card dark flat bordered class="bg-grey-10 q-mb-md q-pl-md">
    <q-card-section>
      <div class="text-h6">New Subreddit</div>
    </q-card-section>
    <q-card-actions align="center">
      <q-btn @click="$router.push('/subreddits/create')" class="q-mb-md" flat style="background: white; color: black; width: 70%" label="Create" />
    </q-card-actions>
  </q-card>
</template>
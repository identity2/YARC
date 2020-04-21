<template>
  <q-card dark flat bordered class="bg-grey-10 q-mt-xl q-mb-md q-pa-md">
    <div class="text-h6 q-mb-md">{{title}}</div>
    <q-list dark separator>
      <q-item v-for="(item, index) in tips" :key="index">{{item}}</q-item>
    </q-list>
  </q-card>
</template>

<script>
export default {
  props: {
    title: String,
    tips: Array
  }
}
</script>
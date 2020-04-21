<template>
  <q-card dark flat bordered class="bg-grey-10 q-mb-md">
    <q-card-section>
      <div class="text-h6 q-mb-md">About u/{{username}}</div>
      <div>
        <span class="text-blue">u/{{username}}</span>
         joined YARC at 
        <span class="text-grey">{{joined}}</span>.
      </div>
      <br />
      <div>Is good user!</div>
    </q-card-section>

    <q-separator dark inset />

    <q-card-section>
      <div class="text-center text-weight-bold">Karma Points: {{karma}}</div>
    </q-card-section>
  </q-card>
</template>

<script>
export default {
  props: {
    username: String,
    karma: Number,
    joined: String
  }
}
</script>
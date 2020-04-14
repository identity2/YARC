<template>
  <div class="column q-ml-md">
    <div class="col">
      <q-btn round flat size="xs" icon="arrow_upward"></q-btn>
    </div>
    <div class="col text-center">
      {{votes}}
    </div>
    <div class="col">
      <q-btn round flat size="xs" icon="arrow_downward"></q-btn>
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
  }
}
</script>
<template>
  <div>
    <q-card-section>
      <div class="text-h5">{{title}}</div>
      <div v-if="!editMode" class="q-mt-lg">{{textBody}}</div>
      <div v-if="editMode" class="q-mt-lg">
        <div class="q-ml-xs q-mb-sm">Editing post:</div>
        <q-input dark outlined standout v-model="textBody" type="textarea" counter :maxlength="textPostMaxLen" />
        <q-btn class="q-mr-md" @click="saveClicked" style="background: white; color: black" label="Save" />
        <q-btn outline @click="cancelClicked" style="color: red" label="Cancel" />
      </div>
    </q-card-section>
  </div>
</template>

<script>
import Limits from '../../limits';

export default {
  props: {
    title: String,
    textBody: String,
    editMode: Boolean
  },
  data() {
    return {
      textPostMaxLen: Limits.textPostMaxLen
    };
  },
  methods: {
    saveClicked() {

    },
    cancelClicked() {
      this.$emit('cancelEdit');
    }
  }
}
</script>

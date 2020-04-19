<template>
  <div>
    <q-card-section>
      <div class="text-h5">{{title}}</div>
      
      <!-- Text Post -->
      <div v-if="!editMode && postType === 'text'" class="q-mt-lg">{{textBody}}</div>
      <div v-if="editMode" class="q-mt-lg">
        <div class="q-ml-xs q-mb-sm">Editing post:</div>
        <q-input dark outlined standout v-model="textBody" type="textarea" counter :maxlength="textPostMaxLen" />
        <q-btn class="q-mr-md" @click="saveClicked" style="background: white; color: black" label="Save" />
        <q-btn outline @click="cancelClicked" style="color: red" label="Cancel" />
      </div>

      <!-- Image Post -->
      <div v-if="postType === 'image'" class="q-mt-lg">
        <img :src="imageUrl" style="max-width: 100%; max-height: 100%" />
      </div>

      <!-- Link Post -->
      <div v-if="postType === 'link'" class="q-mt-lg">
        <a :href="linkUrl" class="text-blue" target="_blank">
          {{linkUrl}}
          <q-icon name="open_in_new" size="xs" />
        </a>
      </div>
    </q-card-section>
  </div>
</template>

<script>
import Limits from '../../limits';

export default {
  props: {
    title: String,
    postType: String,
    textBody: String,
    imageUrl: String,
    linkUrl: String,
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

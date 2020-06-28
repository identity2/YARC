<template>
  <div>
    <q-card-section>
      <div class="text-h5">{{title}}</div>
      
      <!-- Text Post -->
      <div v-if="!editMode && postType === 'text'" class="q-mt-lg">{{textPostBody}}</div>
      <div v-if="editMode" class="q-mt-lg">
        <div class="q-ml-xs q-mb-sm">Editing post:</div>
        <q-input dark outlined standout v-model="textPostBody" type="textarea" counter :maxlength="textPostMaxLen" />
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
      textPostMaxLen: Limits.textPostMaxLen,
      textPostBody: this.textBody
    };
  },
  methods: {
    saveClicked() {
      // Inform the parent for the change
      this.$emit('saveEdit', this.textPostBody, () => {
        // Failure method.
        this.textPostBody = this.textBody;  // Reset to the original content.
      });
    },
    cancelClicked() {
      this.textPostBody = this.textBody;  // Reset to the original content.
      this.$emit('cancelEdit');
    }
  }
}
</script>

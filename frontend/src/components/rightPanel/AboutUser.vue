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

      <!-- Short Bio -->
      <div v-if="!editMode">{{shortBio}}</div>
      <div v-if="editMode">
        <q-input dark outlined standout v-model="shortBio" type="textarea" counter :maxlength="shortBioMaxLen" />
      </div>
    
    </q-card-section>

    <q-card-actions align="center">
      <q-btn @click="editMode = !editMode" class="q-mb-md" flat style="background: white; color: black; width: 50%" :label="editMode ? 'Save' : 'Edit'" />
    </q-card-actions>

    <q-separator dark inset />

    <q-card-section>
      <div class="text-center text-weight-bold">Karma Points: {{karma}}</div>
    </q-card-section>
  </q-card>
</template>

<script>
import Limits from '../../limits';

export default {
  props: {
    username: String,
    karma: Number,
    joined: String
  },
  data() {
    return {
      editMode: false,
      shortBio: "I am a very good user, mark my words!",
      shortBioMaxLen: Limits.shortBioMaxLen
    };
  },
  methods: {
    btnClicked() {
      
    }
  }
}
</script>
<template>
  <v-app id="forbidden">
    <v-main>
      <v-container fill-height>
        <v-row>
          <v-col cols="4" offset="4">
            <v-card class="justify-center">
              <v-card-title class="justify-center">
                You are not permitted to carry out this action.
              </v-card-title>
              <v-card-actions class="justify-center">
                <v-btn color="primary" to="/">
                  Home
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";

@Component
export default  class Forbidden extends Vue{

}
</script>
<template>
  <v-app id="logout">
    <v-main v-if="message">
      <v-container fill-height>
        <v-row>
          <v-col cols="4" offset="4">
            <v-card class="justify-center">
              <v-card-title class="justify-center">
                Logged out Successfully!
              </v-card-title>
              <v-card-actions class="justify-center">
                <v-btn color="primary" to="/login">
                  Login Again
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
import Vue from "vue";
import {Component} from "vue-property-decorator";
import {USERTOKEN} from "@/components/types";

@Component
export default class Logout extends Vue{
  message: string | null = null;
  beforeMount() {
    localStorage.removeItem(USERTOKEN);
    this.message = "Logged Out Successfully!";
  }
}
</script>
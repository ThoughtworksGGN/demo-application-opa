<template>
  <div id="logout">
    {{ this.message }}
  </div>
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
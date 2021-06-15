<template>
  <v-app id="home">
    <v-main>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import Vue from "vue";
import {Component} from "vue-property-decorator";
import {USERTOKEN} from "@/components/types";
import {getAccountForUser, getUserDataFromToken} from "@/components/service";

@Component
export default class Home extends Vue{
  async beforeMount() {
    const token = localStorage.getItem(USERTOKEN)
    const user = getUserDataFromToken();

    if (user.role === "CUSTOMER") {
      const accountId = (await getAccountForUser()).account_id;
      await this.$router.push({path: `account/${accountId}`});
    } else if (user.role === "SUPPORT") {
      await this.$router.push({path: `tickets`});
    } else {
      await this.$router.push({path: `forbidden`});
    }
  }
}
</script>
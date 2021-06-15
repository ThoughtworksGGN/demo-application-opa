<template>
  <div id="login">
    <h1>Login</h1>
    <input type="text" name="username" v-model="username" placeholder="Username" />
    <input type="password" name="password" v-model="password" placeholder="Password" />
    <button type="button" v-on:click="login()">Login</button>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator';
import {getAccountForUser, getUserDataFromToken, login} from './service'
import {loginRequest} from "@/components/types";

@Component
export default class LoginPage extends Vue{
  username = '';
  password = '';

  async login(): Promise<void> {
    const loginRequest: loginRequest = {
      name: this.username,
      password: this.password
    }

    await login(loginRequest);

    const user = await getUserDataFromToken();

    if (user.role === "CUSTOMER") {
      const accountId = (await getAccountForUser()).account_id;
      console.log(`MUBARAK HO CUSTOMER HUA HAI ! Account ID: ${accountId}`);
    }
  }
}
</script>
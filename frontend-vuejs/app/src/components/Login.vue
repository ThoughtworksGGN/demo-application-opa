<template>
  <v-app id="login">
    <v-main>
      <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="primary">
                <v-toolbar-title>Welcome to Internet Movies</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <v-form>
                  <v-text-field
                      name="login"
                      label="Login"
                      type="text"
                      v-model="username"
                  ></v-text-field>
                  <v-text-field
                      id="password"
                      name="password"
                      label="Password"
                      type="password"
                      v-model="password"
                  ></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" v-on:click="login()">Login</v-btn>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
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

    const user = getUserDataFromToken();

    if (user.role === "CUSTOMER") {
      const accountId = (await getAccountForUser()).account_id;
      await this.$router.push({path: `account/${accountId}`});
    } else if (user.role === "SUPPORT") {
      await this.$router.push({path: `tickets`});
    }
  }
}
</script>
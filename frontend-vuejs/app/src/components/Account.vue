<template>
  <v-app id="account-details">
    <v-main v-if="account">
      <v-app-bar fixed color="primary" dark class="text-no-wrap">
        <v-app-bar-title>
          <span class="text-no-wrap-demo">Account Details</span>
        </v-app-bar-title>
        <v-menu offset-y>
          <template v-slot:activator="{ on,attrs }">
            <v-btn light absolute right v-bind="attrs" v-on="on">{{ user.name }}</v-btn>
          </template>
          <v-list>
            <v-list-item to="/logout">
              <v-list-item-title>Logout</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </v-app-bar>
      <v-container fill-height>
        <v-row>
          <v-col offset="1" cols="4">
            <v-card id="details">
              <v-toolbar dark color="primary">
                <v-toolbar-title>
                  Details
                </v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <v-simple-table>
                  <tbody>
                  <tr>
                    <td>Account ID</td>
                    <td>{{account.account_id}}</td>
                  </tr>
                  <tr>
                    <td>User</td>
                    <td>{{account.user.name}}</td>
                  </tr>
                  </tbody>
                </v-simple-table>
              </v-card-text>
            </v-card>
          </v-col>
          <v-col offset="1" cols="4">
            <v-card id="notifications">
              <v-toolbar dark color="orange">
                <v-toolbar-title>Notifications</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <v-simple-table>
                  <thead>
                  <tr>
                    <th>Notification Id</th>
                    <th>Info</th>
                  </tr>
                  </thead>
                  <tbody>
                  <tr v-for="notification in account.notifications" :key="notification.notificationid">
                    <td>{{notification.notificationid}}</td>
                    <td>{{notification.info}}</td>
                  </tr>
                  </tbody>
                </v-simple-table>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
        <v-row>
          <v-col offset="1" cols="4">
            <v-card id="payments">
            <v-toolbar dark color="green">
              <v-toolbar-title>Payments</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
              <v-simple-table>
                <thead>
                <tr>
                  <th>Payment Id</th>
                  <th>Amount</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="payment in account.payments" :key="payment.paymentid">
                  <td>{{payment.paymentid}}</td>
                  <td>{{payment.amount}}</td>
                </tr>
                </tbody>
              </v-simple-table>
            </v-card-text>
          </v-card>
          </v-col>
          <v-col offset="1" cols="4">
            <v-card id="tickets">
              <v-toolbar dark color="red">
                <v-toolbar-title>Tickets</v-toolbar-title>
              </v-toolbar>
              <v-card-text>
                <v-simple-table>
                  <thead>
                  <tr>
                    <th>Ticket Id</th>
                    <th>Info</th>
                  </tr>
                  </thead>
                  <tbody>
                  <tr v-for="ticket in account.tickets" :key="ticket.ticketid">
                    <td>{{ticket.ticketid}}</td>
                    <td>{{ticket.info}}</td>
                  </tr>
                  </tbody>
                </v-simple-table>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import {Vue} from "vue-property-decorator";
import Component from "vue-class-component";
import {account, user} from "@/components/types";
import {getAccount, getUserDataFromToken} from "@/components/service";

@Component
export default class AccountDetails extends Vue{
 account: account | null = null;
 user: user | null = null ;
  async beforeMount(){
    this.user = getUserDataFromToken();

    const accountId = this.$route.params.accountid;
    const accountForUserResponse = await getAccount(accountId);
    if (accountForUserResponse.responseCode === 403) {
      await this.$router.push({path:"forbidden"});
    } else {
      this.account = accountForUserResponse.account;
    }
  }
}
</script>

<style scoped>
.text-no-wrap-demo {

}
</style>
<template>
  <v-app id="tickets">
    <v-main v-if="tickets">
      <v-app-bar fixed color="red" dark>
        <v-app-bar-title>
          <span class="text-no-wrap-demo">Tickets</span>
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
          <v-card id="tickets">
            <v-toolbar dark color="red">
              <v-toolbar-title>Tickets</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
              <v-simple-table>
                <thead>
                <tr>
                  <th>Ticket Id</th>
                  <th>Account Id</th>
                  <th>Info</th>
                  <th>Assigned to me</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="ticket in tickets" :key="ticket.ticketid">
                  <td>{{ticket.ticketid}}</td>
                  <td>
                    <router-link :to="`/account/${ticket.accountid}`">{{ticket.accountid}}</router-link>
                  </td>
                  <td>{{ticket.info}}</td>
                  <td>{{ticket.assignee == user.userid}}</td>
                </tr>
                </tbody>
              </v-simple-table>
            </v-card-text>
          </v-card>
        </v-row>
      </v-container>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {getTickets, getUserDataFromToken} from "@/components/service";
import {ticket, user} from "@/components/types";

@Component
export default class Tickets extends Vue{
  tickets: ticket[] | null = null;
  user: user | null = null;
  async beforeMount() {
    this.user = getUserDataFromToken();
    const ticketsResponse = await getTickets();
    if (ticketsResponse.responseCode === 403) {
      await this.$router.push({path: "forbidden"});
    }
    this.tickets = ticketsResponse.tickets
  }
}
</script>
<template>
  <v-app id="tickets">
    <v-main v-if="tickets">
      <v-app-bar fixed color="red" dark>
        <v-app-bar-title>
          <span class="text-no-wrap-demo">Tickets</span>
        </v-app-bar-title>
        <v-btn light absolute right to="/logout">
          Logout
        </v-btn>
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
                </tr>
                </thead>
                <tbody>
                <tr v-for="ticket in tickets" :key="ticket.ticketid">
                  <td>{{ticket.ticketid}}</td>
                  <td>
                    <router-link :to="`/account/${ticket.accountid}`">{{ticket.accountid}}</router-link>
                  </td>
                  <td>{{ticket.info}}</td>
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
import {getTickets} from "@/components/service";
import {ticket} from "@/components/types";

@Component
export default class Tickets extends Vue{
  tickets: ticket[] | null = null;
  async beforeMount() {
    const ticketsResponse = await getTickets();
    if (ticketsResponse.responseCode === 403) {
      await this.$router.push({path: "forbidden"});
    }
    this.tickets = ticketsResponse.tickets
  }
}
</script>
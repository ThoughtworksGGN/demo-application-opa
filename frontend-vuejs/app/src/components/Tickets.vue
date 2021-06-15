<template>
  <div id="tickets" v-if="tickets">
    <div>
      <h3>Tickets</h3>
      <table>
        <thead>
        <td>Ticket Id</td>
        <td>Account Id</td>
        <td>Info</td>
        </thead>
        <tbody>
        <template v-for="ticket in tickets">
          <tr :key="ticket.ticketid">
            <td>{{ ticket.ticketid }}</td>
            <td>
              <router-link :to="`/account/${ticket.accountid}`">{{ ticket.accountid }}</router-link>
            </td>
            <td>{{ ticket.info }}</td>
          </tr>
        </template>
        </tbody>
      </table>
    </div>
  </div>
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
    this.tickets = ticketsResponse.tickets
  }
}
</script>
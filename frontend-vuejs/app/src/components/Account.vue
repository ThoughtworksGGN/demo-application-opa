<template>
  <div id="account-details" v-if="account">
    <div>
      <label>Account Id: </label>
      <span>{{ account.account_id }}</span>
    </div>
    <div>
      <h3>Payments</h3>
      <table>
        <thead>
        <td>Payment Id</td>
        <td>Amount</td>
        </thead>
        <tbody>
        <template v-for="payment in account.payments">
          <tr :key="payment.paymentid">
            <td>{{ payment.paymentid }}</td>
            <td>{{ payment.amount }}</td>
          </tr>
        </template>
        </tbody>
      </table>
    </div>
    <div>
      <h3>Notifications</h3>
      <table>
        <thead>
        <td>Notification Id</td>
        <td>Info</td>
        </thead>
        <tbody>
        <template v-for="notification in account.notifications">
          <tr :key="notification.notificationid">
            <td>{{ notification.notificationid }}</td>
            <td>{{ notification.info }}</td>
          </tr>
        </template>
        </tbody>
      </table>
    </div>
    <div>
      <h3>Tickets</h3>
      <table>
        <thead>
        <td>Ticket Id</td>
        <td>Info</td>
        </thead>
        <tbody>
        <template v-for="ticket in account.tickets">
          <tr :key="ticket.ticketid">
            <td>{{ ticket.ticketid }}</td>
            <td>{{ ticket.info }}</td>
          </tr>
        </template>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script lang="ts">
import {Vue} from "vue-property-decorator";
import Component from "vue-class-component";
import {account} from "@/components/types";
import {getAccount} from "@/components/service";

@Component
export default class AccountDetails extends Vue{
 account: account | null = null;
  async beforeMount(){
    const accountId = this.$route.params.accountid;
    const accountForUserResponse = await getAccount(accountId);
    this.account = accountForUserResponse.account;
  }
}
</script>
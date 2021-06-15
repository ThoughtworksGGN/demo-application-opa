export interface loginRequest {
    name: string,
    password: string
}

export interface user {
    userid: string,
    name: string,
    role: string
}

export interface notification {
    notificationid: string,
    accountid: string,
    info: string
}

export interface ticket {
    ticketid: string,
    assignee: string,
    accountid: string,
    info: string
}

export interface payment {
    paymentid: string,
    accountid: string,
    amount: number
}

export interface account {
    account_id: string,
    payments: payment[],
    tickets: ticket[],
    notifications: notification[],
    user: user
}

export interface loginResponse {
    token: string
    responseCode: number,
    error: string
}

export interface accountForUserResponse {
    account_id: string,
    responseCode: number,
    error: string
}

export interface ticketsResponse {
    tickets: ticket[],
    responseCode: number,
    error: string
}

export interface accountResponse {
    account: account,
    responseCode: number,
    error: string
}

export const USERTOKEN = 'demo-app-opa-token'
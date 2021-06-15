import axios, {AxiosRequestConfig, AxiosResponse} from 'axios';
import {
    account,
    accountForUserResponse,
    accountResponse,
    loginRequest, loginResponse,
    ticketsResponse,
    user
} from "@/components/types";
import jwtDecode from "jwt-decode";

const baseUrl = "http://localhost:8124"

export const login = async(data: loginRequest):Promise<loginResponse> => {
    const loginResponse: loginResponse = {
        token: '',
        responseCode: 0,
        error: ''
    }
    const requestConfig = {} as AxiosRequestConfig;
    const url = `${baseUrl}/login`;
    const response = await axios.post(url,data,requestConfig);

    if (response.data && response.data.token) {
        localStorage.setItem('demo-app-opa-token', response.data.token);
        loginResponse.token = response.data.token;
    }

    loginResponse.responseCode = response.status;
    loginResponse.error = response.statusText;

    return loginResponse;

};

export const getUserDataFromToken = async():Promise<user> => {
    const token = localStorage.getItem('demo-app-opa-token');
    const user: user = {
        name: "", role: "", userid: ""
    };
    if (token) {
        const decodedToken:any = jwtDecode(token);
        user.name = decodedToken.user_name;
        user.role = decodedToken.role;
        user.userid = decodedToken.user_id;
    }
    return user;
}

export const getAccountForUser = async():Promise<accountForUserResponse> => {
    const requestConfig = {} as AxiosRequestConfig;
    let accountForUserResponse: accountForUserResponse = {
        account_id: "",
        responseCode: 0,
        error: ""
    };
    const url = `${baseUrl}/account`;
    const token = localStorage.getItem('demo-app-opa-token');
    requestConfig.headers = { Authorization: token }
    const response = await axios.get(url,requestConfig);

    if (response.data && response.data.account_id) {
        accountForUserResponse = response.data;
    }
    accountForUserResponse.responseCode = response.status;
    accountForUserResponse.error = response.statusText;

    return accountForUserResponse;
}

export const getTickets = async():Promise<ticketsResponse> => {
    const url = `${baseUrl}/tickets`;
    const token = localStorage.getItem('demo-app-opa-token');
    const requestConfig = {} as AxiosRequestConfig;
    const ticketsResponse: ticketsResponse = {
        tickets: [],
        responseCode: 0,
        error: ""
    }
    requestConfig.headers = { Authorization: token };
    const response = await axios.get(url,requestConfig);

    if (response.data && response.data.tickets) {
        ticketsResponse.tickets = response.data.tickets
    }

    ticketsResponse.responseCode = response.status;
    ticketsResponse.error = response.statusText;

    return ticketsResponse
}

export const getAccount = async(accountId: string):Promise<accountResponse>  => {
    // @ts-ignore
    const accountResponse: accountResponse = {};
    const url = `${baseUrl}/accounts/${accountId}`;
    const token = localStorage.getItem('demo-app-opa-token');
    const requestConfig = {} as AxiosRequestConfig;
    requestConfig.headers = { Authorization: token };

    const response = await axios.get(url,requestConfig);

    if (response.data && response.data.account_id) {
        accountResponse.account = response.data;
    }

    accountResponse.responseCode = response.status;
    accountResponse.error = response.statusText;

    return accountResponse;
}

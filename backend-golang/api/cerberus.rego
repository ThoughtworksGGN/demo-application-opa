package api

default viewAccount = false

customer = true {
    user := data.users[input.userid]
    user.role == "CUSTOMER"
}

support  = true {
    user := data.users[input.userid]
    user.role == "SUPPORT"
}

accountOwner = true {
  userid := input.userid
  accountid := input.accountid
  customer
  data.accounts[accountid].userid == userid
}

viewAccount = true {
	accountOwner
} else = true {
  support
    data.tickets[idx].accountid = input.accountid
    data.tickets[idx].assignee == input.userid
}
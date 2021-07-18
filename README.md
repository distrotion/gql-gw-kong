# gql-gw-kong


/auth
{"Query":"mutation {login(input: {username: \"parin@mail.com\", password: \"123456789\",attempt:1,signintime:\"\"}){Status Sessionid Userid Clinicid Roleid Accountstatus Sessionlifetime}}"}


/noti
{"Query":"mutation{callNoti(input:{userid: \"1-60f2318e16926f9cbdf068e8\"}){notilist{name type related doctor appointment}}}"}
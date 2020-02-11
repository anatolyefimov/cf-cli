package utils

//Host of cf
const Host = `https://codeforces.com/`

//DumpName user information
const DumpName = `.cfrc`

//Help text
const Help = `
сf login                                Login to Codeforces. Note that password is entered in silence mode.
cf enter <contest-id>                   Entering the contest with this id.
сf submit <id|index> <path-to-source>   The command send the solution to the problemset, if before that the user entered the contest using 'cf enter', otherwise  it will send by the index of the corresponding contest.
cf help                                 Get help
`

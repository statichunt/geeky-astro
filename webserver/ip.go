package main

var url = "http://api-production-netinfo.hyvesdp.com/v2/network?ip_address=197.94.4.80"
var api_header_key = "x-api-key"
var secret = "NbVQ^82IM3fOH^hd"

//Offnet example
// {
//     "data": {
//         "type": "network",
//         "attributes": {
//             "network": "OPTINET",
//             "mcc": null,
//             "mnc": null,
//             "verified_onnet": false,
//             "country_iso_code": "ZA",
//             "range": "197.94.0.0/21"
//         }
//     },
//     "meta": {}
// }

//Onnet example
// {
//     "data": {
//         "type": "network",
//         "attributes": {
//             "network": "Telkom Internet",
//             "mcc": "655",
//             "mnc": "02",
//             "verified_onnet": false,
//             "country_iso_code": "ZA",
//             "range": "102.249.4.0/23"
//         }
//     },
//     "meta": {}
// }

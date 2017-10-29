# `diags://` interception
because why not? for ios 6.1.6, format appears to have changed in 7.

## this is really cool how can i use it
it's actually pretty easy. or so i think
grab yourself 3 things:
- dns server, configuration.apple.com + iosdiags.apple.com pointing to your local ip
- sslstrip
- some certs

```
openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```
put these in the `files` dir

at this point, on your ios device, install sslstrip and change your dns server.
go to safari, `diags://` and enter any ticket number you want.

example output:
```
> sudo go run *.go
Password:
2017/10/29 13:25:59 GET configuration.apple.com/configurations/retail/mobileBehaviorScan_1.1.plist
iOS%20Diagnostics/1.0 CFNetwork/609.1.4 Darwin/13.0.0
2017/10/29 13:25:59 GET iosdiags.apple.com:80/MR3Server/GetOptinText
iOS%20Diagnostics/1.0 CFNetwork/609.1.4 Darwin/13.0.0
2017/10/29 13:26:02 GET iosdiags.apple.com:80/MR3Server/ValidateTicket?ticket_number=888888
iOS%20Diagnostics/1.0 CFNetwork/609.1.4 Darwin/13.0.0
2017/10/29 13:26:03 POST iosdiags.apple.com:80/MR3Server/MR3Post
iOS%20Diagnostics/1.0 CFNetwork/609.1.4 Darwin/13.0.0
2017/10/29 13:26:03 Values:
2017/10/29 13:26:03 result => [okay]
2017/10/29 13:26:03 serial_number => [nice try]
2017/10/29 13:26:03 ticket_number => [whatever ticket number you entered]
2017/10/29 13:26:03 device_version => [6.1.6]
2017/10/29 13:26:03 FullChargeCapacity => [930]
2017/10/29 13:26:03 battery_level => [56]
2017/10/29 13:26:03 device_name => [iPod Touch I Think]
2017/10/29 13:26:03 CycleCount => [0]
2017/10/29 13:26:03 DesignCapacity => [930]
2017/10/29 13:26:03 properties => [a bunch of debugging stuff I omit for brevity]
2017/10/29 13:26:03 device_type => [iPod4,1]
2017/10/29 13:26:03 application_version => [1.0]
2017/10/29 13:26:03 Files:
2017/10/29 13:26:03 Working with log_archive
2017/10/29 13:26:03 /tmp/com.apple.behaviorscan.Rg269SJA => files/_tmp_com.apple.behaviorscan.Rg269SJA
```

log_archive is sent as a gzipped file, which this script gunzips and places in files.

merry christmas, happy new year
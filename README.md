# IoT Registery

A simple REST service to act as a registery for IoT devices.

----

/
        GET  - Return basic info, including:
                * current time in UTC
                * ntp server IP
                * environment
                  - timezone
                  - realm name
                  - location name
                * crypto proof - CA signed something.

/reg/
        GET - return a list of UUIDs
        POST - registration (no uuid), incl 8266 serial #
               return:  full registration record w/ UUID

/reg/<uuid>
        GET - return full record
        DEL - destroy registration (retain config <-> s/n)

/reg/<uuid>/ping
        GET - keep alive / not sure if/what to return.
              - indicate if a config update occurred.

/sn/
        GET - list all registered s/n's

/sn/<serial-number>
        GET - return config for s/n
        PUT - set config for s/n
        DEL - destroy config for s/n

----

Configuration record:

{
  id:  <uuid>       ## unique per config version
  sn:  <serial-number>
  ver: <version>    ## monotonic increasing
  ts:  <timestamp>  ## ? ISO-8601 ?

  cfg:  {map of key / value _strings_}
}

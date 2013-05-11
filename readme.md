synopsis
========

    http_plug <port>

    port is an integer in 1..65535 range. The `http_plug` process will listen on that TCP port and respond with request data to http requests.

purpose
=======

When troubleshooting network configuration I needed a porocess I knew was listening on a specific TCP port to test firewall configuration. Windows did not seem to have a simple way of doing that.
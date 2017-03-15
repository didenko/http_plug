synopsis
========

    http_plug <port>

_port_ is an integer in 1..65535 range. The `http_plug` process will listen on that TCP port and echo back with some data to HTTP requests.

purpose
=======

When troubleshooting network configuration I needed a process I knew was listening on a specific TCP port to test firewall configuration. Windows did not seem to have a simple way of doing that.

binaries
========

Only `darwin/amd64` and `windows/amd64` pre-build binaries are tested.

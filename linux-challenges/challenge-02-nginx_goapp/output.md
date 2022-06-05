# KodeKloud - Linux Challenge - Challenge 2 Output

- [x] Inspect OS

```
[bob@centos-host ~]$ uname -r
5.4.0-1028-gcp
```

```
[bob@centos-host ~]$ cat /etc/os-release
NAME="CentOS Stream"
VERSION="8"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="8"
PLATFORM_ID="platform:el8"
PRETTY_NAME="CentOS Stream 8"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:8"
HOME_URL="https://centos.org/"
BUG_REPORT_URL="https://bugzilla.redhat.com/"
REDHAT_SUPPORT_PRODUCT="Red Hat Enterprise Linux 8"
REDHAT_SUPPORT_PRODUCT_VERSION="CentOS Stream"
```

- [x] Install "nginx" package.
- [x] Install "firewalld" package.

```
[bob@centos-host ~]$ sudo yum install -y tmux nginx firewalld
CentOS Stream 8 - AppStream                                                                                                                                0.0  B/s |   0  B     00:00
Errors during downloading metadata for repository 'appstream':
  - Curl error (6): Couldn't resolve host name for http://mirrorlist.centos.org/?release=8-stream&arch=x86_64&repo=AppStream&infra=stock [Could not resolve host: mirrorlist.centos.org]
  - Error: Failed to download metadata for repo 'appstream': Cannot prepare internal mirrorlist: Curl error (6): Couldn't resolve host name for http://mirrorlist.centos.org/?release=8-stream&arch=x86_64&repo=AppStream&infra=stock [Could not resolve host: mirrorlist.centos.org]
```

- [x] Troubleshoot the issues with "yum/dnf" and make sure you are able to install the packages on "centos-host".

Notes:
- No host (bind-utils) or ping (iputils) commands.
- Since curl can't resolve the name, it's probably a DNS issue.

```
[bob@centos-host ~]$ cat /etc/resolv.conf
search us-central1-a.c.kk-lab-prod.internal c.kk-lab-prod.internal google.internal
options ndots:0
```

```
[bob@centos-host ~]$ sudoedit /etc/resolv.conf
```

```
[bob@centos-host ~]$ tail -v /etc/resolv.conf
==> /etc/resolv.conf <==
search us-central1-a.c.kk-lab-prod.internal c.kk-lab-prod.internal google.internal
nameserver 8.8.8.8
options ndots:0
```

- [x] Install "nginx" package.
- [x] Install "firewalld" package.

Notes:
- Also add tmux and net-tools (netstat). Output for installing net-tools not shown below.

```
[bob@centos-host ~]$ sudo yum install -y tmux nginx firewalld
CentOS Stream 8 - AppStream                                                                                                                                 15 MB/s |  22 MB     00:01
CentOS Stream 8 - BaseOS                                                                                                                                    19 MB/s |  22 MB     00:01
CentOS Stream 8 - Extras                                                                                                                                    44 kB/s |  18 kB     00:00
CentOS Stream 8 - Extras common packages                                                                                                                   8.9 kB/s | 4.3 kB     00:00
Dependencies resolved.
===========================================================================================================================================================================================
 Package                                            Architecture                  Version                                                           Repository                        Size
===========================================================================================================================================================================================
Installing:
 firewalld                                          noarch                        0.9.3-13.el8                                                      baseos                           503 k
 nginx                                              x86_64                        1:1.14.1-9.module_el8.0.0+1060+3ab382d3                           appstream                        570 k
 tmux                                               x86_64                        2.7-1.el8                                                         baseos                           317 k
Installing dependencies:
 firewalld-filesystem                               noarch                        0.9.3-13.el8                                                      baseos                            78 k
 ipset                                              x86_64                        7.1-1.el8                                                         baseos                            45 k
 ipset-libs                                         x86_64                        7.1-1.el8                                                         baseos                            71 k
 nginx-all-modules                                  noarch                        1:1.14.1-9.module_el8.0.0+1060+3ab382d3                           appstream                         23 k
 nginx-filesystem                                   noarch                        1:1.14.1-9.module_el8.0.0+1060+3ab382d3                           appstream                         24 k
 nginx-mod-http-image-filter                        x86_64                        1:1.14.1-9.module_el8.0.0+1060+3ab382d3                           appstream                         35 k
 nginx-mod-http-perl                                x86_64                        1:1.14.1-9.module_el8.0.0+1060+3ab382d3                           appstream                         46 k
 nginx-mod-http-xslt-filter                         x86_64                        1:1.14.1-9.module_el8.0.0+1060+3ab382d3                           appstream                         33 k
 nginx-mod-mail                                     x86_64                        1:1.14.1-9.module_el8.0.0+1060+3ab382d3                           appstream                         64 k
 nginx-mod-stream                                   x86_64                        1:1.14.1-9.module_el8.0.0+1060+3ab382d3                           appstream                         85 k
 python3-firewall                                   noarch                        0.9.3-13.el8                                                      baseos                           434 k
 python3-nftables                                   x86_64                        1:0.9.3-25.el8                                                    baseos                            30 k
 python3-slip                                       noarch                        0.6.4-13.el8                                                      baseos                            39 k
 python3-slip-dbus                                  noarch                        0.6.4-13.el8                                                      baseos                            39 k
Enabling module streams:
 nginx                                                                            1.14

Transaction Summary
===========================================================================================================================================================================================
Install  17 Packages

Total download size: 2.4 M
Installed size: 7.2 M
Downloading Packages:
(1/17): nginx-all-modules-1.14.1-9.module_el8.0.0+1060+3ab382d3.noarch.rpm                                                                                 188 kB/s |  23 kB     00:00
(2/17): nginx-filesystem-1.14.1-9.module_el8.0.0+1060+3ab382d3.noarch.rpm                                                                                  160 kB/s |  24 kB     00:00
(3/17): nginx-mod-http-image-filter-1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64.rpm                                                                       606 kB/s |  35 kB     00:00
(4/17): nginx-mod-http-perl-1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64.rpm                                                                               492 kB/s |  46 kB     00:00
(5/17): nginx-mod-http-xslt-filter-1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64.rpm                                                                        386 kB/s |  33 kB     00:00
(6/17): nginx-1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64.rpm                                                                                             1.8 MB/s | 570 kB     00:00
(7/17): nginx-mod-mail-1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64.rpm                                                                                    767 kB/s |  64 kB     00:00
(8/17): nginx-mod-stream-1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64.rpm                                                                                  961 kB/s |  85 kB     00:00
(9/17): firewalld-filesystem-0.9.3-13.el8.noarch.rpm                                                                                                       392 kB/s |  78 kB     00:00
(10/17): ipset-7.1-1.el8.x86_64.rpm                                                                                                                        221 kB/s |  45 kB     00:00
(11/17): ipset-libs-7.1-1.el8.x86_64.rpm                                                                                                                   998 kB/s |  71 kB     00:00
(12/17): firewalld-0.9.3-13.el8.noarch.rpm                                                                                                                 1.3 MB/s | 503 kB     00:00
(13/17): python3-nftables-0.9.3-25.el8.x86_64.rpm                                                                                                          304 kB/s |  30 kB     00:00
(14/17): python3-slip-0.6.4-13.el8.noarch.rpm                                                                                                              519 kB/s |  39 kB     00:00
(15/17): python3-slip-dbus-0.6.4-13.el8.noarch.rpm                                                                                                         501 kB/s |  39 kB     00:00
(16/17): python3-firewall-0.9.3-13.el8.noarch.rpm                                                                                                          1.7 MB/s | 434 kB     00:00
(17/17): tmux-2.7-1.el8.x86_64.rpm                                                                                                                         3.0 MB/s | 317 kB     00:00
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                                                                      2.6 MB/s | 2.4 MB     00:00
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                                                                   1/1
  Installing       : python3-slip-0.6.4-13.el8.noarch                                                                                                                                 1/17
warning: Unable to get systemd shutdown inhibition lock: Unit systemd-logind.service is masked.

  Installing       : python3-slip-dbus-0.6.4-13.el8.noarch                                                                                                                            2/17
  Installing       : python3-nftables-1:0.9.3-25.el8.x86_64                                                                                                                           3/17
  Installing       : python3-firewall-0.9.3-13.el8.noarch                                                                                                                             4/17
  Installing       : ipset-libs-7.1-1.el8.x86_64                                                                                                                                      5/17
  Running scriptlet: ipset-libs-7.1-1.el8.x86_64                                                                                                                                      5/17
  Installing       : ipset-7.1-1.el8.x86_64                                                                                                                                           6/17
  Installing       : firewalld-filesystem-0.9.3-13.el8.noarch                                                                                                                         7/17
  Running scriptlet: nginx-filesystem-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.noarch                                                                                                  8/17
  Installing       : nginx-filesystem-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.noarch                                                                                                  8/17
  Installing       : nginx-mod-http-image-filter-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                       9/17
  Running scriptlet: nginx-mod-http-image-filter-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                       9/17
  Installing       : nginx-mod-http-perl-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                              10/17
  Running scriptlet: nginx-mod-http-perl-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                              10/17
  Installing       : nginx-mod-http-xslt-filter-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                       11/17
  Running scriptlet: nginx-mod-http-xslt-filter-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                       11/17
  Installing       : nginx-mod-mail-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                                   12/17
  Running scriptlet: nginx-mod-mail-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                                   12/17
  Installing       : nginx-all-modules-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.noarch                                                                                                13/17
  Installing       : nginx-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                                            14/17
  Running scriptlet: nginx-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                                            14/17
  Installing       : nginx-mod-stream-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                                 15/17
  Running scriptlet: nginx-mod-stream-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                                 15/17
  Installing       : firewalld-0.9.3-13.el8.noarch                                                                                                                                   16/17
  Running scriptlet: firewalld-0.9.3-13.el8.noarch                                                                                                                                   16/17
  Installing       : tmux-2.7-1.el8.x86_64                                                                                                                                           17/17
  Running scriptlet: tmux-2.7-1.el8.x86_64                                                                                                                                           17/17
  Verifying        : nginx-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                                             1/17
  Verifying        : nginx-all-modules-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.noarch                                                                                                 2/17
  Verifying        : nginx-filesystem-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.noarch                                                                                                  3/17
  Verifying        : nginx-mod-http-image-filter-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                       4/17
  Verifying        : nginx-mod-http-perl-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                               5/17
  Verifying        : nginx-mod-http-xslt-filter-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                        6/17
  Verifying        : nginx-mod-mail-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                                    7/17
  Verifying        : nginx-mod-stream-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                                                                                  8/17
  Verifying        : firewalld-0.9.3-13.el8.noarch                                                                                                                                    9/17
  Verifying        : firewalld-filesystem-0.9.3-13.el8.noarch                                                                                                                        10/17
  Verifying        : ipset-7.1-1.el8.x86_64                                                                                                                                          11/17
  Verifying        : ipset-libs-7.1-1.el8.x86_64                                                                                                                                     12/17
  Verifying        : python3-firewall-0.9.3-13.el8.noarch                                                                                                                            13/17
  Verifying        : python3-nftables-1:0.9.3-25.el8.x86_64                                                                                                                          14/17
  Verifying        : python3-slip-0.6.4-13.el8.noarch                                                                                                                                15/17
  Verifying        : python3-slip-dbus-0.6.4-13.el8.noarch                                                                                                                           16/17
  Verifying        : tmux-2.7-1.el8.x86_64                                                                                                                                           17/17

Installed:
  firewalld-0.9.3-13.el8.noarch                                                            firewalld-filesystem-0.9.3-13.el8.noarch
  ipset-7.1-1.el8.x86_64                                                                   ipset-libs-7.1-1.el8.x86_64
  nginx-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                                     nginx-all-modules-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.noarch
  nginx-filesystem-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.noarch                          nginx-mod-http-image-filter-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64
  nginx-mod-http-perl-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                       nginx-mod-http-xslt-filter-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64
  nginx-mod-mail-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64                            nginx-mod-stream-1:1.14.1-9.module_el8.0.0+1060+3ab382d3.x86_64
  python3-firewall-0.9.3-13.el8.noarch                                                     python3-nftables-1:0.9.3-25.el8.x86_64
  python3-slip-0.6.4-13.el8.noarch                                                         python3-slip-dbus-0.6.4-13.el8.noarch
  tmux-2.7-1.el8.x86_64

Complete!
```

- [x] Start and Enable "firewalld" service.

```
[bob@centos-host ~]$ sudo systemctl enable firewalld
```

```
[bob@centos-host ~]$ sudo systemctl start firewalld
```

- [x] Add firewall rules to allow only incoming port "22", "80" and "8081".

```
[bob@centos-host ~]$ sudo firewall-cmd --permanent --zone=public --list-ports
```

```
[bob@centos-host ~]$ sudo firewall-cmd --permanent --zone=public --add-port=22/tcp
success
```

```
[bob@centos-host ~]$ sudo firewall-cmd --permanent --zone=public --add-port=80/tcp
success
```

```
[bob@centos-host ~]$ sudo firewall-cmd --permanent --zone=public --add-port=8081/tcp
success
```

- [x] The firewall rules must be permanent and effective immediately.

```
[bob@centos-host ~]$ sudo firewall-cmd --reload
success
```

```
[bob@centos-host ~]$ sudo firewall-cmd --permanent --zone=public --list-ports
22/tcp 80/tcp 8081/tcp
```


- [x] Configure Nginx as a reverse proxy for the GoApp so that we can access the GoApp on port "80"

```
[bob@centos-host ~]$ sudo bash -c 'cat > /etc/nginx/conf.d/goapp.conf' <<EOF
# /etc/nginx/conf.d/goapp.conf

server {
    listen 80;
	# we don't know the name of the server
    server_name *;

    location / {
        proxy_pass http://localhost:8081;
    }
}
EOF
```

```
[bob@centos-host ~]$ cat /etc/nginx/conf.d/goapp.conf
# /etc/nginx/conf.d/goapp.conf

server {
    listen 80;
	# we don't know the name of the server
    server_name *;

    location / {
        proxy_pass http://localhost:8081;
    }
}
```

- [x] Start and Enable "nginx" service.

```
[bob@centos-host ~]$ sudo systemctl enable nginx
Created symlink /etc/systemd/system/multi-user.target.wants/nginx.service → /usr/lib/systemd/system/nginx.service.
```

```
[bob@centos-host ~]$ sudo systemctl start nginx
```

```
[bob@centos-host ~]$ curl -v http://localhost/
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 80 (#0)
> GET / HTTP/1.1
> Host: localhost
> User-Agent: curl/7.61.1
> Accept: */*
>
< HTTP/1.1 502 Bad Gateway
< Server: nginx/1.14.1
< Date: Sun, 05 Jun 2022 17:48:47 GMT
< Content-Type: text/html
< Content-Length: 173
< Connection: keep-alive
<
<html>
<head><title>502 Bad Gateway</title></head>
<body bgcolor="white">
<center><h1>502 Bad Gateway</h1></center>
<hr><center>nginx/1.14.1</center>
</body>
</html>
* Connection #0 to host localhost left intact
```

- [x] Start GoApp by running the "nohup go run main.go &" command from "/home/bob/go-app/" directory, it can take few seconds to start.

Notes:
- Why didn't it ask us to add a service for this app? (Forgot to complain about that in the feedback form.)

```
[bob@centos-host ~]$ cd go-app/
```

```
[bob@centos-host go-app]$ ls
CONTRIBUTING.md  README.md                application.k8s.yml  controller        go.mod  main.go     model       router   test                   zaplogger.docker.yml
Dockerfile       application.develop.yml  config               docs              go.sum  middleware  public      service  util                   zaplogger.k8s.yml
LICENSE          application.docker.yml   container            go-webapp-sample  logger  migration   repository  session  zaplogger.develop.yml
```

```
[bob@centos-host go-app]$ nohup go run main.go &
[1] 2769
[bob@centos-host go-app]$ nohup: ignoring input and appending output to 'nohup.out'
```

```
[bob@centos-host go-app]$ tail -v nohup.out
==> nohup.out <==
go: downloading golang.org/x/text v0.3.7
go: downloading github.com/jackc/pgpassfile v1.0.0
go: downloading golang.org/x/sys v0.0.0-20220114195835-da31bd327af9
go: downloading github.com/go-openapi/swag v0.19.15
go: downloading github.com/go-openapi/jsonreference v0.19.6
go: downloading github.com/go-openapi/jsonpointer v0.19.5
go: downloading github.com/PuerkitoBio/purell v1.1.1
go: downloading github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578
go: downloading github.com/mailru/easyjson v0.7.7
go: downloading github.com/josharian/intern v1.0.0
```

```
[bob@centos-host go-app]$ curl -v http://localhost/
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 80 (#0)
> GET / HTTP/1.1
> Host: localhost
> User-Agent: curl/7.61.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Server: nginx/1.14.1
< Date: Sun, 05 Jun 2022 17:53:01 GMT
< Content-Type: text/html; charset=utf-8
< Content-Length: 1865
< Connection: keep-alive
< Accept-Ranges: bytes
< Last-Modified: Sun, 05 Jun 2022 17:10:25 GMT
< Vary: Origin
<
<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><meta http-equiv="X-UA-Compatible" content="IE=edge"><meta name="viewport" content="width=device-width,initial-scale=1"><!--[if IE]><link rel="icon" href="/favicon.ico"><![endif]--><title>vuejs-webapp-sample</title><link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:300,400,500,700,400italic"><link rel="stylesheet" href="//fonts.googleapis.com/icon?family=Material+Icons"><link href="/css/app.750b60b0.css" rel="preload" as="style"><link href="/css/chunk-vendors.533831d3.css" rel="preload" as="style"><link href="/js/app.dbc5a974.js" rel="preload" as="script"><link href="/js/chunk-vendors.0cedba66.js" rel="preload" as="script"><link href="/css/chunk-vendors.533831d3.css" rel="stylesheet"><link href="/css/app.750b60b0.css" rel="stylesheet"><link rel="icon" type="image/png" sizes="32x32" href="/img/icons/favicon-32x32.png"><link rel="icon" type="image/png" sizes="16x16" href="/img/icons/favicon-16x16.png"><link rel="manifest" href="/manif* Connection #0 to host localhost left intact
est.json"><meta name="theme-color" content="#4DBA87"><meta name="apple-mobile-web-app-capable" content="no"><meta name="apple-mobile-web-app-status-bar-style" content="default"><meta name="apple-mobile-web-app-title" content="vuejs-webapp-sample"><link rel="apple-touch-icon" href="/img/icons/apple-touch-icon-152x152.png"><link rel="mask-icon" href="/img/icons/safari-pinned-tab.svg" color="#4DBA87"><meta name="msapplication-TileImage" content="/img/icons/msapplication-icon-144x144.png"><meta name="msapplication-TileColor" content="#000000"></head><body><noscript><strong>We're sorry but vuejs-webapp-sample doesn't work properly without JavaScript enabled. Please enable it to continue.</strong></noscript><div id="app"></div><script src="/js/chunk-vendors.0cedba66.js"></script><script src="/js/app.dbc5a974.js"></script></body></html>
```

```
[bob@centos-host go-app]$ sudo netstat -tlnp
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name
tcp        0      0 127.0.0.11:42855        0.0.0.0:*               LISTEN      -
tcp        0      0 0.0.0.0:111             0.0.0.0:*               LISTEN      1/init
tcp        0      0 0.0.0.0:80              0.0.0.0:*               LISTEN      2701/nginx: master
tcp        0      0 0.0.0.0:8080            0.0.0.0:*               LISTEN      937/ttyd
tcp        0      0 192.168.122.1:53        0.0.0.0:*               LISTEN      1394/dnsmasq
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      1034/sshd
tcp6       0      0 :::111                  :::*                    LISTEN      1/init
tcp6       0      0 :::80                   :::*                    LISTEN      2701/nginx: master
tcp6       0      0 :::8081                 :::*                    LISTEN      6014/main
tcp6       0      0 :::22                   :::*                    LISTEN      1034/sshd
```

- [x] bob is able to login into GoApp using username "test" and password "test"

Notes:
- Sign-in from curl not supported?
- Can't get past the Red Hat default nginx page at: https://80-port-2472b0cc285c4c23.labs.kodekloud.com/
- All the tests passed.

```
[bob@centos-host go-app]$ curl -v --user test:test http://localhost/
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 80 (#0)
* Server auth using Basic with user 'test'
> GET / HTTP/1.1
> Host: localhost
> Authorization: Basic dGVzdDp0ZXN0
> User-Agent: curl/7.61.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Server: nginx/1.14.1
< Date: Sun, 05 Jun 2022 17:53:54 GMT
< Content-Type: text/html; charset=utf-8
< Content-Length: 1865
< Connection: keep-alive
< Accept-Ranges: bytes
< Last-Modified: Sun, 05 Jun 2022 17:10:25 GMT
< Vary: Origin
<
<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><meta http-equiv="X-UA-Compatible" content="IE=edge"><meta name="viewport" content="width=device-width,initial-scale=1"><!--[if IE]><link rel="icon" href="/favicon.ico"><![endif]--><title>vuejs-webapp-sample</title><link rel="stylesheet" href="//fonts.googleapis.com/css?family=Roboto:300,400,500,700,400italic"><link rel="stylesheet" href="//fonts.googleapis.com/icon?family=Material+Icons"><link href="/css/app.750b60b0.css" rel="preload" as="style"><link href="/css/chunk-vendors.533831d3.css" rel="preload" as="style"><link href="/js/app.dbc5a974.js" rel="preload" as="script"><link href="/js/chunk-vendors.0cedba66.js" rel="preload" as="script"><link href="/css/chunk-vendors.533831d3.css" rel="stylesheet"><link href="/css/app.750b60b0.css" rel="stylesheet"><link rel="icon" type="image/png" sizes="32x32" href="/img/icons/favicon-32x32.png"><link rel="icon" type="image/png" sizes="16x16" href="/img/icons/favicon-16x16.png"><link rel="manifest" href="/manif* Connection #0 to host localhost left intact
est.json"><meta name="theme-color" content="#4DBA87"><meta name="apple-mobile-web-app-capable" content="no"><meta name="apple-mobile-web-app-status-bar-style" content="default"><meta name="apple-mobile-web-app-title" content="vuejs-webapp-sample"><link rel="apple-touch-icon" href="/img/icons/apple-touch-icon-152x152.png"><link rel="mask-icon" href="/img/icons/safari-pinned-tab.svg" color="#4DBA87"><meta name="msapplication-TileImage" content="/img/icons/msapplication-icon-144x144.png"><meta name="msapplication-TileColor" content="#000000"></head><body><noscript><strong>We're sorry but vuejs-webapp-sample doesn't work properly without JavaScript enabled. Please enable it to continue.</strong></noscript><div id="app"></div><script src="/js/chunk-vendors.0cedba66.js"></script><script src="/js/app.dbc5a974.js"></script></body></htmlg
```

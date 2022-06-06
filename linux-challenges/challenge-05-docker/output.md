# KodeKloud - Linux Challenge - Challenge 5 Output

- [x] Inspect OS

```
[bob@centos-host ~]$ uname -r
4.18.0-240.1.1.el8_3.x86_64
```

```
[bob@centos-host ~]$ cat /etc/os-release
ME="CentOS Linux"
VERSION="8"
ID="centos"
ID_LIKE="rhel fedora"
VERSION_ID="8"
PLATFORM_ID="platform:el8"
PRETTY_NAME="CentOS Linux 8"
ANSI_COLOR="0;31"
CPE_NAME="cpe:/o:centos:centos:8"
HOME_URL="https://centos.org/"
BUG_REPORT_URL="https://bugs.centos.org/"
CENTOS_MANTISBT_PROJECT="CentOS-8"
CENTOS_MANTISBT_PROJECT_VERSION="8"
```

- [x] Install "mariadb" database server on this server and "start/enable" its service.

```
[bob@centos-host ~]$ sudo yum install -y mariadb
Last metadata expiration check: 0:13:49 ago on Sun 05 Jun 2022 11:53:02 PM UTC.
Module yaml error: Unexpected key in data: static_context [line 9 col 3]
Module yaml error: Unexpected key in data: static_context [line 9 col 3]
Dependencies resolved.
===========================================================================================================================================================================================
 Package                                            Architecture                   Version                                                         Repository                         Size
===========================================================================================================================================================================================
Installing:
 mariadb                                            x86_64                         3:10.3.28-1.module_el8.3.0+757+d382997d                         appstream                         6.0 M
Installing dependencies:
 mariadb-common                                     x86_64                         3:10.3.28-1.module_el8.3.0+757+d382997d                         appstream                          64 k
 mariadb-connector-c                                x86_64                         3.1.11-2.el8_3                                                  appstream                         200 k
 mariadb-connector-c-config                         noarch                         3.1.11-2.el8_3                                                  appstream                          15 k
 perl-Carp                                          noarch                         1.42-396.el8                                                    baseos                             30 k
 perl-Data-Dumper                                   x86_64                         2.167-399.el8                                                   baseos                             58 k
 perl-Digest                                        noarch                         1.17-395.el8                                                    appstream                          27 k
 perl-Digest-MD5                                    x86_64                         2.55-396.el8                                                    appstream                          37 k
 perl-Encode                                        x86_64                         4:2.97-3.el8                                                    baseos                            1.5 M
 perl-Errno                                         x86_64                         1.28-420.el8                                                    baseos                             76 k
 perl-Exporter                                      noarch                         5.72-396.el8                                                    baseos                             34 k
 perl-File-Path                                     noarch                         2.15-2.el8                                                      baseos                             38 k
 perl-File-Temp                                     noarch                         0.230.600-1.el8                                                 baseos                             63 k
 perl-Getopt-Long                                   noarch                         1:2.50-4.el8                                                    baseos                             63 k
 perl-HTTP-Tiny                                     noarch                         0.074-1.el8                                                     baseos                             58 k
 perl-IO                                            x86_64                         1.38-420.el8                                                    baseos                            142 k
 perl-MIME-Base64                                   x86_64                         3.15-396.el8                                                    baseos                             31 k
 perl-Net-SSLeay                                    x86_64                         1.88-1.module_el8.3.0+410+ff426aa3                              appstream                         379 k
 perl-PathTools                                     x86_64                         3.74-1.el8                                                      baseos                             90 k
 perl-Pod-Escapes                                   noarch                         1:1.07-395.el8                                                  baseos                             20 k
 perl-Pod-Perldoc                                   noarch                         3.28-396.el8                                                    baseos                             86 k
 perl-Pod-Simple                                    noarch                         1:3.35-395.el8                                                  baseos                            213 k
 perl-Pod-Usage                                     noarch                         4:1.69-395.el8                                                  baseos                             34 k
 perl-Scalar-List-Utils                             x86_64                         3:1.49-2.el8                                                    baseos                             68 k
 perl-Socket                                        x86_64                         4:2.027-3.el8                                                   baseos                             59 k
 perl-Storable                                      x86_64                         1:3.11-3.el8                                                    baseos                             98 k
 perl-Term-ANSIColor                                noarch                         4.06-396.el8                                                    baseos                             46 k
 perl-Term-Cap                                      noarch                         1.17-395.el8                                                    baseos                             23 k
 perl-Text-ParseWords                               noarch                         3.30-395.el8                                                    baseos                             18 k
 perl-Text-Tabs+Wrap                                noarch                         2013.0523-395.el8                                               baseos                             24 k
 perl-Time-Local                                    noarch                         1:1.280-1.el8                                                   baseos                             34 k
 perl-URI                                           noarch                         1.73-3.el8                                                      appstream                         116 k
 perl-Unicode-Normalize                             x86_64                         1.25-396.el8                                                    baseos                             82 k
 perl-constant                                      noarch                         1.33-396.el8                                                    baseos                             25 k
 perl-interpreter                                   x86_64                         4:5.26.3-420.el8                                                baseos                            6.3 M
 perl-libnet                                        noarch                         3.11-3.el8                                                      appstream                         121 k
 perl-libs                                          x86_64                         4:5.26.3-420.el8                                                baseos                            1.6 M
 perl-macros                                        x86_64                         4:5.26.3-420.el8                                                baseos                             72 k
 perl-parent                                        noarch                         1:0.237-1.el8                                                   baseos                             20 k
 perl-podlators                                     noarch                         4.11-1.el8                                                      baseos                            118 k
 perl-threads                                       x86_64                         1:2.21-2.el8                                                    baseos                             61 k
 perl-threads-shared                                x86_64                         1.58-2.el8                                                      baseos                             48 k
Installing weak dependencies:
 perl-IO-Socket-IP                                  noarch                         0.39-5.el8                                                      appstream                          47 k
 perl-IO-Socket-SSL                                 noarch                         2.066-4.module_el8.3.0+410+ff426aa3                             appstream                         298 k
 perl-Mozilla-CA                                    noarch                         20160104-7.module_el8.3.0+416+dee7bcef                          appstream                          15 k
Enabling module streams:
 mariadb                                                                           10.3
 perl                                                                              5.26
 perl-IO-Socket-SSL                                                                2.066
 perl-libwww-perl                                                                  6.34

Transaction Summary
===========================================================================================================================================================================================
Install  45 Packages

Total download size: 18 M
Installed size: 73 M
Downloading Packages:
(1/45): mariadb-common-10.3.28-1.module_el8.3.0+757+d382997d.x86_64.rpm                                                                                    369 kB/s |  64 kB     00:00
(2/45): mariadb-connector-c-3.1.11-2.el8_3.x86_64.rpm                                                                                                      1.0 MB/s | 200 kB     00:00
(3/45): mariadb-connector-c-config-3.1.11-2.el8_3.noarch.rpm                                                                                               470 kB/s |  15 kB     00:00
(4/45): perl-Digest-1.17-395.el8.noarch.rpm                                                                                                                858 kB/s |  27 kB     00:00
(5/45): perl-Digest-MD5-2.55-396.el8.x86_64.rpm                                                                                                            758 kB/s |  37 kB     00:00
(6/45): perl-IO-Socket-IP-0.39-5.el8.noarch.rpm                                                                                                            1.4 MB/s |  47 kB     00:00
(7/45): perl-IO-Socket-SSL-2.066-4.module_el8.3.0+410+ff426aa3.noarch.rpm                                                                                  6.4 MB/s | 298 kB     00:00
(8/45): perl-Mozilla-CA-20160104-7.module_el8.3.0+416+dee7bcef.noarch.rpm                                                                                  326 kB/s |  15 kB     00:00
(9/45): perl-Net-SSLeay-1.88-1.module_el8.3.0+410+ff426aa3.x86_64.rpm                                                                                      7.3 MB/s | 379 kB     00:00
(10/45): perl-URI-1.73-3.el8.noarch.rpm                                                                                                                    2.1 MB/s | 116 kB     00:00
(11/45): perl-libnet-3.11-3.el8.noarch.rpm                                                                                                                 2.9 MB/s | 121 kB     00:00
(12/45): perl-Carp-1.42-396.el8.noarch.rpm                                                                                                                 945 kB/s |  30 kB     00:00
(13/45): perl-Data-Dumper-2.167-399.el8.x86_64.rpm                                                                                                         1.5 MB/s |  58 kB     00:00
(14/45): mariadb-10.3.28-1.module_el8.3.0+757+d382997d.x86_64.rpm                                                                                           12 MB/s | 6.0 MB     00:00
(15/45): perl-Errno-1.28-420.el8.x86_64.rpm                                                                                                                1.1 MB/s |  76 kB     00:00
(16/45): perl-Encode-2.97-3.el8.x86_64.rpm                                                                                                                  13 MB/s | 1.5 MB     00:00
(17/45): perl-File-Path-2.15-2.el8.noarch.rpm                                                                                                              1.0 MB/s |  38 kB     00:00
(18/45): perl-Exporter-5.72-396.el8.noarch.rpm                                                                                                             645 kB/s |  34 kB     00:00
(19/45): perl-Getopt-Long-2.50-4.el8.noarch.rpm                                                                                                            1.9 MB/s |  63 kB     00:00
(20/45): perl-HTTP-Tiny-0.074-1.el8.noarch.rpm                                                                                                             1.8 MB/s |  58 kB     00:00
(21/45): perl-IO-1.38-420.el8.x86_64.rpm                                                                                                                   4.2 MB/s | 142 kB     00:00
(22/45): perl-MIME-Base64-3.15-396.el8.x86_64.rpm                                                                                                          882 kB/s |  31 kB     00:00
(23/45): perl-File-Temp-0.230.600-1.el8.noarch.rpm                                                                                                         540 kB/s |  63 kB     00:00
(24/45): perl-PathTools-3.74-1.el8.x86_64.rpm                                                                                                              2.8 MB/s |  90 kB     00:00
(25/45): perl-Pod-Escapes-1.07-395.el8.noarch.rpm                                                                                                          522 kB/s |  20 kB     00:00
(26/45): perl-Pod-Perldoc-3.28-396.el8.noarch.rpm                                                                                                          1.8 MB/s |  86 kB     00:00
(27/45): perl-Pod-Usage-1.69-395.el8.noarch.rpm                                                                                                            880 kB/s |  34 kB     00:00
(28/45): perl-Scalar-List-Utils-1.49-2.el8.x86_64.rpm                                                                                                      2.2 MB/s |  68 kB     00:00
(29/45): perl-Socket-2.027-3.el8.x86_64.rpm                                                                                                                2.0 MB/s |  59 kB     00:00
(30/45): perl-Pod-Simple-3.35-395.el8.noarch.rpm                                                                                                           2.1 MB/s | 213 kB     00:00
(31/45): perl-Storable-3.11-3.el8.x86_64.rpm                                                                                                               2.5 MB/s |  98 kB     00:00
(32/45): perl-Term-ANSIColor-4.06-396.el8.noarch.rpm                                                                                                       1.6 MB/s |  46 kB     00:00
(33/45): perl-Term-Cap-1.17-395.el8.noarch.rpm                                                                                                             530 kB/s |  23 kB     00:00
(34/45): perl-Text-Tabs+Wrap-2013.0523-395.el8.noarch.rpm                                                                                                  839 kB/s |  24 kB     00:00
(35/45): perl-Text-ParseWords-3.30-395.el8.noarch.rpm                                                                                                      331 kB/s |  18 kB     00:00
(36/45): perl-Time-Local-1.280-1.el8.noarch.rpm                                                                                                            928 kB/s |  34 kB     00:00
(37/45): perl-Unicode-Normalize-1.25-396.el8.x86_64.rpm                                                                                                    2.4 MB/s |  82 kB     00:00
(38/45): perl-constant-1.33-396.el8.noarch.rpm                                                                                                             668 kB/s |  25 kB     00:00
(39/45): perl-macros-5.26.3-420.el8.x86_64.rpm                                                                                                             1.8 MB/s |  72 kB     00:00
(40/45): perl-libs-5.26.3-420.el8.x86_64.rpm                                                                                                                18 MB/s | 1.6 MB     00:00
(41/45): perl-parent-0.237-1.el8.noarch.rpm                                                                                                                436 kB/s |  20 kB     00:00
(42/45): perl-podlators-4.11-1.el8.noarch.rpm                                                                                                              2.9 MB/s | 118 kB     00:00
(43/45): perl-threads-2.21-2.el8.x86_64.rpm                                                                                                                1.8 MB/s |  61 kB     00:00
(44/45): perl-threads-shared-1.58-2.el8.x86_64.rpm                                                                                                         901 kB/s |  48 kB     00:00
(45/45): perl-interpreter-5.26.3-420.el8.x86_64.rpm                                                                                                         14 MB/s | 6.3 MB     00:00
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                                                                       14 MB/s |  18 MB     00:01
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                                                                   1/1
  Installing       : mariadb-connector-c-config-3.1.11-2.el8_3.noarch                                                                                                                 1/45
  Installing       : mariadb-common-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                    2/45
  Installing       : perl-Digest-1.17-395.el8.noarch                                                                                                                                  3/45
  Installing       : perl-Digest-MD5-2.55-396.el8.x86_64                                                                                                                              4/45
  Installing       : perl-Data-Dumper-2.167-399.el8.x86_64                                                                                                                            5/45
  Installing       : perl-libnet-3.11-3.el8.noarch                                                                                                                                    6/45
  Installing       : perl-Net-SSLeay-1.88-1.module_el8.3.0+410+ff426aa3.x86_64                                                                                                        7/45
  Installing       : perl-URI-1.73-3.el8.noarch                                                                                                                                       8/45
  Installing       : perl-Pod-Escapes-1:1.07-395.el8.noarch                                                                                                                           9/45
  Installing       : perl-Mozilla-CA-20160104-7.module_el8.3.0+416+dee7bcef.noarch                                                                                                   10/45
  Installing       : perl-IO-Socket-IP-0.39-5.el8.noarch                                                                                                                             11/45
  Installing       : perl-Time-Local-1:1.280-1.el8.noarch                                                                                                                            12/45
  Installing       : perl-IO-Socket-SSL-2.066-4.module_el8.3.0+410+ff426aa3.noarch                                                                                                   13/45
  Installing       : perl-Term-ANSIColor-4.06-396.el8.noarch                                                                                                                         14/45
  Installing       : perl-Term-Cap-1.17-395.el8.noarch                                                                                                                               15/45
  Installing       : perl-File-Temp-0.230.600-1.el8.noarch                                                                                                                           16/45
  Installing       : perl-Pod-Simple-1:3.35-395.el8.noarch                                                                                                                           17/45
  Installing       : perl-HTTP-Tiny-0.074-1.el8.noarch                                                                                                                               18/45
  Installing       : perl-podlators-4.11-1.el8.noarch                                                                                                                                19/45
  Installing       : perl-Pod-Perldoc-3.28-396.el8.noarch                                                                                                                            20/45
  Installing       : perl-Text-ParseWords-3.30-395.el8.noarch                                                                                                                        21/45
  Installing       : perl-Pod-Usage-4:1.69-395.el8.noarch                                                                                                                            22/45
  Installing       : perl-MIME-Base64-3.15-396.el8.x86_64                                                                                                                            23/45
  Installing       : perl-Storable-1:3.11-3.el8.x86_64                                                                                                                               24/45
  Installing       : perl-Getopt-Long-1:2.50-4.el8.noarch                                                                                                                            25/45
  Installing       : perl-Errno-1.28-420.el8.x86_64                                                                                                                                  26/45
  Installing       : perl-Socket-4:2.027-3.el8.x86_64                                                                                                                                27/45
  Installing       : perl-Encode-4:2.97-3.el8.x86_64                                                                                                                                 28/45
  Installing       : perl-Carp-1.42-396.el8.noarch                                                                                                                                   29/45
  Installing       : perl-Exporter-5.72-396.el8.noarch                                                                                                                               30/45
  Installing       : perl-libs-4:5.26.3-420.el8.x86_64                                                                                                                               31/45
  Installing       : perl-Scalar-List-Utils-3:1.49-2.el8.x86_64                                                                                                                      32/45
  Installing       : perl-parent-1:0.237-1.el8.noarch                                                                                                                                33/45
  Installing       : perl-macros-4:5.26.3-420.el8.x86_64                                                                                                                             34/45
  Installing       : perl-Text-Tabs+Wrap-2013.0523-395.el8.noarch                                                                                                                    35/45
  Installing       : perl-Unicode-Normalize-1.25-396.el8.x86_64                                                                                                                      36/45
  Installing       : perl-File-Path-2.15-2.el8.noarch                                                                                                                                37/45
  Installing       : perl-IO-1.38-420.el8.x86_64                                                                                                                                     38/45
  Installing       : perl-PathTools-3.74-1.el8.x86_64                                                                                                                                39/45
  Installing       : perl-constant-1.33-396.el8.noarch                                                                                                                               40/45
  Installing       : perl-threads-1:2.21-2.el8.x86_64                                                                                                                                41/45
  Installing       : perl-threads-shared-1.58-2.el8.x86_64                                                                                                                           42/45
  Installing       : perl-interpreter-4:5.26.3-420.el8.x86_64                                                                                                                        43/45
  Installing       : mariadb-connector-c-3.1.11-2.el8_3.x86_64                                                                                                                       44/45
  Installing       : mariadb-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                          45/45
  Running scriptlet: mariadb-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                          45/45
  Verifying        : mariadb-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                           1/45
  Verifying        : mariadb-common-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                    2/45
  Verifying        : mariadb-connector-c-3.1.11-2.el8_3.x86_64                                                                                                                        3/45
  Verifying        : mariadb-connector-c-config-3.1.11-2.el8_3.noarch                                                                                                                 4/45
  Verifying        : perl-Digest-1.17-395.el8.noarch                                                                                                                                  5/45
  Verifying        : perl-Digest-MD5-2.55-396.el8.x86_64                                                                                                                              6/45
  Verifying        : perl-IO-Socket-IP-0.39-5.el8.noarch                                                                                                                              7/45
  Verifying        : perl-IO-Socket-SSL-2.066-4.module_el8.3.0+410+ff426aa3.noarch                                                                                                    8/45
  Verifying        : perl-Mozilla-CA-20160104-7.module_el8.3.0+416+dee7bcef.noarch                                                                                                    9/45
  Verifying        : perl-Net-SSLeay-1.88-1.module_el8.3.0+410+ff426aa3.x86_64                                                                                                       10/45
  Verifying        : perl-URI-1.73-3.el8.noarch                                                                                                                                      11/45
  Verifying        : perl-libnet-3.11-3.el8.noarch                                                                                                                                   12/45
  Verifying        : perl-Carp-1.42-396.el8.noarch                                                                                                                                   13/45
  Verifying        : perl-Data-Dumper-2.167-399.el8.x86_64                                                                                                                           14/45
  Verifying        : perl-Encode-4:2.97-3.el8.x86_64                                                                                                                                 15/45
  Verifying        : perl-Errno-1.28-420.el8.x86_64                                                                                                                                  16/45
  Verifying        : perl-Exporter-5.72-396.el8.noarch                                                                                                                               17/45
  Verifying        : perl-File-Path-2.15-2.el8.noarch                                                                                                                                18/45
  Verifying        : perl-File-Temp-0.230.600-1.el8.noarch                                                                                                                           19/45
  Verifying        : perl-Getopt-Long-1:2.50-4.el8.noarch                                                                                                                            20/45
  Verifying        : perl-HTTP-Tiny-0.074-1.el8.noarch                                                                                                                               21/45
  Verifying        : perl-IO-1.38-420.el8.x86_64                                                                                                                                     22/45
  Verifying        : perl-MIME-Base64-3.15-396.el8.x86_64                                                                                                                            23/45
  Verifying        : perl-PathTools-3.74-1.el8.x86_64                                                                                                                                24/45
  Verifying        : perl-Pod-Escapes-1:1.07-395.el8.noarch                                                                                                                          25/45
  Verifying        : perl-Pod-Perldoc-3.28-396.el8.noarch                                                                                                                            26/45
  Verifying        : perl-Pod-Simple-1:3.35-395.el8.noarch                                                                                                                           27/45
  Verifying        : perl-Pod-Usage-4:1.69-395.el8.noarch                                                                                                                            28/45
  Verifying        : perl-Scalar-List-Utils-3:1.49-2.el8.x86_64                                                                                                                      29/45
  Verifying        : perl-Socket-4:2.027-3.el8.x86_64                                                                                                                                30/45
  Verifying        : perl-Storable-1:3.11-3.el8.x86_64                                                                                                                               31/45
  Verifying        : perl-Term-ANSIColor-4.06-396.el8.noarch                                                                                                                         32/45
  Verifying        : perl-Term-Cap-1.17-395.el8.noarch                                                                                                                               33/45
  Verifying        : perl-Text-ParseWords-3.30-395.el8.noarch                                                                                                                        34/45
  Verifying        : perl-Text-Tabs+Wrap-2013.0523-395.el8.noarch                                                                                                                    35/45
  Verifying        : perl-Time-Local-1:1.280-1.el8.noarch                                                                                                                            36/45
  Verifying        : perl-Unicode-Normalize-1.25-396.el8.x86_64                                                                                                                      37/45
  Verifying        : perl-constant-1.33-396.el8.noarch                                                                                                                               38/45
  Verifying        : perl-interpreter-4:5.26.3-420.el8.x86_64                                                                                                                        39/45
  Verifying        : perl-libs-4:5.26.3-420.el8.x86_64                                                                                                                               40/45
  Verifying        : perl-macros-4:5.26.3-420.el8.x86_64                                                                                                                             41/45
  Verifying        : perl-parent-1:0.237-1.el8.noarch                                                                                                                                42/45
  Verifying        : perl-podlators-4.11-1.el8.noarch                                                                                                                                43/45
  Verifying        : perl-threads-1:2.21-2.el8.x86_64                                                                                                                                44/45
  Verifying        : perl-threads-shared-1.58-2.el8.x86_64                                                                                                                           45/45

Installed:
  mariadb-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64   mariadb-common-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64   mariadb-connector-c-3.1.11-2.el8_3.x86_64
  mariadb-connector-c-config-3.1.11-2.el8_3.noarch         perl-Carp-1.42-396.el8.noarch                                   perl-Data-Dumper-2.167-399.el8.x86_64
  perl-Digest-1.17-395.el8.noarch                          perl-Digest-MD5-2.55-396.el8.x86_64                             perl-Encode-4:2.97-3.el8.x86_64
  perl-Errno-1.28-420.el8.x86_64                           perl-Exporter-5.72-396.el8.noarch                               perl-File-Path-2.15-2.el8.noarch
  perl-File-Temp-0.230.600-1.el8.noarch                    perl-Getopt-Long-1:2.50-4.el8.noarch                            perl-HTTP-Tiny-0.074-1.el8.noarch
  perl-IO-1.38-420.el8.x86_64                              perl-IO-Socket-IP-0.39-5.el8.noarch                             perl-IO-Socket-SSL-2.066-4.module_el8.3.0+410+ff426aa3.noarch
  perl-MIME-Base64-3.15-396.el8.x86_64                     perl-Mozilla-CA-20160104-7.module_el8.3.0+416+dee7bcef.noarch   perl-Net-SSLeay-1.88-1.module_el8.3.0+410+ff426aa3.x86_64
  perl-PathTools-3.74-1.el8.x86_64                         perl-Pod-Escapes-1:1.07-395.el8.noarch                          perl-Pod-Perldoc-3.28-396.el8.noarch
  perl-Pod-Simple-1:3.35-395.el8.noarch                    perl-Pod-Usage-4:1.69-395.el8.noarch                            perl-Scalar-List-Utils-3:1.49-2.el8.x86_64
  perl-Socket-4:2.027-3.el8.x86_64                         perl-Storable-1:3.11-3.el8.x86_64                               perl-Term-ANSIColor-4.06-396.el8.noarch
  perl-Term-Cap-1.17-395.el8.noarch                        perl-Text-ParseWords-3.30-395.el8.noarch                        perl-Text-Tabs+Wrap-2013.0523-395.el8.noarch
  perl-Time-Local-1:1.280-1.el8.noarch                     perl-URI-1.73-3.el8.noarch                                      perl-Unicode-Normalize-1.25-396.el8.x86_64
  perl-constant-1.33-396.el8.noarch                        perl-interpreter-4:5.26.3-420.el8.x86_64                        perl-libnet-3.11-3.el8.noarch
  perl-libs-4:5.26.3-420.el8.x86_64                        perl-macros-4:5.26.3-420.el8.x86_64                             perl-parent-1:0.237-1.el8.noarch
  perl-podlators-4.11-1.el8.noarch                         perl-threads-1:2.21-2.el8.x86_64                                perl-threads-shared-1.58-2.el8.x86_64

Complete!
```

```
[bob@centos-host ~]$ sudo yum install -y mariadb-server
Last metadata expiration check: 0:19:35 ago on Sun 05 Jun 2022 11:53:02 PM UTC.
Module yaml error: Unexpected key in data: static_context [line 9 col 3]
Module yaml error: Unexpected key in data: static_context [line 9 col 3]
Dependencies resolved.
===========================================================================================================================================================================================
 Package                                        Architecture                    Version                                                           Repository                          Size
===========================================================================================================================================================================================
Installing:
 mariadb-server                                 x86_64                          3:10.3.28-1.module_el8.3.0+757+d382997d                           appstream                           16 M
Installing dependencies:
 libaio                                         x86_64                          0.3.112-1.el8                                                     baseos                              33 k
 mariadb-errmsg                                 x86_64                          3:10.3.28-1.module_el8.3.0+757+d382997d                           appstream                          234 k
 perl-DBD-MySQL                                 x86_64                          4.046-3.module_el8.3.0+419+c2dec72b                               appstream                          156 k
 perl-DBI                                       x86_64                          1.641-3.module_el8.3.0+413+9be2aeb5                               appstream                          740 k
 perl-Math-BigInt                               noarch                          1:1.9998.11-7.el8                                                 baseos                             196 k
 perl-Math-Complex                              noarch                          1.59-420.el8                                                      baseos                             108 k
 psmisc                                         x86_64                          23.1-5.el8                                                        baseos                             151 k
Installing weak dependencies:
 mariadb-backup                                 x86_64                          3:10.3.28-1.module_el8.3.0+757+d382997d                           appstream                          6.1 M
 mariadb-gssapi-server                          x86_64                          3:10.3.28-1.module_el8.3.0+757+d382997d                           appstream                           51 k
 mariadb-server-utils                           x86_64                          3:10.3.28-1.module_el8.3.0+757+d382997d                           appstream                          1.1 M
Enabling module streams:
 perl-DBD-MySQL                                                                 4.046
 perl-DBI                                                                       1.641

Transaction Summary
===========================================================================================================================================================================================
Install  11 Packages

Total download size: 25 M
Installed size: 119 M
Downloading Packages:
(1/11): mariadb-gssapi-server-10.3.28-1.module_el8.3.0+757+d382997d.x86_64.rpm                                                                             321 kB/s |  51 kB     00:00
(2/11): mariadb-errmsg-10.3.28-1.module_el8.3.0+757+d382997d.x86_64.rpm                                                                                    1.1 MB/s | 234 kB     00:00
(3/11): mariadb-server-utils-10.3.28-1.module_el8.3.0+757+d382997d.x86_64.rpm                                                                               12 MB/s | 1.1 MB     00:00
(4/11): perl-DBD-MySQL-4.046-3.module_el8.3.0+419+c2dec72b.x86_64.rpm                                                                                      3.9 MB/s | 156 kB     00:00
(5/11): mariadb-backup-10.3.28-1.module_el8.3.0+757+d382997d.x86_64.rpm                                                                                     17 MB/s | 6.1 MB     00:00
(6/11): libaio-0.3.112-1.el8.x86_64.rpm                                                                                                                    751 kB/s |  33 kB     00:00
(7/11): perl-Math-BigInt-1.9998.11-7.el8.noarch.rpm                                                                                                        5.2 MB/s | 196 kB     00:00
(8/11): perl-DBI-1.641-3.module_el8.3.0+413+9be2aeb5.x86_64.rpm                                                                                            5.1 MB/s | 740 kB     00:00
(9/11): perl-Math-Complex-1.59-420.el8.noarch.rpm                                                                                                          2.0 MB/s | 108 kB     00:00
(10/11): mariadb-server-10.3.28-1.module_el8.3.0+757+d382997d.x86_64.rpm                                                                                    41 MB/s |  16 MB     00:00
(11/11): psmisc-23.1-5.el8.x86_64.rpm                                                                                                                      2.1 MB/s | 151 kB     00:00
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                                                                       45 MB/s |  25 MB     00:00
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                                                                   1/1
  Installing       : libaio-0.3.112-1.el8.x86_64                                                                                                                                      1/11
  Installing       : psmisc-23.1-5.el8.x86_64                                                                                                                                         2/11
  Installing       : perl-Math-Complex-1.59-420.el8.noarch                                                                                                                            3/11
  Installing       : perl-Math-BigInt-1:1.9998.11-7.el8.noarch                                                                                                                        4/11
  Installing       : perl-DBI-1.641-3.module_el8.3.0+413+9be2aeb5.x86_64                                                                                                              5/11
  Installing       : perl-DBD-MySQL-4.046-3.module_el8.3.0+419+c2dec72b.x86_64                                                                                                        6/11
  Installing       : mariadb-errmsg-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                    7/11
  Installing       : mariadb-gssapi-server-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                             8/11
  Installing       : mariadb-server-utils-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                              9/11
  Running scriptlet: mariadb-server-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                   10/11
  Installing       : mariadb-server-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                   10/11
  Running scriptlet: mariadb-server-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                   10/11
  Installing       : mariadb-backup-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                   11/11
  Running scriptlet: mariadb-backup-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                   11/11
  Verifying        : mariadb-backup-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                    1/11
  Verifying        : mariadb-errmsg-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                    2/11
  Verifying        : mariadb-gssapi-server-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                             3/11
  Verifying        : mariadb-server-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                                    4/11
  Verifying        : mariadb-server-utils-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                                                                                              5/11
  Verifying        : perl-DBD-MySQL-4.046-3.module_el8.3.0+419+c2dec72b.x86_64                                                                                                        6/11
  Verifying        : perl-DBI-1.641-3.module_el8.3.0+413+9be2aeb5.x86_64                                                                                                              7/11
  Verifying        : libaio-0.3.112-1.el8.x86_64                                                                                                                                      8/11
  Verifying        : perl-Math-BigInt-1:1.9998.11-7.el8.noarch                                                                                                                        9/11
  Verifying        : perl-Math-Complex-1.59-420.el8.noarch                                                                                                                           10/11
  Verifying        : psmisc-23.1-5.el8.x86_64                                                                                                                                        11/11

Installed:
  libaio-0.3.112-1.el8.x86_64                                                              mariadb-backup-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64
  mariadb-errmsg-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                            mariadb-gssapi-server-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64
  mariadb-server-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64                            mariadb-server-utils-3:10.3.28-1.module_el8.3.0+757+d382997d.x86_64
  perl-DBD-MySQL-4.046-3.module_el8.3.0+419+c2dec72b.x86_64                                perl-DBI-1.641-3.module_el8.3.0+413+9be2aeb5.x86_64
  perl-Math-BigInt-1:1.9998.11-7.el8.noarch                                                perl-Math-Complex-1.59-420.el8.noarch
  psmisc-23.1-5.el8.x86_64

Complete!
```

```
[bob@centos-host ~]$ sudo systemctl enable mariadb
Created symlink /etc/systemd/system/mysql.service → /usr/lib/systemd/system/mariadb.service.
Created symlink /etc/systemd/system/mysqld.service → /usr/lib/systemd/system/mariadb.service.
Created symlink /etc/systemd/system/multi-user.target.wants/mariadb.service → /usr/lib/systemd/system/mariadb.service.
```

```
[bob@centos-host ~]$ sudo systemctl start mariadb
```

- [x] Set a password for mysql root user to "S3cure#321"

```
[bob@centos-host ~]$ sudo mysql -u root <<EOF
ALTER USER 'root'@'localhost' IDENTIFIED BY 'S3cure#321';
flush privileges;
EOF
```

```
[bob@centos-host ~]$ sudo mysql -u root -p <<EOF
EOF

Enter password:
```

- [x] The "root" account is currently locked on "centos-host", please unlock it.

```
[bob@centos-host ~]$ sudo passwd -u root -u
Unlocking password for user root.
passwd: Success
```

- [x] Make user "root" a member of "wheel" group

```
[bob@centos-host ~]$ sudo usermod -a -G wheel root
```

```
[bob@centos-host ~]$ id root
uid=0(root) gid=0(root) groups=0(root),10(wheel)
```

- [x] Edit the PAM configuration file for the "su" utility so that this utility only accepts the requests from the users that are part of the "wheel" group and the requests from the users should be accepted immediately, without asking for any password.

```
[bob@centos-host ~]$ grep wheel /etc/pam.d/su
# Uncomment the following line to implicitly trust users in the "wheel" group.
#auth           sufficient      pam_wheel.so trust use_uid
# Uncomment the following line to require a user to be in the "wheel" group.
#auth           required        pam_wheel.so use_uid
```

```
[bob@centos-host ~]$ sudoedit /etc/pam.d/su
```

```
[bob@centos-host ~]$ grep wheel /etc/pam.d/su
# Uncomment the following line to implicitly trust users in the "wheel" group.
auth            sufficient      pam_wheel.so trust use_uid
# Uncomment the following line to require a user to be in the "wheel" group.
auth            required        pam_wheel.so use_uid
```

- [x] Add an extra IP to "eth1" interface on this system: 10.0.0.50/24

```
[bob@centos-host ~]$ ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 52:54:00:17:26:ad brd ff:ff:ff:ff:ff:ff
    inet 192.168.121.35/24 brd 192.168.121.255 scope global dynamic noprefixroute eth0
       valid_lft 2171sec preferred_lft 2171sec
    inet6 fe80::5054:ff:fe17:26ad/64 scope link
       valid_lft forever preferred_lft forever
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 52:54:00:21:24:bb brd ff:ff:ff:ff:ff:ff
    inet 172.28.128.195/24 brd 172.28.128.255 scope global dynamic noprefixroute eth1
       valid_lft 2196sec preferred_lft 2196sec
    inet6 fe80::5054:ff:fe21:24bb/64 scope link
       valid_lft forever preferred_lft forever
4: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default
    link/ether 02:42:7c:c1:a0:bd brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
```

```
[bob@centos-host ~]$ sudo grep 'NM_CONTROLLED' /etc/sysconfig/network-scripts/ifcfg-eth1
NM_CONTROLLED=yes
```

```
[bob@centos-host ~]$ sudo nmtui
```

```
[bob@centos-host ~]$ ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 52:54:00:17:26:ad brd ff:ff:ff:ff:ff:ff
    inet 192.168.121.35/24 brd 192.168.121.255 scope global dynamic noprefixroute eth0
       valid_lft 3562sec preferred_lft 3562sec
    inet6 fe80::5054:ff:fe17:26ad/64 scope link
       valid_lft forever preferred_lft forever
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 52:54:00:21:24:bb brd ff:ff:ff:ff:ff:ff
    inet 10.0.0.50/24 brd 10.0.0.255 scope global noprefixroute eth1
       valid_lft forever preferred_lft forever
    inet 172.28.128.195/24 brd 172.28.128.255 scope global dynamic noprefixroute eth1
       valid_lft 3582sec preferred_lft 3582sec
    inet6 fe80::5054:ff:fe21:24bb/64 scope link
       valid_lft forever preferred_lft forever
4: docker0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default
    link/ether 02:42:7c:c1:a0:bd brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
```

- [x] Add a local DNS entry for the database hostname "mydb.kodekloud.com" so that it can resolve to "10.0.0.50" IP address.

```
[bob@centos-host ~]$ sudo bash -c 'cat >> /etc/hosts' <<EOF
> 10.0.0.50    mydb.kodekloud.com
> EOF
```

```
[bob@centos-host ~]$ ping -c 3 mydb.kodekloud.com
PING mydb.kodekloud.com (10.0.0.50) 56(84) bytes of data.
64 bytes from mydb.kodekloud.com (10.0.0.50): icmp_seq=1 ttl=64 time=0.137 ms
64 bytes from mydb.kodekloud.com (10.0.0.50): icmp_seq=2 ttl=64 time=0.048 ms
64 bytes from mydb.kodekloud.com (10.0.0.50): icmp_seq=3 ttl=64 time=0.045 ms

--- mydb.kodekloud.com ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 82ms
rtt min/avg/max/mdev = 0.045/0.076/0.137/0.043 ms
```

- [x] Pull "nginx" docker image.

```
[bob@centos-host ~]$ sudo docker pull nginx
Using default tag: latest
latest: Pulling from library/nginx
42c077c10790: Pull complete
62c70f376f6a: Pull complete
915cc9bd79c2: Pull complete
75a963e94de0: Pull complete
7b1fab684d70: Pull complete
db24d06d5af4: Pull complete
Digest: sha256:2bcabc23b45489fb0885d69a06ba1d648aeda973fae7bb981bafbb884165e514
Status: Downloaded newer image for nginx:latest
docker.io/library/nginx:latest
```

```
[bob@centos-host ~]$ sudo docker images
REPOSITORY   TAG       IMAGE ID       CREATED      SIZE
nginx        latest    0e901e68141f   8 days ago   142MB
```

- [x] Create and run a new Docker container based on the "nginx" image. The container should be named as "myapp" and the port "80" on the host should be mapped to the port "80" on the container.

```
[bob@centos-host ~]$ sudo docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```

```
[bob@centos-host ~]$ sudo docker run --name myapp -d -p 80:80 nginx:latest
b876970bedc1ecc4e95b74c58cc400dab50acff13d5d95d2818b77f864f684f4
```

```
[bob@centos-host ~]$ sudo docker ps
CONTAINER ID   IMAGE          COMMAND                  CREATED         STATUS        PORTS                               NAMES
b876970bedc1   nginx:latest   "/docker-entrypoint.…"   3 seconds ago   Up 1 second   0.0.0.0:80->80/tcp, :::80->80/tcp   myapp
```

- [x] Create a bash script called "container-start.sh" under "/home/bob/" which should be able to "start" the "myapp" container. It should also display a message "myapp container started!"

```
[bob@centos-host ~]$ cat > container-start.sh <<EOF
> #!/bin/bash
>
> sudo docker start myapp
>
> printf "%s\n" "myapp container started!"
> EOF
```

```
[bob@centos-host ~]$ chmod -v a+x ./container-start.sh
mode of './container-start.sh' changed from 0664 (rw-rw-r--) to 0775 (rwxrwxr-x)
```

- [x] Create a bash script called "container-stop.sh" under "/home/bob/" which should be able to stop the "myapp" container. It should also display a message "myapp container stopped!"

```
[bob@centos-host ~]$ cat > container-stop.sh <<EOF
#!/bin/bash

sudo docker stop myapp

printf "%s\n" "myapp container stopped!"
EOF
```

```
[bob@centos-host ~]$ chmod -v a+x ./container-stop.sh
mode of './container-stop.sh' changed from 0664 (rw-rw-r--) to 0775 (rwxrwxr-x)
```

- [x] Add a cron job for the "root" user which should run "container-stop.sh" script at "12am" everyday.

```
[bob@centos-host ~]$ sudo crontab -u root -e
no crontab for root - using an empty one
crontab: installing new crontab
```

```
[bob@centos-host ~]$ sudo crontab -l -u root
# Example of job definition:
# .---------------- minute (0 - 59)
# |  .------------- hour (0 - 23)
# |  |  .---------- day of month (1 - 31)
# |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...
# |  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat
# |  |  |  |  |
# *  *  *  *  * user-name  command to be executed
0  0  *  *  * root  /home/bob/container-stop.sh
```

- [x] Add a cron job for the "root" user which should run "container-start.sh" script at "8am" everyday.

```
[bob@centos-host ~]$ sudo crontab -u root -e
crontab: installing new crontab
```

```
[bob@centos-host ~]$ sudo crontab -l -u root
# Example of job definition:
# .---------------- minute (0 - 59)
# |  .------------- hour (0 - 23)
# |  |  .---------- day of month (1 - 31)
# |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...
# |  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat
# |  |  |  |  |
# *  *  *  *  * user-name  command to be executed
0  0  *  *  * root  /home/bob/container-stop.sh
0  8  *  *  * root  /home/bob/container-start.sh
```

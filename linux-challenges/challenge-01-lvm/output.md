# KodeKloud - Linux Challenge - Challenge 1 Output

- [x] Inspect OS

```
[bob@centos-host ~]$ uname -r
4.18.0-240.1.1.el8_3.x86_64
```

```
[bob@centos-host ~]$ cat /etc/os-release
NAME="CentOS Linux"
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

- [x] Install the correct packages that will allow the use of "lvm" on the centos machine.

```
[bob@centos-host ~]$ sudo yum install lvm2 tmux
Last metadata expiration check: 0:25:19 ago on Sun 05 Jun 2022 03:48:21 PM UTC.
Module yaml error: Unexpected key in data: static_context [line 9 col 3]
Module yaml error: Unexpected key in data: static_context [line 9 col 3]
Dependencies resolved.
===========================================================================================================================================================================================
 Package                                                    Architecture                        Version                                          Repository                           Size
===========================================================================================================================================================================================
Installing:
 lvm2                                                       x86_64                              8:2.03.12-10.el8                                 baseos                              1.6 M
 tmux                                                       x86_64                              2.7-1.el8                                        baseos                              317 k
Upgrading:
 device-mapper                                              x86_64                              8:1.02.177-10.el8                                baseos                              377 k
 device-mapper-libs                                         x86_64                              8:1.02.177-10.el8                                baseos                              409 k
Installing dependencies:
 device-mapper-event                                        x86_64                              8:1.02.177-10.el8                                baseos                              271 k
 device-mapper-event-libs                                   x86_64                              8:1.02.177-10.el8                                baseos                              270 k
 device-mapper-persistent-data                              x86_64                              0.9.0-4.el8                                      baseos                              925 k
 libaio                                                     x86_64                              0.3.112-1.el8                                    baseos                               33 k
 lvm2-libs                                                  x86_64                              8:2.03.12-10.el8                                 baseos                              1.2 M

Transaction Summary
===========================================================================================================================================================================================
Install  7 Packages
Upgrade  2 Packages

Total download size: 5.3 M
Is this ok [y/N]: y
Downloading Packages:
(1/9): device-mapper-event-libs-1.02.177-10.el8.x86_64.rpm                                                                                                 1.4 MB/s | 270 kB     00:00
(2/9): device-mapper-persistent-data-0.9.0-4.el8.x86_64.rpm                                                                                                4.4 MB/s | 925 kB     00:00
(3/9): device-mapper-event-1.02.177-10.el8.x86_64.rpm                                                                                                      1.2 MB/s | 271 kB     00:00
(4/9): libaio-0.3.112-1.el8.x86_64.rpm                                                                                                                     785 kB/s |  33 kB     00:00
(5/9): tmux-2.7-1.el8.x86_64.rpm                                                                                                                           6.1 MB/s | 317 kB     00:00
(6/9): lvm2-libs-2.03.12-10.el8.x86_64.rpm                                                                                                                  11 MB/s | 1.2 MB     00:00
(7/9): lvm2-2.03.12-10.el8.x86_64.rpm                                                                                                                       12 MB/s | 1.6 MB     00:00
(8/9): device-mapper-1.02.177-10.el8.x86_64.rpm                                                                                                            4.2 MB/s | 377 kB     00:00
(9/9): device-mapper-libs-1.02.177-10.el8.x86_64.rpm                                                                                                       3.6 MB/s | 409 kB     00:00
-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                                                                       12 MB/s | 5.3 MB     00:00
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                                                                   1/1
  Installing       : libaio-0.3.112-1.el8.x86_64                                                                                                                                      1/11
  Upgrading        : device-mapper-libs-8:1.02.177-10.el8.x86_64                                                                                                                      2/11
  Upgrading        : device-mapper-8:1.02.177-10.el8.x86_64                                                                                                                           3/11
  Installing       : device-mapper-event-libs-8:1.02.177-10.el8.x86_64                                                                                                                4/11
  Installing       : device-mapper-event-8:1.02.177-10.el8.x86_64                                                                                                                     5/11
  Running scriptlet: device-mapper-event-8:1.02.177-10.el8.x86_64                                                                                                                     5/11
  Installing       : lvm2-libs-8:2.03.12-10.el8.x86_64                                                                                                                                6/11
  Installing       : device-mapper-persistent-data-0.9.0-4.el8.x86_64                                                                                                                 7/11
  Installing       : lvm2-8:2.03.12-10.el8.x86_64                                                                                                                                     8/11
  Running scriptlet: lvm2-8:2.03.12-10.el8.x86_64                                                                                                                                     8/11
  Installing       : tmux-2.7-1.el8.x86_64                                                                                                                                            9/11
  Running scriptlet: tmux-2.7-1.el8.x86_64                                                                                                                                            9/11
  Cleanup          : device-mapper-8:1.02.171-5.el8.x86_64                                                                                                                           10/11
  Cleanup          : device-mapper-libs-8:1.02.171-5.el8.x86_64                                                                                                                      11/11
  Running scriptlet: device-mapper-libs-8:1.02.171-5.el8.x86_64                                                                                                                      11/11
  Verifying        : device-mapper-event-8:1.02.177-10.el8.x86_64                                                                                                                     1/11
  Verifying        : device-mapper-event-libs-8:1.02.177-10.el8.x86_64                                                                                                                2/11
  Verifying        : device-mapper-persistent-data-0.9.0-4.el8.x86_64                                                                                                                 3/11
  Verifying        : libaio-0.3.112-1.el8.x86_64                                                                                                                                      4/11
  Verifying        : lvm2-8:2.03.12-10.el8.x86_64                                                                                                                                     5/11
  Verifying        : lvm2-libs-8:2.03.12-10.el8.x86_64                                                                                                                                6/11
  Verifying        : tmux-2.7-1.el8.x86_64                                                                                                                                            7/11
  Verifying        : device-mapper-8:1.02.177-10.el8.x86_64                                                                                                                           8/11
  Verifying        : device-mapper-8:1.02.171-5.el8.x86_64                                                                                                                            9/11
  Verifying        : device-mapper-libs-8:1.02.177-10.el8.x86_64                                                                                                                     10/11
  Verifying        : device-mapper-libs-8:1.02.171-5.el8.x86_64                                                                                                                      11/11

Upgraded:
  device-mapper-8:1.02.177-10.el8.x86_64                                                    device-mapper-libs-8:1.02.177-10.el8.x86_64

Installed:
  device-mapper-event-8:1.02.177-10.el8.x86_64    device-mapper-event-libs-8:1.02.177-10.el8.x86_64    device-mapper-persistent-data-0.9.0-4.el8.x86_64    libaio-0.3.112-1.el8.x86_64
  lvm2-8:2.03.12-10.el8.x86_64                    lvm2-libs-8:2.03.12-10.el8.x86_64                    tmux-2.7-1.el8.x86_64

Complete!
```

- [x] Create a group called "dba_users" and add the user called 'bob' to this group

```
[bob@centos-host ~]$ sudo groupadd dba_users
```

```
[bob@centos-host ~]$ sudo usermod -a -G dba_users bob
```

```
[bob@centos-host ~]$ sudo grep dba_users /etc/group
dba_users:x:1003:bob
```

- [x] Create a Physical Volume for "/dev/vdb"

```
[bob@centos-host ~]$ ls /dev/vd*
/dev/vda  /dev/vda1  /dev/vdb  /dev/vdc  /dev/vdd  /dev/vde
```

```
[bob@centos-host ~]$ sudo parted /dev/vdb p
Error: /dev/vdb: unrecognised disk label
Model: Virtio Block Device (virtblk)
Disk /dev/vdb: 1074MB
Sector size (logical/physical): 512B/512B
Partition Table: unknown
Disk Flags:
```

```
[bob@centos-host ~]$ sudo parted /dev/vdb -s -- mklabel gpt p mkpart lvm 1m -1 set 1 lvm on p
Model: Virtio Block Device (virtblk)
Disk /dev/vdb: 1074MB
Sector size (logical/physical): 512B/512B
Partition Table: gpt
Disk Flags:

Number  Start  End  Size  File system  Name  Flags

Model: Virtio Block Device (virtblk)
Disk /dev/vdb: 1074MB
Sector size (logical/physical): 512B/512B
Partition Table: gpt
Disk Flags:

Number  Start   End     Size    File system  Name  Flags
 1      1049kB  1073MB  1072MB               lvm   lvm
```

```
[bob@centos-host ~]$ sudo pvcreate /dev/vdb1
  Physical volume "/dev/vdb1" successfully created.
```

```
[bob@centos-host ~]$ sudo pvs
  PV         VG Fmt  Attr PSize    PFree
  /dev/vdb1     lvm2 ---  1022.00m 1022.00m
```

- [x] Create a Physical Volume for "/dev/vdc"

```
[bob@centos-host ~]$ sudo parted /dev/vdc p
Error: /dev/vdc: unrecognised disk label
Model: Virtio Block Device (virtblk)
Disk /dev/vdc: 1074MB
Sector size (logical/physical): 512B/512B
Partition Table: unknown
Disk Flags:
```

```
[bob@centos-host ~]$ sudo parted /dev/vdc -s -- mklabel gpt p mkpart lvm 1m -1 set 1 lvm on p
Model: Virtio Block Device (virtblk)
Disk /dev/vdc: 1074MB
Sector size (logical/physical): 512B/512B
Partition Table: gpt
Disk Flags:

Number  Start  End  Size  File system  Name  Flags

Model: Virtio Block Device (virtblk)
Disk /dev/vdc: 1074MB
Sector size (logical/physical): 512B/512B
Partition Table: gpt
Disk Flags:

Number  Start   End     Size    File system  Name  Flags
 1      1049kB  1073MB  1072MB               lvm   lvm
```

```
[bob@centos-host ~]$ sudo pvcreate /dev/vdc1
  Physical volume "/dev/vdc1" successfully created.
```

```
[bob@centos-host ~]$ sudo pvs
  PV         VG Fmt  Attr PSize    PFree
  /dev/vdb1     lvm2 ---  1022.00m 1022.00m
  /dev/vdc1     lvm2 ---  1022.00m 1022.00m
```

- [x] Create a volume group called "dba_storage" using the physical volumes "/dev/vdb" and "/dev/vdc"

```
[bob@centos-host ~]$ sudo vgcreate dba_storage /dev/vd{b,c}1
  Volume group "dba_storage" successfully created
```

```
[bob@centos-host ~]$ sudo vgs
  VG          #PV #LV #SN Attr   VSize VFree
  dba_storage   2   0   0 wz--n- 1.99g 1.99g
```

```
[bob@centos-host ~]$ sudo pvs
  PV         VG          Fmt  Attr PSize    PFree
  /dev/vdb1  dba_storage lvm2 a--  1020.00m 1020.00m
  /dev/vdc1  dba_storage lvm2 a--  1020.00m 1020.00m
```

- [x] Create an "lvm" called "volume_1" from the volume group called "dba_storage". Make use of the entire space available in the volume group.

```
[bob@centos-host ~]$ sudo lvcreate --name volume_1 --extents 100%VG dba_storage
  Logical volume "volume_1" created.
  ```

```
[bob@centos-host ~]$ sudo lvs
  LV       VG          Attr       LSize Pool Origin Data%  Meta%  Move Log Cpy%Sync Convert
  volume_1 dba_storage -wi-ao---- 1.99g
```

```
[bob@centos-host ~]$ sudo vgs
  VG          #PV #LV #SN Attr   VSize VFree
  dba_storage   2   1   0 wz--n- 1.99g    0
```

```
[bob@centos-host ~]$ sudo pvs
  PV         VG          Fmt  Attr PSize    PFree
  /dev/vdb1  dba_storage lvm2 a--  1020.00m    0
  /dev/vdc1  dba_storage lvm2 a--  1020.00m    0
```

- [x] Format the lvm volume "volume_1" as an "XFS" filesystem

```
[bob@centos-host ~]$ sudo mkfs.xfs -L dba_storage /dev/mapper/dba_storage-volume_1
meta-data=/dev/mapper/dba_storage-volume_1 isize=512    agcount=4, agsize=130560 blks
         =                       sectsz=512   attr=2, projid32bit=1
         =                       crc=1        finobt=1, sparse=1, rmapbt=0
         =                       reflink=1
data     =                       bsize=4096   blocks=522240, imaxpct=25
         =                       sunit=0      swidth=0 blks
naming   =version 2              bsize=4096   ascii-ci=0, ftype=1
log      =internal log           bsize=4096   blocks=2560, version=2
         =                       sectsz=512   sunit=0 blks, lazy-count=1
realtime =none                   extsz=4096   blocks=0, rtextents=0
```

- [x] Make sure that this mount point is persistent across reboots with the correct default options

Notes:
- This one worked once using UUID=, hasn't worked with LABEL= or with a device node.
- Not sure if the test prefers tabs vs spaces. The only time it worked was with tabs and edited by hand using vi. The rest of the file uses spaces.

```
[bob@centos-host ~]$ sudo bash -c 'printf "UUID=%s\t/mnt/dba_storage\txfs\tdefaults\t0 0\n" "$(lsblk --fs /dev/mapper/dba_storage-volume_1 | awk "/dba_storage/ { print \$4 }")" >> /etc/fstab'
```

```
[bob@centos-host ~]$ tail -v /etc/fstab==> /etc/fstab <==
#
# After editing this file, run 'systemctl daemon-reload' to update systemd
# units generated from this file.
#
UUID=a62c5b49-755e-41b0-9d36-de3d95e17232 /                       xfs     defaults        0 0
/swapfile none swap defaults 0 0
#VAGRANT-BEGIN
# The contents below are automatically generated by Vagrant. Do not modify.
#VAGRANT-END
UUID=4e35881b-4c0a-408c-9a3d-9959afc1b65c        /mnt/dba_storage        xfs        defaults        0 0
```

```
[bob@centos-host ~]$ sudo systemctl daemon-reload
```

- [x] Mount the filesystem at the path "/mnt/dba_storage"

```
[bob@centos-host ~]$ sudo mkdir /mnt/dba_storage
```

```
[bob@centos-host ~]$ sudo mount -av
/                        : ignored
none                     : ignored
mount: /mnt/dba_storage does not contain SELinux labels.
       You just mounted an file system that supports labels which does not
       contain labels, onto an SELinux box. It is likely that confined
       applications will generate AVC messages and not be allowed access to
       this file system.  For more details see restorecon(8) and mount(8).
/mnt/dba_storage         : successfully mounted
```

```
[bob@centos-host ~]$ sudo restorecon -R /mnt/dba_storage
```

```
[bob@centos-host ~]$ df -h /mnt/dba_storage
Filesystem                        Size  Used Avail Use% Mounted on
/dev/mapper/dba_storage-volume_1  2.0G   47M  2.0G   3% /mnt/dba_storage
```

- [x] Ensure that the mountpoint "/mnt/dba_storage" has the group ownership set to the "dba_users" group

```
[bob@centos-host ~]$ sudo chgrp -v dba_users /mnt/dba_storage/
changed group of '/mnt/dba_storage/' from root to dba_users
```

- [x] Ensure that the mount point "/mnt/dba_storage" has "read/write" and execute permissions for the owner and group and no permissions for anyone else.

```
[bob@centos-host ~]$ sudo chmod -v 0770 /mnt/dba_storage/
mode of '/mnt/dba_storage/' changed from 0755 (rwxr-xr-x) to 0770 (rwxrwx---)
```

```
[bob@centos-host ~]$ ls -dl /mnt/dba_storage/
drwxrwx---. 2 root dba_users 6 Jun  5 16:22 /mnt/dba_storage/
```

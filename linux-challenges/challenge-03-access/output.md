# KodeKloud - Linux Challenge - Challenge 3 Output

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

- [x] Create a group called "admins"

```
[bob@centos-host ~]$ sudo groupadd admins
```

```
[bob@centos-host ~]$ sudo tail -v -n 1 /etc/group
==> /etc/group <==
admins:x:1003:
```

- [x] Create a group called "devs"

```
[bob@centos-host ~]$ sudo groupadd devs
```

```
[bob@centos-host ~]$ sudo tail -v -n 2 /etc/group
==> /etc/group <==
admins:x:1003:
devs:x:1004:
```

- [x] Make sure "/data" directory is owned by user "bob".

```
[bob@centos-host ~]$ sudo chown -v bob /data
changed ownership of '/data' from root to bob
```

- [x] Make sure "/data" directory is owned by user "bob" and group "devs" and "user/group" owner has "full" permissions but "other" should not have any permissions.

```
[bob@centos-host ~]$ sudo chown -v bob:devs /data
changed ownership of '/data' from bob:root to bob:devs
```

```
[bob@centos-host ~]$ sudo chmod -v 0770 /data
mode of '/data' changed from 0755 (rwxr-xr-x) to 0770 (rwxrwx---)
```

- [x] Give some additional permissions to "admins" group on "/data" directory so that any user who is the member the "admins" group has "full permissions" on this directory.

```
[bob@centos-host ~]$ getfacl /data
getfacl: Removing leading '/' from absolute path names
# file: data
# owner: bob
# group: devs
user::rwx
group::rwx
other::---
```

```
[bob@centos-host ~]$ sudo setfacl -d -m g:admins:rwx /data
```

```
[bob@centos-host ~]$ getfacl /data
getfacl: Removing leading '/' from absolute path names
# file: data
# owner: bob
# group: devs
user::rwx
group::rwx
other::---
default:user::rwx
default:group::rwx
default:group:admins:rwx
default:mask::rwx
default:other::---
```

- [x] Create a user called "ray" , change his login shell to "/bin/sh" and set "D3vU3r321" password for this user.
- [x] Make user "ray" a member of "devs" group.

```
[bob@centos-host ~]$ sudo useradd -c "Ray" -s /bin/sh -G devs ray
```

```
[bob@centos-host ~]$ sudo passwd ray
Changing password for user ray.
New password:
Retype new password:
passwd: all authentication tokens updated successfully.
```

```
[bob@centos-host ~]$ sudo grep ray /etc/passwd /etc/group
/etc/passwd:ray:x:1003:1005:Ray:/home/ray:/bin/sh
/etc/group:devs:x:1004:ray
/etc/group:ray:x:1005:
```

- [x] Create a user called "lisa", change her login shell to "/bin/sh" and set "D3vUd3r123" password for this user.
- [x] Make user "lisa" a member of "devs" group.

```
[bob@centos-host ~]$ sudo useradd -c "Lisa" -s /bin/sh -G devs lisa
```

```
[bob@centos-host ~]$ sudo passwd lisa
Changing password for user lisa.
New password:
Retype new password:
passwd: all authentication tokens updated successfully.
```

```
[bob@centos-host ~]$ sudo grep lisa /etc/passwd /etc/group
/etc/passwd:lisa:x:1004:1006:Lisa:/home/lisa:/bin/sh
/etc/group:devs:x:1004:ray,lisa
/etc/group:lisa:x:1006:
```

- [x] Create a user called "natasha" , change her login shell to "/bin/zsh" and set "DwfawUd113" password for this user.
- [x] Make user "natasha" a member of "admins" group.

```
[bob@centos-host ~]$ sudo useradd -c "Natasha" -s /bin/zsh -G admins natasha
```

```
[bob@centos-host ~]$ sudo passwd natasha
Changing password for user natasha.
New password:
Retype new password:
passwd: all authentication tokens updated successfully.
```

```
[bob@centos-host ~]$ sudo grep natasha /etc/passwd /etc/group
/etc/passwd:natasha:x:1005:1007:Natasha:/home/natasha:/bin/zsh
/etc/group:admins:x:1003:natasha
/etc/group:natasha:x:1007:
```

- [x] Create a user called "david" , change his login shell to "/bin/zsh" and set "D3vUd3raaw" password for this user.
- [x] Make user "david" a member of "admins" group.

```
[bob@centos-host ~]$ sudo useradd -c "David" -s /bin/zsh -G admins david
```

```
[bob@centos-host ~]$ sudo passwd david
Changing password for user david.
New password:
Retype new password:
passwd: all authentication tokens updated successfully.
```

```
[bob@centos-host ~]$ sudo grep david /etc/passwd /etc/group
/etc/passwd:david:x:1006:1008:David:/home/david:/bin/zsh
/etc/group:admins:x:1003:natasha,david
/etc/group:david:x:1008:
```

- [x] Configure a "resource limit" for the "devs" group so that this group (members of the group) can not run more than "30 processes" in their session. This should be both a "hard limit" and a "soft limit", written in a single line.

```
[bob@centos-host ~]$ sudo bash -c 'cat >> /etc/limits.conf' <<EOF
> devs	hard	nproc	30
> devs	soft	nproc	30
> EOF
```

```
[bob@centos-host ~]$ grep nproc /etc/limits.conf
devs    hard    nproc   30
devs    soft    nproc   30
```

- [x] Edit the disk quota for the group called "devs". Limit the amount of storage space it can use (not inodes). Set a "soft" limit of "100MB" and a "hard" limit of "500MB" on "/data" partition.

Notes:
- As far as I can tell, quotas are setup correctly, but tests are failing to pass this step.

```
[bob@centos-host ~]$ mount | grep /data
/dev/vdb1 on /data type xfs (rw,relatime,seclabel,attr2,inode64,logbufs=8,logbsize=32k,usrquota,prjquota,grpquota)
```

```
[bob@centos-host ~]$ sudo xfs_quota -x -c 'limit -g bsoft=100g bhard=500g devs' /data
```

```
[bob@centos-host ~]$ sudo xfs_quota -x -c 'report -h -g' /data
Group quota on /data (/dev/vdb1)
                        Blocks
Group ID     Used   Soft   Hard Warn/Grace
---------- ---------------------------------
root            0      0      0  00 [------]
devs            0   100G   500G  00 [------]
```

```
[bob@centos-host ~]$ sudo quota -g devs
Disk quotas for group devs (gid 1004):
Filesystem  blocks   quota   limit   grace   files   quota   limit   grace
/dev/vdb1       0  104857600 524288000               1       0       0
```

- [x] Make sure all users under "admins" group can run all commands with "sudo" and without entering any password.

```
[bob@centos-host ~]$ sudo visudo
```

```
[bob@centos-host ~]$ sudo tail -v -n 2 /etc/sudoers
==> /etc/sudoers <==
# admins rules
%admins ALL=(ALL) NOPASSWD: ALL
```

- [x] Make sure all users under "devs" group can only run the "dnf" command with "sudo" and without entering any password.

```
[bob@centos-host ~]$ sudo visudo
```

```
[bob@centos-host ~]$ sudo tail -v -n 2 /etc/sudoers
# devs rules
%devs ALL=(ALL) NOPASSWD: /usr/bin/dnf
```

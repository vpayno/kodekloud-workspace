# KodeKloud - Linux Challenge - Challenge 4 Output

- [x] Inspect OS

```
[bob@centos-host ~]$ uname -r
5.4.0-1075-gcp
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

- [x] Create "/opt/appdata" directory.

```
[bob@centos-host ~]$ sudo mkdir -pv /opt/appdata
mkdir: created directory '/opt/appdata'
```

- [x] Do not delete any files from "/home/bob/preserved" directory.

```
[bob@centos-host ~]$ find /home/bob/preserved/ -type f | wc -l
9206
```

```
[bob@centos-host ~]$ ls /home/bob/preserved/.[a-z0-9A-Z]* | wc -l
1200
```

```
[bob@centos-host ~]$ ls /home/bob/preserved/[a-z0-9A-Z]* | wc -l
8006
```

```
[bob@centos-host ~]$ echo $(( 8006 + 1200 ))
9206
```

```
[bob@centos-host ~]$ du -shc /home/bob/preserved
8.9M    /home/bob/preserved
8.9M    total
```

- [x] Find the "hidden" files in "/home/bob/preserved" directory and copy them in "/opt/appdata/hidden/" directory (create the destination directory if doesn't exist).

```
[bob@centos-host ~]$ sudo mkdir -pv /opt/appdata/hidden/
mkdir: created directory '/opt/appdata/hidden/'
```

```
[bob@centos-host ~]$ ls /home/bob/preserved/.[a-z0-9A-Z]* | wc -l
1200
```

```
[bob@centos-host ~]$ sudo cp -vp /home/bob/preserved/.[a-z0-9A-Z]* /opt/appdata/hidden/ | wc -l
1200
```

```
[bob@centos-host ~]$ ls /opt/appdata/hidden/.[a-z0-9A-Z]* | wc -l
1200
```

- [x] Find the "non-hidden" files in "/home/bob/preserved" directory and copy them in "/opt/appdata/files/" directory (create the destination directory if doesn't exist).

```
[bob@centos-host ~]$ sudo mkdir -pv /opt/appdata/files/
mkdir: created directory '/opt/appdata/files/'
```

```
[bob@centos-host ~]$ find /home/bob/preserved/ -type f -name '[a-z0-9A-Z]*' | sudo xargs cp -vp -t /opt/appdata/files/ | wc -l
8006
```

```
[bob@centos-host ~]$ find /opt/appdata/files/ -type f | wc -l
8006
```

- [x] Find and delete the files in "/opt/appdata" directory that contain a word ending with the letter "t" (case sensitive).

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E '[a-z]+t\b' | head
/opt/appdata/hidden/.hfile109:My raw datat
/opt/appdata/hidden/.hfile047:My raw datat
/opt/appdata/hidden/.hfile011:My raw datat
/opt/appdata/hidden/.hfile143:My raw datat
/opt/appdata/hidden/.hfile113:My raw datat
/opt/appdata/hidden/.hfile031:My raw datat
/opt/appdata/hidden/.hfile020:My raw datat
/opt/appdata/hidden/.hfile023:My raw datat
/opt/appdata/hidden/.hfile119:My raw datat
/opt/appdata/hidden/.hfile035:My raw datat
```

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E '[a-z]+t\b' | cut -f1 -d: | sort -u | wc -l
300
```

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E '[a-z]+t\b' | cut -f1 -d: | sort -u | sudo xargs rm -v | wc -l
300
```

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E 't\b' | cut -f1 -d: | sort -u | wc -l
0
```

- [x] Change all the occurrences of the word "yes" to "no" in all files present under "/opt/appdata/" directory.

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E '\byes\b' | head
/opt/appdata/hidden/.hfile092:yes maybe yes
/opt/appdata/hidden/.hfile057:yes maybe yes
/opt/appdata/hidden/.hfile088:yes maybe yes
/opt/appdata/hidden/.hfile051:yes maybe yes
/opt/appdata/hidden/.hfile061:yes maybe yes
/opt/appdata/hidden/.hfile099:yes maybe yes
/opt/appdata/hidden/.hfile086:yes maybe yes
/opt/appdata/hidden/.hfile066:yes maybe yes
/opt/appdata/hidden/.hfile054:yes maybe yes
/opt/appdata/hidden/.hfile055:yes maybe yes
```

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E '\byes\b' | cut -f1 -d: | sort -u | wc -l
50
```

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E '\byes\b' | cut -f1 -d: | sort -u | sudo xargs sed -r -i -e 's/\byes\b/no/g'
```

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E '\byes\b' | cut -f1 -d: | sort -u | wc -l
0
```

- [x] Change all the occurrences of the word "raw" to "processed" in all files present under "/opt/appdata/" directory. It must be a "case-insensitive" replacement, means all words must be replaced like "raw , Raw , RAW" etc.

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E '\b[rR][aA][wW]\b' | head
/opt/appdata/hidden/.hfile173:My raw data
/opt/appdata/hidden/.hfile188:My raw data
/opt/appdata/hidden/.hfile159:My raw data
/opt/appdata/hidden/.hfile183:My raw data
/opt/appdata/hidden/.hfile197:My raw data
/opt/appdata/hidden/.hfile165:My raw data
/opt/appdata/hidden/.hfile182:My raw data
/opt/appdata/hidden/.hfile180:My raw data
/opt/appdata/hidden/.hfile153:My raw data
/opt/appdata/hidden/.hfile158:My raw data
```

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E '\b[rR][aA][wW]\b' | cut -f1 -d: | sort -u | wc -l
1850
```

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E '\b[rR][aA][wW]\b' | cut -f1 -d: | sort -u | sudo xargs sed -r -i -e 's/\b[rR][aA][wW]\b/processed/g'
```

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -E '\b[rR][aA][wW]\b' | cut -f1 -d: | sort -u | wc -l
0
```

- [x] Add the "sticky bit" special permission on "/opt/appdata" directory (keep the other permissions as it is).

```
[bob@centos-host ~]$ sudo chmod -v +t /opt/appdata/
mode of '/opt/appdata/' changed from 0755 (rwxr-xr-x) to 1755 (rwxr-xr-t)
```

- [x] Create a "tar.gz" archive of "/opt/appdata" directory and save the archive to this file: "/opt/appdata.tar.gz"

```
[bob@centos-host ~]$ sudo tar czf /opt/appdata.tar.gz /opt/appdata
tar: Removing leading `/' from member names
```

```
[bob@centos-host ~]$ file /opt/appdata.tar.gz
/opt/appdata.tar.gz: gzip compressed data, last modified: Sun Jun  5 23:06:58 2022, from Unix, original size 6318080
```

- [x] The "appdata.tar.gz" archive should have the final processed data.

```
[bob@centos-host ~]$ tar tzf /opt/appdata.tar.gz | wc -l
8909
```

- [x] Create a "softlink" called "/home/bob/appdata.tar.gz" of "/opt/appdata.tar.gz" file.

```
[bob@centos-host ~]$ ln -sv /opt/appdata.tar.gz /home/bob/appdata.tar.gz
'/home/bob/appdata.tar.gz' -> '/opt/appdata.tar.gz'
```

```
[bob@centos-host ~]$ file appdata.tar.gz
appdata.tar.gz: symbolic link to /opt/appdata.tar.gz
```

- [x] Make "bob" the "user" and the "group" owner of "/opt/appdata.tar.gz" file.

```
[bob@centos-host ~]$ sudo chown -v bob:bob /opt/appdata.tar.gz
changed ownership of '/opt/appdata.tar.gz' from root:root to bob:bob
```

- [x] The "user/group" owner should have "read only" permissions on "/opt/appdata.tar.gz" file and "others" should not have any permissions.

```
[bob@centos-host ~]$ sudo chmod -v 0440 /opt/appdata.tar.gz
mode of '/opt/appdata.tar.gz' changed from 0644 (rw-r--r--) to 0440 (r--r-----)
```

- [x] Create a script called "/home/bob/filter.sh".

```
[bob@centos-host ~]$ cat > /home/bob/filter.sh <<EOF
>#!/bin/bash
>
>EOF
```

```
[bob@centos-host ~]$ chmod -v a+x /home/bob/filter.sh
mode of '/home/bob/filter.sh' changed from 0664 (rw-rw-r--) to 0775 (rwxrwxr-x)
```

- [ ] The script should filter the lines from "/opt/appdata.tar.gz" file which contain the word "processed", and save the filtered output in "/home/bob/filtered.txt" file. It must "overwrite" the existing contents of "/home/bob/filtered.txt" file

Note:
- I can't figure out what it wants. Is it greping the filter.sh script for something?
- I'm not even sure why we want to have a file called filtered.txt with the file list of files that have processed in them. Or the opposite, the file list without those files.

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | wc -l
8906
```

```
[bob@centos-host ~]$ find /opt/appdata/ -type f | xargs grep -v -E '\bprocessed\b' | cut -f1 -d: | sort -u | wc -l
50
```

```
[bob@centos-host ~]$ cat >> /home/bob/filter.sh <<EOF
>mkdir -pv /tmp/deleteme
>
>tar -C /tmp/deleteme/ -xzf /opt/appdata.tar.gz
>
>find /tmp/deleteme/ -type f | xargs grep -v -E '\bprocessed\b' | cut -f1 -d: | sort -u > /home/bob/filtered.txt
>
>rm -rf /tmp/deleteme
>EOF
```

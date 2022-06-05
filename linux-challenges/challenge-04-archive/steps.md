# https://kodekloud.com/topic/linux-challenge-4/

Step overview for challenge 4.

```
$ uname -r
$ cat /etc/os-release
```

```
$ sudo mkdir -pv /opt/appdata
```

```
$ find /home/bob/preserved/ -type f | wc -l
$ ls /home/bob/preserved/.[a-z0-9A-Z]* | wc -l
$ ls /home/bob/preserved/[a-z0-9A-Z]* | wc -l
$ echo $(( 8006 + 1200 ))
$ du -shc /home/bob/preserved
```

```
$ sudo mkdir -pv /opt/appdata/hidden/
$ ls /home/bob/preserved/.[a-z0-9A-Z]* | wc -l
$ sudo cp -p /home/bob/preserved/.[a-z0-9A-Z]* /opt/appdata/hidden/
$ ls /opt/appdata/hidden/.[a-z0-9A-Z]* | wc -l
```

```
$ sudo mkdir -pv /opt/appdata/files/
$ find /home/bob/preserved/ -type f -name '[a-z0-9A-Z]*' | sudo xargs cp -vp -t /opt/appdata/files/ | wc -l
$ find /opt/appdata/files/ -type f | wc -l
$ find /opt/appdata/ -type f | xargs grep -E '[a-z]+t\b' | head
$ find /opt/appdata/ -type f | xargs grep -E '[a-z]+t\b' | cut -f1 -d: | sort -u | wc -l
$ find /opt/appdata/ -type f | xargs grep -E '[a-z]+t\b' | cut -f1 -d: | sort -u | sudo xargs rm -v | wc -l
$ find /opt/appdata/ -type f | xargs grep -E 't\b' | cut -f1 -d: | sort -u | wc -l
```

```
$ find /opt/appdata/ -type f | xargs grep -E '\byes\b' | head
$ find /opt/appdata/ -type f | xargs grep -E '\byes\b' | cut -f1 -d: | sort -u | wc -l
$ find /opt/appdata/ -type f | xargs grep -E '\byes\b' | cut -f1 -d: | sort -u | sudo xargs sed -r -i -e 's/\byes\b/no/g'
$ find /opt/appdata/ -type f | xargs grep -E '\byes\b' | cut -f1 -d: | sort -u | wc -l
```

```
$ find /opt/appdata/ -type f | xargs grep -E '\b[rR][aA][wW]\b' | head
$ find /opt/appdata/ -type f | xargs grep -E '\b[rR][aA][wW]\b' | cut -f1 -d: | sort -u | wc -l
$ find /opt/appdata/ -type f | xargs grep -E '\b[rR][aA][wW]\b' | cut -f1 -d: | sort -u | sudo xargs sed -r -i -e 's/\b[rR][aA][wW]\b/processed/g'
$ find /opt/appdata/ -type f | xargs grep -E '\b[rR][aA][wW]\b' | cut -f1 -d: | sort -u | wc -l
```

```
$ sudo chmod -v +t /opt/appdata/
$ sudo tar czf /opt/appdata.tar.gz /opt/appdata
$ file /opt/appdata.tar.gz
$ tar tzf /opt/appdata.tar.gz | wc -l
```

```
$ ln -sv /opt/appdata.tar.gz /home/bob/appdata.tar.gz
$ file appdata.tar.gz
```

```
$ sudo chown -v bob:bob /opt/appdata.tar.gz
$ sudo chmod -v 0440 /opt/appdata.tar.gz
```

```
$ cat > /home/bob/filter.sh <<EOF
#!/bin/bash

EOF

$ chmod -v a+x /home/bob/filter.sh
```

```
$ find /opt/appdata/ -type f | xargs grep -v -E '\bprocessed\b' | cut -f1 -d: | sort -u | wc -l
```

```
$ cat >> /home/bob/filter.sh <<EOF
mkdir -pv /tmp/deleteme

tar -C /tmp/deleteme/ -xzf /opt/appdata.tar.gz

find /tmp/deleteme/ -type f | xargs grep -v -E '\bprocessed\b' | cut -f1 -d: | sort -u > /home/bob/filtered.txt

rm -rf /tmp/deleteme
EOF

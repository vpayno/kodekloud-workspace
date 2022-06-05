# https://kodekloud.com/topic/linux-challenge-3/

Step overview for challenge 3.

```
$ uname -r
$ cat /etc/os-release
```

```
$ sudo groupadd admins
$ sudo tail -v -n 1 /etc/group
```

```
$ sudo groupadd devs
$ sudo tail -v -n 2 /etc/group
```

```
$ sudo chown -v bob /data
$ sudo chown -v bob:devs /data
$ sudo chmod -v 0770 /data
```

```
$ getfacl /data
$ sudo setfacl -d -m g:admins:rwx /data
$ getfacl /data
```

```
$ sudo useradd -c "Ray" -s /bin/sh -G devs ray
$ sudo passwd ray
$ sudo grep ray /etc/passwd /etc/group

$ sudo useradd -c "Lisa" -s /bin/sh -G devs lisa
$ sudo passwd lisa
$ sudo grep lisa /etc/passwd /etc/group

$ sudo useradd -c "Natasha" -s /bin/zsh -G admins natasha
$ sudo passwd natasha
$ sudo grep natasha /etc/passwd /etc/group

$ sudo useradd -c "David" -s /bin/zsh -G admins david
$ sudo passwd david
$ sudo grep david /etc/passwd /etc/group
```

```
$ sudo bash -c 'cat >> /etc/limits.conf' <<EOF
$ grep nproc /etc/limits.conf
```

```
$ mount | grep /data
$ sudo xfs_quota -x -c 'limit -g bsoft=100g bhard=500g devs' /data
$ sudo xfs_quota -x -c 'report -h -g' /data
$ sudo quota -g devs
```

```
$ sudo visudo
$ sudo tail -v -n 2 /etc/sudoers

$ sudo visudo
$ sudo tail -v -n 2 /etc/sudoers
```

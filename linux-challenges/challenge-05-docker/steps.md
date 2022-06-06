# https://kodekloud.com/topic/linux-challenge-5/

Step overview for challenge 5.

```
$ uname -r
$ cat /etc/os-release
```

```
$ sudo yum install -y mariadb mariadb-server
```

```
$ sudo systemctl enable mariadb
$ sudo systemctl start mariadb
$ sudo mysql -u root <<EOF
$ sudo mysql -u root -p <<EOF
```

```
$ sudo passwd -u root -u
$ sudo usermod -a -G wheel root
$ id root
```

```
$ grep wheel /etc/pam.d/su
$ sudoedit /etc/pam.d/su
$ grep wheel /etc/pam.d/su
```

```
$ ip addr
$ sudo grep 'NM_CONTROLLED' /etc/sysconfig/network-scripts/ifcfg-eth1
$ sudo nmtui
$ ip addr
```

```
$ sudo bash -c 'cat >> /etc/hosts' <<EOF
$ ping -c 3 mydb.kodekloud.com
```

```
$ sudo docker pull nginx
$ sudo docker images
$ sudo docker ps
$ sudo docker run --name myapp -d -p 80:80 nginx:latest
$ sudo docker ps
```

```
$ cat > container-start.sh <<EOF
#!/bin/bash

sudo docker start myapp

printf "%s\n" "myapp container started!"
EOF

$ chmod -v a+x ./container-start.sh
```

```
$ cat > container-stop.sh <<EOF
#!/bin/bash

sudo docker stop myapp

printf "%s\n" "myapp container stopped!"
EOF

$ chmod -v a+x ./container-stop.sh
```

```
$ sudo crontab -u root -e
> 0  0  *  *  * root  /home/bob/container-stop.sh
$ sudo crontab -l -u root
```

```
$ sudo crontab -u root -e
> 0  8  *  *  * root  /home/bob/container-start.sh
$ sudo crontab -l -u root
```

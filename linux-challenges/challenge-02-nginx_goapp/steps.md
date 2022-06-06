# https://kodekloud.com/topic/linux-challenge-2/

Step overview for challenge 2.

```
$ uname -r
$ cat /etc/os-release

$ sudo yum install -y tmux net-tools nginx firewalld

$ cat /etc/resolv.conf
$ sudoedit /etc/resolv.conf
$ tail -v /etc/resolv.conf

$ sudo yum install -y tmux net-tools nginx firewalld

$ sudo systemctl enable firewalld
$ sudo systemctl start firewalld
$ sudo firewall-cmd --permanent --zone=public --list-ports
$ sudo firewall-cmd --permanent --zone=public --add-port=22/tcp
$ sudo firewall-cmd --permanent --zone=public --add-port=80/tcp
$ sudo firewall-cmd --permanent --zone=public --add-port=8081/tcp
$ sudo firewall-cmd --reload
$ sudo firewall-cmd --permanent --zone=public --list-ports

$ sudo bash -c 'cat > /etc/nginx/conf.d/goapp.conf' <<EOF
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

$ sudo systemctl enable nginx
$ sudo systemctl start nginx
$ curl -v --user test:test http://localhost/

$ cd go-app/
$ nohup go run main.go &
$ sleep 5s #
$ tail -v nohup.out
$ curl -v --user test:test http://localhost/
$ sudo netstat -tlnp
```

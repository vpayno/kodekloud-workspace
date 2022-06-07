# https://kodekloud.com/lessons/project-e-commerce-application/

Step overview for e-commerce lab.

They used a vm/container on katacode for the demo.

```
sudo yum install firewalld
sudo service firewalld start
sudo systemctl enable firewalld

sudo yum install mariadb-server
# change port
sudoedit /etc/my.cnf
sudo service mariadb start
sudo systemctl enable mariadb

sudo firewall-cmd --permanent --zone=public --add-port=3306/tcp
sudo firewall-cmd --reload

mysql <<EOF
CREATE DATABASE ecomdb;
CREATE USER 'ecomuser'@'localhost' IDENTIFIED BY 'ecompassword';
GRANT ALL PRIVILEGES ON *.* TO 'econuser'@'localhost';
FLUSH PRIVILEGES;
EOF

mysql < db-load-script.sql
```

```
sudo yum install -y httpd php php-mysql
sudo firewall-cmd --permament --zone=public --add-port=80/tcp
sudo firewall-cmd --reload

# use index.php instead of index.html
sudo vi /etc/httpd/conf/httpd.conf

sudo serviec httpd start
sudo systemctl enable httpd

sudo yum install -y git
git clone https://github.com/${app_name}.git /var/www/html/

# update index.php to use the correct db addr, name and pw.

curl -v http://localhost/
```

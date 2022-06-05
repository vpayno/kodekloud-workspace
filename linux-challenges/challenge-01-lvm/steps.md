# https://kodekloud.com/topic/linux-challenge-1/

Step overview for challenge 1.

```
sudo yum install lvm2 tmux

sudo groupadd dba_users
sudo usermod -a -G dba_users bob
sudo grep dba_users /etc/group

tmux

sudo parted /dev/vdb p
sudo parted /dev/vdc p
sudo parted /dev/vdb -s -- mklabel gpt p mkpart lvm 1m -1 set 1 lvm on p
sudo parted /dev/vdc -s -- mklabel gpt p mkpart lvm 1m -1 set 1 lvm on p

sudo pvcreate /dev/vd{b,c}1
sudo pvs

sudo vgcreate dba_storage /dev/vd{b,c}1
sudo vgs
sudo pvs

sudo lvcreate --name volume_1 --extents 100%VG dba_storage
sudo lvs
sudo vgs
sudo pvs

sudo mkfs.xfs -L dba_storage /dev/mapper/dba_storage-volume_1

sudo bash -c 'printf "UUID=%s        /mnt/dba_storage        xfs        defaults        0 0\n" "$(lsblk --fs /dev/mapper/dba_storage-volume_1 | awk "/dba_storage/ { print \$4 }")" >> /etc/fstab'
tail -v /etc/fstab

sudo mkdir /mnt/dba_storage
sudo mount -av
sudo restorecon -R /mnt/dba_storage
df -h

sudo chgrp -v dba_users /mnt/dba_storage/
sudo chmod -v 0770 /mnt/dba_storage/
ls -dl /mnt/dba_storage/
```

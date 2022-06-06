# KodeCloud - BashX Commands

```
bob@caleston-lp10:~$ rocket-ls

--------------------------------------------
          ROCKET MISSIONS
--------------------------------------------
  Name                   Status
--------------------------------------------

          Total Missions - 0
```

```
bob@caleston-lp10:~$ mkdir lunar-mission

bob@caleston-lp10:~$ rocket-add lunar-mission

--------------------------------------------
          PROJECT lunar-mission
--------------------------------------------
Creating a new rocket....Done!

bob@caleston-lp10:~$ find lunar-mission/
lunar-mission/
lunar-mission/status

bob@caleston-lp10:~$ cat lunar-mission/status
created
```

```
bob@caleston-lp10:~$ rocket-start-power lunar-mission

Starting power ....Done!

bob@caleston-lp10:~$ find lunar-mission/
lunar-mission/
lunar-mission/status

bob@caleston-lp10:~$ cat lunar-mission/status
start-power
```

```
bob@caleston-lp10:~$ rocket-internal-power lunar-mission

Switching to internal ....Done!

bob@caleston-lp10:~$ find lunar-mission/
lunar-mission/
lunar-mission/status

bob@caleston-lp10:~$ cat lunar-mission/status
internal-power
```

```
bob@caleston-lp10:~$ rocket-start-sequence lunar-mission

Starting auto sequence ....Done!

bob@caleston-lp10:~$ find lunar-mission/
lunar-mission/
lunar-mission/status

bob@caleston-lp10:~$ cat lunar-mission/status
start-sequence
```

```
bob@caleston-lp10:~$ rocket-start-engine lunar-mission

Starting engine ....Done!

bob@caleston-lp10:~$ find lunar-mission/
lunar-mission/
lunar-mission/status

bob@caleston-lp10:~$ cat lunar-mission/status
start-engine
```

```
bob@caleston-lp10:~$ rocket-lift-off lunar-mission

Initiating lift off ....
        Countdown
          10
          9
          8
          7
          6
...
bob@caleston-lp10:~$ find lunar-mission/
lunar-mission/
lunar-mission/status

bob@caleston-lp10:~$ cat lunar-mission/status
success
```

```
bob@caleston-lp10:~$ rocket-status lunar-mission
success

bob@caleston-lp10:~$ find lunar-mission/
lunar-mission/
lunar-mission/status

bob@caleston-lp10:~$ cat lunar-mission/status
success
```

```
bob@caleston-lp10:~$ rocket-ls

--------------------------------------------
          ROCKET MISSIONS
--------------------------------------------
  Name                   Status
--------------------------------------------
lunar-mission            success
mars-mission             success

          Total Missions - 2
```

```
bob@caleston-lp10:~$ bash /usr/local/bin/_test-script.sh
mkdir: cannot create directory 'lunar-mission': File exists

--------------------------------------------
          PROJECT lunar-mission
--------------------------------------------
Creating a new rocket....Done!

Starting power ....Done!

Switching to internal ....Done!

Starting auto sequence ....Done!

Starting engine ....Done!

Initiating lift off ....
        Countdown
          10
          9
```

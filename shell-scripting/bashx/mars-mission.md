# Mars Mission

```
bob@caleston-lp10:~$ for cmd in rocket-add rocket-start-power rocket-internal-power rocket-start-sequence rocket-start-engine rocket-lift-off rocket-status; do "${cmd}" mars-mission; tail -v mars-mission/status; printf "\n"; done

--------------------------------------------
          PROJECT mars-mission
--------------------------------------------
Creating a new rocket....Done!
==> mars-mission/status <==
created


Starting power ....Done!
==> mars-mission/status <==
start-power


Switching to internal ....Done!
==> mars-mission/status <==
internal-power


Starting auto sequence ....Done!
==> mars-mission/status <==
start-sequence


Starting engine ....Done!
==> mars-mission/status <==
start-engine


Initiating lift off ....
        Countdown
          10
          9
```

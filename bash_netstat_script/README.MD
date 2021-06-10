### IP information script
The script is intended to get information about the IP addresses that processes established connection with.You can receive additional information as an option( you will be prompted to it during the script execution)
execution example:You can enter the "ENTER" key to get the default value
```
Enter the PID/Process Name: firefox \n
What types of connections would you like to see?
>  l - for LISTENING , e - for ESTABLISHED , (all-default,any other letter or ENTER) 
You'd choosed to type everything
To see all the connections you must login as root. Do you need such an option? (y/n?) default -n
USERTYPE IS n
How many lines would you like to see as output/press Enter to default all-print 
TEST_________________________________________________________________
      1 34.210.128.24  ESTABLISHED 3681/firefox
      2 87.250.251.210  ESTABLISHED 3681/firefox
      ____________________________________________________________________
   Amazon Technologies Inc. (AT-88-Z)Connection Counts :1 IP:34.210.128.24 State:ESTABLISHED PID/NAME:3681/firefox

Connection Counts :2 IP:87.250.251.210 State:ESTABLISHED PID/NAME:3681/firefox
````

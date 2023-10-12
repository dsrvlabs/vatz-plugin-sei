# Script for vatz & vatz-plugin-sei
## 1. install_vatz-plugin-sei.sh
- Run this script to install vatz and vatz-plugin-sei. 
```
bash install_vatz-plugin-sei.sh
```
## 2. default_example.yaml
- Replace default.yaml with this file after `vatz init`.
- You must enter hostname and webhook.
- Change the port if necessary.
```
cp default_example.yaml /root/vatz/default.yaml
```
## 3. vatz_start.sh
- Running this script will run vatz and vatz-plugin-sei. 
- The log is stored in `/var/log/vatz`.
  - You can change the log path if necessary.
- Enter valoper_address and voter_address.
- Modify sei home to suit your needs.
```
bash vatz_start.sh
```
## 4. vatz_stop.sh
- Running this script will stop both vatz and vatz-plugin-sei as they are currently running.
```
bash vatz_stop.sh
```

echo 100 facebook >> /etc/iproute2/rt_tables

ip route add default dev tun0 table facebook

  ipset create blockedIP  iphash -exist
  ipset add blockedIP  8.8.2.2

iptables -t mangle -I prerouting -m set --match-set blockedIP dst -j  Mark --set-mark 998


ip rule add fwmark 998 priority 1999 table facebook

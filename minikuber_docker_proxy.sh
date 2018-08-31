# cat proxy_minikuber_docker.sh| minikube ssh

#vi /usr/lib/systemd/system/docker.service

#Environment="DOCKER_RAMDISK=yes" "HTTP_PROXY=http://192.168.99.1:1087/" "HTTPS_PROXY=http://192.168.99.1:1087/"

ip=`ifconfig| grep -E '\binet\b' | grep 10  | cut -d' ' -f 2`
line="Environment=\"DOCKER_RAMDISK=yes\" \"HTTP_PROXY=http://$ip:1087/\" \"HTTPS_PROXY=http://$ip:1087/\""

#cat << EOF | ssh jd 
#echo $ip $line
#echo $line >> ss
#date >> ss
#EOF

cat << EOF 
sudo -i
sed -i s+^Environment.*+"$line"+ /usr/lib/systemd/system/docker.service 
systemctl daemon-reload
systemctl restart docker
exit
exit
EOF

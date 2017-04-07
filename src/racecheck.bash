#!/bin/bash

export NOTI_DEFAULT=banner

echo "all.bash"
./all.bash
noti -t racecheck -m "all.bash done"

echo "install race cmd/compile"
go install -race cmd/compile
noti -t racecheck -m "install -race done"


echo "HOST"
go build -a -gcflags="-c=128" std cmd
go build -a -gcflags="-c=128 -N -l" std cmd
go build -a -gcflags="-c=128 -race" std cmd
noti -t racecheck -m "HOST done"

echo "linux/ppc64"
GOOS=linux GOARCH=ppc64 go build -a -gcflags="-c=128" std cmd 

echo "linux/ppc64le"
GOOS=linux GOARCH=ppc64le go build -a -gcflags="-c=128" std cmd 


echo "linux/s390x"
GOOS=linux GOARCH=s390x go build -a -gcflags="-c=128" std cmd 
echo "linux/arm64"
GOOS=linux GOARCH=arm64 go build -a -gcflags="-c=128" std cmd 
echo "linux/mips"
GOOS=linux GOARCH=mips go build -a -gcflags="-c=128" std cmd 
echo "linux/arm"
GOOS=linux GOARCH=arm go build -a -gcflags="-c=128" std cmd 
echo "nacl/arm"
GOOS=nacl GOARCH=arm go build -a -gcflags="-c=128" std cmd 
echo "linux/arm5"
GOOS=linux GOARCH=arm GOARM=5 go build -a -gcflags="-c=128" std cmd 
echo "linux/arm6"
GOOS=linux GOARCH=arm GOARM=6 go build -a -gcflags="-c=128" std cmd 
echo "linux/arm7"
GOOS=linux GOARCH=arm GOARM=7 go build -a -gcflags="-c=128" std cmd 
echo "nacl/arm5"
GOOS=nacl GOARCH=arm GOARM=5 go build -a -gcflags="-c=128" std cmd 
echo "nacl/arm6"
GOOS=nacl GOARCH=arm GOARM=6 go build -a -gcflags="-c=128" std cmd 
echo "nacl/arm7"
GOOS=nacl GOARCH=arm GOARM=7 go build -a -gcflags="-c=128" std cmd 

noti -t racecheck -m "other architectures done"

echo "checking reproducibility"
./make.bash
toolstash save
for i in `seq 20`
do
	go build -a -toolexec='toolstash -cmp' std cmd
	echo "ok"
done
noti -t racecheck -m "reproducibility done"

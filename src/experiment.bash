#!/bin/bash

while true
do
	LIKELY=`python3 -c "import numpy; print(1+int(numpy.random.exponential(20)))"`
	NORMAL=`python3 -c "import numpy; print($LIKELY+int(numpy.random.exponential(30)))"`
	UNLIKELY=`python3 -c "import numpy; print($NORMAL+int(numpy.random.exponential(50)))"`
	J=1
	OUT=`J=1 LIKELY=$LIKELY NORMAL=$NORMAL UNLIKELY=$UNLIKELY go build -a std cmd 2>&1 | wc -l`
	# go build -a -o x cmd/go
	# go build -a -o y cmd/compile
	# go build -a -o z cmd/pprof
	# go build -a -o w cmd/trace
	# go build -a -o v cmd/vet
	# echo $LIKELY $'\t' $NORMAL $'\t' $UNLIKELY $'\t' `ls -s x y z w v`
	echo $LIKELY $'\t' $NORMAL $'\t' $UNLIKELY $'\t' $OUT
done

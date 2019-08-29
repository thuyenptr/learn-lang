#!/bin/bash
WORKING_DIR=`pwd`

protoc -I ${WORKING_DIR}/proto/binlog ${WORKING_DIR}/proto/binlog/*.proto --go_out=plugins=grpc:${WORKING_DIR}/proto

cd ${WORKING_DIR}/proto
ls
#sed -i.bak -E 's/import _ \"gogoproto\"//g' *.pb.go
#sed -i.bak -E 's/import fmt \"fmt\"//g' *.pb.go
#rm -f *.bak
#goimports -w *.pb.go
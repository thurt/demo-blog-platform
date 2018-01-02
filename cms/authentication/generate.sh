#! /bin/bash

#generate mc.Conn interface
ifacemaker -f $GOPATH/src/github.com/memcachier/mc/mc.go -s Conn -i Conn -p authentication -o ./mc.go && \
#generate mc.Conn mock
mockgen -package=mock_mc -source=./mc.go >  ../mock_mc/mock_mc.go 


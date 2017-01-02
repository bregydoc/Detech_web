#!/usr/bin/env bash

export GOPATH=$HOME/goLibraries/

cd PublicGo

$HOME/goLibraries/bin/gopherjs build dashboardScript.go -o ../Public/dashboardScript.js
$HOME/goLibraries/bin/gopherjs build newPatientScript.go -o ../Public/newPatientScript.js

cd ..

cd Backend
go build -o ../Detech30
cd ..

./Detech30

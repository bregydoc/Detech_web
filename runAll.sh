#!/usr/bin/env bash

export GOPATH=$HOME/goLibraries/

cd PublicGo

$HOME/goLibraries/bin/gopherjs build -m dashboardScript.go -o ../Public/dashboardScript.js
$HOME/goLibraries/bin/gopherjs build -m newPatientScript.go -o ../Public/newPatientScript.js
$HOME/goLibraries/bin/gopherjs build -m patientDashboardScript.go -o ../Public/patientDashboardScript.js
$HOME/goLibraries/bin/gopherjs build -m newClinicHistoryScript.go -o ../Public/newClinicHistoryScript.js
$HOME/goLibraries/bin/gopherjs build -m evaluationFileScript.go -o ../Public/evaluationFileScript.js


cd ..

cd Backend
go build -o ../Detech30
cd ..

./Detech30

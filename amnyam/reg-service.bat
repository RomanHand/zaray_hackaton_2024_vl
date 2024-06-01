chcp 65001
@echo off

set SrvcName="zarya_backend"
set CurrentDir=%~dp0

set BinPath="%CurrentDir%back.exe"

sc stop %SrvcName%
sc delete %SrvcName%
sc create %SrvcName% binPath= %BinPath% start= demand displayname= %SrvcName%
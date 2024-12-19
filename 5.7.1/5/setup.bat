@echo off
:: Check if script is with administrator privileges
net session >nul 2>&1
if %errorlevel% neq 0 (
    echo This script must be run with administrator privileges!
    exit /b 1
)
:: Create schedule
schtasks /create /tn "RunAsSystem" /tr "cmd.exe /c echo HD{Endi_Sen_Men_Bo'lding} > C:\Users\Haady\Desktop\SystemFlag.txt" /sc once /st 00:00 /ru "SYSTEM"

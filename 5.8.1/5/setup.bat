@echo off

:: Check if script is running with admin rights
net session >nul 2>&1
if %errorlevel% neq 0 (
    echo The script is NOT running with administrator privileges. Please run this script as an administrator!
    exit /b 1
)

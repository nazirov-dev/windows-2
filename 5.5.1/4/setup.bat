@echo off

:: Check script is running with admin privileges
net session >nul 2>&1
if %errorlevel% neq 0 (
    echo You need to run this script with admin privileges.
    pause
    exit /b
)

:: Run PowerShell with elevated privileges to change the creation date
(Get-Item "C:\\Windows\\System32\\netapi32.dll").CreationTime = $(Get-Date "20/09/2021 17:07:21")

:: Check if the PowerShell command succeeded
if %errorlevel% neq 0 (
    echo Failed to change creation date of netapi32.dll.
    pause
    exit /b
) else (
    echo Successfully changed creation date of netapi32.dll.
    pause
)

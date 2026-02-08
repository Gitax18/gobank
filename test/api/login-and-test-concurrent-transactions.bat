@echo off
setlocal EnableDelayedExpansion

set BASE_URL=http://localhost:8080
set COOKIE_FILE=cookies.txt

set LOGIN_EMAIL=john@example.com
set LOGIN_PASSWORD=john123

set RECEIVER_ID=7
set AMOUNT=100
set TOTAL_REQUESTS=100

echo ======================================
echo STEP 1: Logging in user
echo ======================================

curl -s -X POST %BASE_URL%/user/login ^
    -H "Content-Type: application/json" ^
    --cookie-jar %COOKIE_FILE% ^
    --data-raw "{\"email\":\"%LOGIN_EMAIL%\",\"password\":\"%LOGIN_PASSWORD%\"}"

echo.
echo Login completed. Cookies saved to %COOKIE_FILE%

echo.
echo ======================================
echo STEP 2: Sending %TOTAL_REQUESTS% concurrent transactions
echo Receiver ID: %RECEIVER_ID%
echo Amount per transaction: %AMOUNT%
echo ======================================

for /L %%i in (1,1,%TOTAL_REQUESTS%) do (
    start "" /B cmd /c ^
    curl -s -X POST %BASE_URL%/transaction ^
        -H "Content-Type: application/json" ^
        --cookie %COOKIE_FILE% ^
        --data-raw "{\"receiver_id\":%RECEIVER_ID%,\"amount\":%AMOUNT%}"
)

echo.
echo 🚀 All %TOTAL_REQUESTS% transaction requests dispatched concurrently
echo Check DB balances and server logs

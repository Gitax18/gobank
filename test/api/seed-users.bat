@echo off
setlocal EnableDelayedExpansion

set URL=http://localhost:8080/user/register

set NAME1=Rohit Sharma
set EMAIL1=rohit.sharma@gmail.com
set PHONE1=9876543210
set ACC1=1000000001

set NAME2=Ananya Verma
set EMAIL2=ananya.verma@gmail.com
set PHONE2=9123456789
set ACC2=1000000002

set NAME3=Karan Mehta
set EMAIL3=karan.mehta@gmail.com
set PHONE3=9988776655
set ACC3=1000000003

set NAME4=Pooja Nair
set EMAIL4=pooja.nair@gmail.com
set PHONE4=9090909090
set ACC4=1000000004

set NAME5=Arjun Singh
set EMAIL5=arjun.singh@gmail.com
set PHONE5=9012345678
set ACC5=1000000005

for /L %%i in (1,1,5) do (

    set /a BALANCE=!RANDOM! %% 1001 * 10

    call set NAME=%%NAME%%i%%
    call set EMAIL=%%EMAIL%%i%%
    call set PHONE=%%PHONE%%i%%
    call set ACCOUNT=%%ACC%%i%%

    echo Creating user %%i: !NAME! with balance !BALANCE!

    curl -s -X POST %URL% ^
      -H "Content-Type: application/json" ^
      --data-raw "{\"email\":\"!EMAIL!\",\"password\":\"john123\",\"name\":\"!NAME!\",\"number\":!PHONE!,\"account_number\":!ACCOUNT!,\"balance\":!BALANCE!}"

    echo.
)

echo.
echo ✅ Successfully created 5 realistic users.

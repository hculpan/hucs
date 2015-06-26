@echo off

set GOFILES=

for %%f in (*.go) do call :concat %%f

goto :run

:concat
set GOFILES=%GOFILES% %1
goto :eof

:run
go run %GOFILES% %*

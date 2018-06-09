@echo off
for /L %%A in (1,1,40) do (
    hmac-sha256-go.exe "https://www.bitmex.com/api/v1/trade/bucketed?binSize=5m&partial=false&symbol=XBTUSD&count=20&reverse=false&startTime=2018-06-02T16%3A20%3A00.000Z'" 8b5f48702995c1598c573db1e21866a9b825d4a794d169d7060a03605796360b
    timeout /t 1 > NUL
)


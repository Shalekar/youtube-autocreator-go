mkdir .\ffmpeg
curl -L https://www.gyan.dev/ffmpeg/builds/ffmpeg-release-essentials.zip > .\ffmpeg\ffmpeg-release-essentials.zip
tar -xf .\ffmpeg\ffmpeg-release-essentials.zip -C .\ffmpeg
del ".\ffmpeg\ffmpeg-release-essentials.zip"
FOR /F "tokens=*" %%g IN ('dir /b ffmpeg') do (SET firstFile=%%g)
move .\ffmpeg\%firstFile%\bin .\ffmpeg 
move .\ffmpeg\%firstFile%\doc .\ffmpeg 
move .\ffmpeg\%firstFile%\presets .\ffmpeg
del /S .\ffmpeg\%firstFile%\*
rmdir -f .\ffmpeg\%firstFile%
setx /M PATH "%PATH%;%CD%\ffmpeg\bin;"
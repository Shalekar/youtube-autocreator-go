New-Item '.\ffmpeg' -type directory
Invoke-WebRequest -Uri "https://www.gyan.dev/ffmpeg/builds/ffmpeg-release-essentials.zip" -OutFile ".\ffmpeg\ffmpeg-release-essentials.zip"
Expand-Archive .\ffmpeg\ffmpeg-release-essentials.zip -DestinationPath .\ffmpeg
Remove-Item .\ffmpeg\ffmpeg-release-essentials.zip
$dir = (Get-ChildItem -Path .\ffmpeg\ -Directory  | Select-Object -First 1)[0].Name
Get-Item -Path .\ffmpeg\$dir\* | Move-Item -Destination .\ffmpeg
Remove-Item .\ffmpeg\$dir
$newPath = [Environment]::GetEnvironmentVariable("PATH", "Machine") + [IO.Path]::PathSeparator + (Get-Item .).FullName + "\ffmpeg\bin" + [IO.Path]::PathSeparator
[Environment]::SetEnvironmentVariable( "Path", $newPath, "Machine" )
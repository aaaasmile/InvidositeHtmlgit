#This script is used to update the vido-preproc. It try to install the latest zip package into the.zips directory
# A backup of the ./ directory is done before into the ./old dir.
#version 0.0.1

$serviceName = "vido-preproc"

$dirZip = Resolve-Path ".\zips"
$dirCurrent = Resolve-Path "."
$dirOld = Resolve-Path ".\old"

$serviceCmd = "$dirCurrent\vido-preproc.exe"


#Get the latest zip file
$latestZipPkg = Get-ChildItem -Path $dirZip | Sort-Object Name -Descending | Select-Object -First 1

Write-Host "Lets start the update for $serviceCmd"

# create folder for the backup of the current installation
$TimeStamp = Get-Date -Format o | foreach {$_ -replace ":", "."}
$oldFolder = New-Item -ItemType Directory -Path "$dirOld\$TimeStamp"
Write-Host "Use backup folder $oldFolder"

#unzip the latest package
Write-Host "Unzip the package $latestZipPkg"
$destZipExpand = "$dirZip\$TimeStamp"

$confirmation = Read-Host "`nHello, this will update the $serviceName using the package `n$latestZipPkg`nZip expand is $destZipExpand dir current is  $dirCurrent`nContinue (y/n)?"
if ($confirmation -ne 'y') {
  Write-Host "As you like, nothing to do here. Bye."
  return
}

#Powershell version 5
Expand-Archive "$dirZip\$latestZipPkg" -DestinationPath $destZipExpand


#move the current version into the backup (CAUTION: move exe, static and template folder only)
if ([string]::IsNullOrEmpty($dirCurrent)) {
  Write-Host("Current dir is empty.`nSorry I can't continue.")
  return
}
if ([string]::IsNullOrEmpty($oldFolder)) {
  Write-Host("Old folder dir is empty.`nSorry I can't continue.")
  return
}
if ([string]::IsNullOrEmpty($destZipExpand)) {
  Write-Host("Expand dir folder dir is empty.`nSorry I can't continue.")
  return
}

Write-Host "Archive current into the folder $oldFolder"
Move-Item -Path "$dirCurrent\vido-preproc.exe" -Destination $oldFolder
Move-Item -Path "$dirCurrent\static*" -Destination $oldFolder
Move-Item -Path "$dirCurrent\templates*" -Destination $oldFolder
Copy-Item -Path "$dirCurrent\config.toml" -Destination $oldFolder
Move-Item -Path "$dirCurrent\config.toml" -Destination "$dirCurrent\config.toml_$TimeStamp"

#copy the zip content into the current
Write-Host "Copy the zip content into the dir $dirCurrent"
Move-Item -Path "$destZipExpand\*" -Destination $dirCurrent


Write-Host "That's all what I can do now.`nPlease check if the config.toml is the correct one."

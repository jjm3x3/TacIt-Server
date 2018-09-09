$Env:DB_PORT=33306
$Env:GOPATH=Split-Path $(Split-Path $(pwd))
$pwPath='.\kube\password.txt'
if(!(Test-Path $pwPath)){
	Write-Error "$pwPath file missing"
} else {
	$Env:DB_PASSWORD=$(Get-Content $pwPath)
}

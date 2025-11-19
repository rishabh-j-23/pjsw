function pjsw-project
{
    param(
        [string]$arg1
    )

    $dir = ""

    if (Test-Path $arg1)
    {
        Write-Host "switching using pjsw sw"
        $dir = & pjsw sw $arg1
    } else
    {
        $dir = & pjsw project
    }

    if ($dir)
    {
        Set-Location $dir
    }
}

param (
    [Parameter(Mandatory = $false)]
    [string]$Task = "build",
    [string]$Cmd = "dmrserver"
)

function Fmt {
    gofmt -s -w ./.
}

function Build {
    go build -o ./build/ "./cmd/$($Cmd)"
}

function Run {
    go run "./cmd/$($Cmd)" "-config=$($Config)"
}

switch ($Cmd) {
    Default { $Config = "./server.toml" }
}

switch ($Task) {
    "fmt" { Fmt }
    "build" { Build }
    "run" { Run }
    Default { Write-Host "Unknown task: $Task" }
}

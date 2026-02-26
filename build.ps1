param (
    [Parameter(Mandatory = $false)]
    [string]$Task = "build"
)

function Fmt {
    gofmt -s -w ./.
}

function Build {
    go build -o ./build/ ./cmd/dmrserver
}

function Run {
    go run ./cmd/dmrserver
}

switch ($Task) {
    "fmt" { Fmt }
    "build" { Build }
    "run" { Run }
    Default { Write-Host "Unknown task: $Task" }
}

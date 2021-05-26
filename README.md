# purple
purple is a container-optimized tool for continuous security testing. purple can execute Atomic Red Team definitions.

## Usage
`./purplecli test -f <path to test spec>`

Example:
```bigquery
./purplecli test -f examples/simple.yaml
INFO[0000] Running test suite 'Uname Example'
INFO[0000] Test: UnameExec
INFO[0000] Executing uname -a
Darwin siyeh.local 20.4.0 Darwin Kernel Version 20.4.0: Thu Apr 22 21:46:47 PDT 2021; root:xnu-7195.101.2~1/RELEASE_X86_64 x86_64
```
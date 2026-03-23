# Doctor CLI - Environment Checker

Doctor is a CLI tool to check and fix your developer environment. It validates tools, ports, and dependencies to ensure your setup is ready.

Usage:
./doctor                     # interactive mode
./doctor diagnose            # run full environment checks
./doctor diagnose --json     # run checks and output JSON
./doctor check <tool>        # check a single tool (docker, git, go, node, python3, sleep)
./doctor check <tool> --version  # show version of a specific tool
./doctor check <tool> --path     # show path of a specific tool
./doctor fix                 # fix all actionable issues
./doctor fix docker          # fix Docker only
./doctor port <port>         # check a specific port
./doctor list                # list all available checks
./doctor version             # show Doctor version
./doctor help                # show this help

Available Checks:
docker
git
go
node
python3
sleep

Notes:
- Docker must be running for Docker checks/fixes.
- Ports 3000 and 5432 are checked by default.
- Interactive mode prompts you to fix issues automatically.
- Auto-fix mode runs diagnose and applies fixes sequentially.

Example workflow:
1. ./doctor                  # see interactive overview
2. ./doctor diagnose          # see status of all tools
3. ./doctor fix               # fix all problems automatically
4. ./doctor list              # list available checks
5. ./doctor port 3000         # check if a port is free

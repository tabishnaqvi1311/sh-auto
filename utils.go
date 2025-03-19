package main

import (
	"fmt"
	"os"
	"os/exec"
)

const SystemScript = `
START_TIME=$(date +%%s)
echo "compiling"
%s
EXIT_CODE=$?
echo "Build complete."
END_TIME=$(date +%%s)
ELAPSED_TIME=$((END_TIME - START_TIME))
GREEN="\033[0;32m"
RED="\033[0;31m"
NC="\033[0m" # No Color
if [ $EXIT_CODE -eq 0 ]; then
    echo "${GREEN}✔ Compiled in ${ELAPSED_TIME} seconds${NC}"
else
    echo "${RED}✘ Compilation failed with exit code ${EXIT_CODE} in ${ELAPSED_TIME} seconds${NC}"
fi
exit $EXIT_CODE
`

func WriteToSh(buildCmd string) error {
    if buildCmd == "" {
        return fmt.Errorf("build command missing")
    }

    fmt.Println("Baking script... 🍞")

    fullScript := fmt.Sprintf(SystemScript, buildCmd)

    err := os.WriteFile("build.sh", []byte(fullScript), 0755)
    if err != nil {
        return err
    }

    fmt.Println("Running build.sh... 🚀")

    cmd := exec.Command("sh", "build.sh")
    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        return err
    }

    return nil
}

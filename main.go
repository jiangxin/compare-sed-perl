package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

var (
	input = []string{
		"beginning\n",
		"commit c014567890123456789012345678901234567890\n",
		"commit c004567890123456789012345678901234567890\n",
		"commit c024567890123456789012345678901234567890\n",
		"commit c034567890123456789012345678901234567890\n",
		"commit c044567890123456789012345678901234567890\n",
		"commit c054567890123456789012345678901234567890\n",
		"commit c064567890123456789012345678901234567890\n",
		"commit c074567890123456789012345678901234567890\n",
		"commit c084567890123456789012345678901234567890\n",
		"commit c094567890123456789012345678901234567890\n",
		"commit c0a4567890123456789012345678901234567890\n",
		"commit c0b4567890123456789012345678901234567890\n",
		"commit c0c4567890123456789012345678901234567890\n",
		"commit c0d4567890123456789012345678901234567890\n",
		"commit c0e4567890123456789012345678901234567890\n",
		"tag 7a91567890123456789012345678901234567890\n",
		"tag 7a92567890123456789012345678901234567890\n",
		"tag 7a93567890123456789012345678901234567890\n",
		"commit c0f4567890123456789012345678901234567890\n",
		"error: this error message should be striped.\n",
		"end.\n",
	}

	expect = []byte(`beginning
commit <COMMIT-1>
commit <COMMIT-0>
commit <COMMIT-2>
commit <COMMIT-3>
commit <COMMIT-4>
commit <COMMIT-5>
commit <COMMIT-6>
commit <COMMIT-7>
commit <COMMIT-8>
commit <COMMIT-9>
commit <COMMIT-A>
commit <COMMIT-B>
commit <COMMIT-C>
commit <COMMIT-D>
commit <COMMIT-E>
tag <TAG-1>
tag <TAG-2>
tag <TAG-3>
commit <COMMIT-F>
end.
`)
)

func GetCmdResUsage(args ...string) (int64, error) {
	cmd := exec.Command(args[0], args[1:]...)
	in, err := cmd.StdinPipe()
	if err != nil {
		return 0, err
	}
	go func() {
		defer in.Close()
		for _, line := range input {
			io.WriteString(in, line)
		}
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}
	if bytes.Compare(expect, out) != 0 {
		fmt.Fprintf(os.Stderr, "ERROR: unexpected output.\n")
		fmt.Fprintf(os.Stderr, "expect: %s\n", expect)
		fmt.Fprintf(os.Stderr, "actual: %s\n", out)
		fmt.Fprintln(os.Stderr, "")
	}
	if sysUsage, ok := cmd.ProcessState.SysUsage().(*syscall.Rusage); ok {
		return sysUsage.Maxrss, nil
	}
	return 0, errors.New("fail to get sysusage")
}

func main() {
	cmds := [][]string{
		[]string{"sed",
			"-e", "s/  *$//",
			"-e", "s/  */ /g",
			"-e", "s/'/\"/g",
			"-e", "s/	/    /g",
			"-e", "s/c004567[0-9a-f]*/<COMMIT-0>/g",
			"-e", "s/c014567[0-9a-f]*/<COMMIT-1>/g",
			"-e", "s/c024567[0-9a-f]*/<COMMIT-2>/g",
			"-e", "s/c034567[0-9a-f]*/<COMMIT-3>/g",
			"-e", "s/c044567[0-9a-f]*/<COMMIT-4>/g",
			"-e", "s/c054567[0-9a-f]*/<COMMIT-5>/g",
			"-e", "s/c064567[0-9a-f]*/<COMMIT-6>/g",
			"-e", "s/c074567[0-9a-f]*/<COMMIT-7>/g",
			"-e", "s/c084567[0-9a-f]*/<COMMIT-8>/g",
			"-e", "s/c094567[0-9a-f]*/<COMMIT-9>/g",
			"-e", "s/c0a4567[0-9a-f]*/<COMMIT-A>/g",
			"-e", "s/c0b4567[0-9a-f]*/<COMMIT-B>/g",
			"-e", "s/c0c4567[0-9a-f]*/<COMMIT-C>/g",
			"-e", "s/c0d4567[0-9a-f]*/<COMMIT-D>/g",
			"-e", "s/c0e4567[0-9a-f]*/<COMMIT-E>/g",
			"-e", "s/c0f4567[0-9a-f]*/<COMMIT-F>/g",
			"-e", "s/7a91567[0-9a-f]*/<TAG-1>/g",
			"-e", "s/7a92567[0-9a-f]*/<TAG-2>/g",
			"-e", "s/7a93567[0-9a-f]*/<TAG-3>/g",
			"-e", "/^error: / d",
		},

		// GNU sed
		[]string{"gsed",
			"-e", "s/  *$//",
			"-e", "s/  */ /g",
			"-e", "s/'/\"/g",
			"-e", "s/	/    /g",
			"-e", "s/c004567[0-9a-f]*/<COMMIT-0>/g",
			"-e", "s/c014567[0-9a-f]*/<COMMIT-1>/g",
			"-e", "s/c024567[0-9a-f]*/<COMMIT-2>/g",
			"-e", "s/c034567[0-9a-f]*/<COMMIT-3>/g",
			"-e", "s/c044567[0-9a-f]*/<COMMIT-4>/g",
			"-e", "s/c054567[0-9a-f]*/<COMMIT-5>/g",
			"-e", "s/c064567[0-9a-f]*/<COMMIT-6>/g",
			"-e", "s/c074567[0-9a-f]*/<COMMIT-7>/g",
			"-e", "s/c084567[0-9a-f]*/<COMMIT-8>/g",
			"-e", "s/c094567[0-9a-f]*/<COMMIT-9>/g",
			"-e", "s/c0a4567[0-9a-f]*/<COMMIT-A>/g",
			"-e", "s/c0b4567[0-9a-f]*/<COMMIT-B>/g",
			"-e", "s/c0c4567[0-9a-f]*/<COMMIT-C>/g",
			"-e", "s/c0d4567[0-9a-f]*/<COMMIT-D>/g",
			"-e", "s/c0e4567[0-9a-f]*/<COMMIT-E>/g",
			"-e", "s/c0f4567[0-9a-f]*/<COMMIT-F>/g",
			"-e", "s/7a91567[0-9a-f]*/<TAG-1>/g",
			"-e", "s/7a92567[0-9a-f]*/<TAG-2>/g",
			"-e", "s/7a93567[0-9a-f]*/<TAG-3>/g",
			"-e", "/^error: / d",
		},

		[]string{"perl", "-ne",
			`s/  *$//;
			s/  */ /g;
			s/'/"/g;
			s/	/    /g;
			s/c004567[0-9a-f]*/<COMMIT-0>/g;
			s/c014567[0-9a-f]*/<COMMIT-1>/g;
			s/c024567[0-9a-f]*/<COMMIT-2>/g;
			s/c034567[0-9a-f]*/<COMMIT-3>/g;
			s/c044567[0-9a-f]*/<COMMIT-4>/g;
			s/c054567[0-9a-f]*/<COMMIT-5>/g;
			s/c064567[0-9a-f]*/<COMMIT-6>/g;
			s/c074567[0-9a-f]*/<COMMIT-7>/g;
			s/c084567[0-9a-f]*/<COMMIT-8>/g;
			s/c094567[0-9a-f]*/<COMMIT-9>/g;
			s/c0a4567[0-9a-f]*/<COMMIT-A>/g;
			s/c0b4567[0-9a-f]*/<COMMIT-B>/g;
			s/c0c4567[0-9a-f]*/<COMMIT-C>/g;
			s/c0d4567[0-9a-f]*/<COMMIT-D>/g;
			s/c0e4567[0-9a-f]*/<COMMIT-E>/g;
			s/c0f4567[0-9a-f]*/<COMMIT-F>/g;
			s/7a91567[0-9a-f]*/<TAG-1>/g;
			s/7a92567[0-9a-f]*/<TAG-2>/g;
			s/7a93567[0-9a-f]*/<TAG-3>/g;
			next if /^error: .*$/;
			print`,
		},
	}

	for _, cmd := range cmds {
		res, err := GetCmdResUsage(cmd...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: fail to run %s...: %s\n", cmd[0], err)
			continue
		}
		fmt.Printf("Command: %-5s..., ", cmd[0])
		fmt.Printf("MaxRSS: %d\n", res)
	}
}

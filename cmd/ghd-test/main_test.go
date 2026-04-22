package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	originalVersion, originalCommit, originalDate := version, commit, date
	version, commit, date = "v1.2.3", "abc123", "2026-04-22T13:00:00Z"
	t.Cleanup(func() {
		version, commit, date = originalVersion, originalCommit, originalDate
	})

	tests := []struct {
		name       string
		args       []string
		wantCode   int
		wantStdout string
		wantStderr string
	}{
		{
			name:       "default command prints hello world",
			args:       nil,
			wantCode:   0,
			wantStdout: "hello from ghd-test\n",
		},
		{
			name:       "version subcommand prints build metadata",
			args:       []string{"version"},
			wantCode:   0,
			wantStdout: "ghd-test v1.2.3 (abc123) built 2026-04-22T13:00:00Z\n",
		},
		{
			name:       "version flag prints build metadata",
			args:       []string{"--version"},
			wantCode:   0,
			wantStdout: "ghd-test v1.2.3 (abc123) built 2026-04-22T13:00:00Z\n",
		},
		{
			name:       "unknown command fails",
			args:       []string{"nope"},
			wantCode:   2,
			wantStderr: "unknown command \"nope\"\nUsage: ghd-test [version|--version]\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout strings.Builder
			var stderr strings.Builder

			gotCode := run(tt.args, &stdout, &stderr)

			if gotCode != tt.wantCode {
				t.Fatalf("exit code = %d, want %d", gotCode, tt.wantCode)
			}
			if stdout.String() != tt.wantStdout {
				t.Fatalf("stdout = %q, want %q", stdout.String(), tt.wantStdout)
			}
			if stderr.String() != tt.wantStderr {
				t.Fatalf("stderr = %q, want %q", stderr.String(), tt.wantStderr)
			}
		})
	}
}

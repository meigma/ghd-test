# Contributing

`ghd-test` is intentionally small. Changes should preserve its role as a stable
release fixture for `ghd` functional tests.

## Local Checks

Run these before opening a pull request:

```sh
go test ./...
go build ./cmd/ghd-test
goreleaser check
```

Use a snapshot build when changing release packaging:

```sh
goreleaser release --snapshot --clean --skip=sign,sbom
```

The full release workflow signs checksums, generates SBOMs, and creates GitHub
artifact attestations on GitHub-hosted runners.

## Pull Requests

- Keep changes focused.
- Use Conventional Commit titles, such as `feat: add release fixture CLI`.
- Update `ghd.toml` when release asset names or binary paths change.
- Keep GitHub Actions pinned to full commit SHAs.

## Releases

Release Please owns version PRs and tags. Do not create manual release tags
unless recovering from broken automation.

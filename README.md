# ghd-test

`ghd-test` is a small Go CLI used as a release fixture for
[GitHub Downloader](https://github.com/meigma/ghd). It exists so `ghd` can run
functional tests against real GitHub releases, release assets, immutable release
attestations, and GitHub artifact attestations.

## Usage

```sh
go run ./cmd/ghd-test
go run ./cmd/ghd-test version
```

Expected output:

```text
hello from ghd-test
```

Release binaries print linker-injected version metadata:

```text
ghd-test v0.1.0 (<commit>) built <date>
```

## `ghd` Fixture Contract

The root `ghd.toml` describes one package named `ghd-test`. The package uses
`v${version}` tags, GoReleaser archive assets, and the release workflow
identity `meigma/ghd-test/.github/workflows/release.yml` for provenance
verification.

The release pipeline is intentionally simple:

1. Release Please opens and maintains the release PR.
2. Merging the release PR creates a tag and draft GitHub release.
3. The tag-triggered Release workflow builds GoReleaser archives, signs
   `checksums.txt`, generates SBOMs, creates GitHub artifact attestations, and
   publishes the draft release.

## Local Development

```sh
go test ./...
go build ./cmd/ghd-test
goreleaser check
goreleaser release --snapshot --clean --skip=sign,sbom
```

The release workflow installs Syft and uses GitHub OIDC for keyless Cosign
signing. Local snapshot builds may need `--skip=sign,sbom` unless the same
tools and identity flow are available on the developer machine.

## Release Verification

After a release is published:

```sh
gh release verify v0.1.0 -R meigma/ghd-test
gh release verify-asset v0.1.0 ./ghd-test_0.1.0_linux_amd64.tar.gz -R meigma/ghd-test
gh attestation verify ./ghd-test_0.1.0_linux_amd64.tar.gz \
  --repo meigma/ghd-test \
  --signer-workflow meigma/ghd-test/.github/workflows/release.yml \
  --source-ref refs/tags/v0.1.0 \
  --deny-self-hosted-runners
```

## Security

Do not report vulnerabilities in public issues or discussions. See
[SECURITY.md](SECURITY.md).

# Migration Guide: Serenity-Go -> Verity-BDD

This release introduces a hard rename from `Serenity-Go` to `Verity-BDD`.

There are no compatibility aliases for old names.

## What Changed

- Module path changed: `github.com/nchursin/serenity-go` -> `github.com/nchursin/verity-bdd`
- Root package directory changed: `serenity/` -> `verity/`
- Public testing API renamed from `Serenity*` to `Verity*`

## Import Path Migration

Before:

```go
import (
    "github.com/nchursin/serenity-go/serenity/abilities/api"
    serenity "github.com/nchursin/serenity-go/serenity/testing"
)
```

After:

```go
import (
    "github.com/nchursin/verity-bdd/verity_abilities/api"
    verity "github.com/nchursin/verity-bdd"
)
```

## Public API Rename Map

- `SerenityTest` -> `VerityTest`
- `NewSerenityTest` -> `NewVerityTest`
- `NewSerenityTestWithContext` -> `NewVerityTestWithContext`
- `NewSerenityTestWithReporter` -> `NewVerityTestWithReporter`

## Example Migration

Before:

```go
test := serenity.NewSerenityTest(t, serenity.Scene{})
```

After:

```go
test := verity.NewVerityTest(t, verity.Scene{})
```

## Suggested Migration Steps

1. Update `go.mod` dependencies to use `github.com/nchursin/verity-bdd`.
2. Replace import paths from `/serenity/...` to `/verity/...`.
3. Rename API symbols from `Serenity*` to `Verity*`.
4. Run formatting and tests:

```bash
go mod tidy
go test ./...
```

## Breaking Change Notice

This is a breaking rename release. Existing projects using Serenity-Go names will not compile until migrated.

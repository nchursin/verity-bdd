# Verity-BDD Rename Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Выполнить жесткий переход проекта с Verity-BDD на Verity-BDD без обратной совместимости по старым именам API.

**Architecture:** Миграция делится на 4 слоя: (1) safety tests для нового API-нейминга, (2) rename модуля и import path, (3) rename директории/пакета и публичных символов, (4) инфраструктура CI/release/docs. Каждый шаг идет по TDD: один RED тест, минимальный GREEN, рефакторинг, повтор.

**Tech Stack:** Go modules, go test, golangci-lint, GitHub Actions, git-cliff, semantic-release.

---

### Task 1: Зафиксировать fail по новому публичному API

**Files:**
- Modify: `verity/testing/serenity_test_manager_test.go`
- Test: `verity/testing/serenity_test_manager_test.go`

**Step 1: Write the failing test**

Добавить один тест, который использует только новые имена API (`VerityTest`, `NewVerityTest`, `NewVerityTestWithContext`, `NewVerityTestWithReporter`) и не использует `Serenity*`.

**Step 2: Run test to verify it fails**

Run: `go test ./verity/testing -run TestNewVerityNames -v`
Expected: FAIL with undefined identifiers for `Verity*`.

**Step 3: Write minimal implementation**

В `verity/testing/serenity_test_manager.go` добавить минимальные определения новых имен (без удаления старых на этом шаге).

**Step 4: Run test to verify it passes**

Run: `go test ./verity/testing -run TestNewVerityNames -v`
Expected: PASS.

**Step 5: Commit**

```bash
git add verity/testing/serenity_test_manager_test.go verity/testing/serenity_test_manager.go
git commit -m "test(rename): add failing coverage for Verity API names"
```

### Task 2: Переименовать модуль в go.mod

**Files:**
- Modify: `go.mod`
- Test: `examples/basic_test_new_api_test.go`

**Step 1: Write the failing test**

Добавить один compile-oriented тест/пример импорта с `github.com/nchursin/verity-bdd/...`.

**Step 2: Run test to verify it fails**

Run: `go test ./examples -run TestImportPathVerityBDD -v`
Expected: FAIL with module/import path mismatch.

**Step 3: Write minimal implementation**

Изменить строку модуля в `go.mod`:

```go
module github.com/nchursin/verity-bdd
```

**Step 4: Run test to verify it passes**

Run: `go test ./examples -run TestImportPathVerityBDD -v`
Expected: PASS or compile success for this scope.

**Step 5: Commit**

```bash
git add go.mod examples/basic_test_new_api_test.go
git commit -m "feat(rename)!: switch module path to github.com/nchursin/verity-bdd"
```

### Task 3: Массово обновить internal import path

**Files:**
- Modify: `serenity/**/*.go`
- Modify: `examples/**/*.go`
- Modify: `README.md`
- Modify: `docs/**/*.md`
- Test: `verity/testing/integration_test.go`

**Step 1: Write the failing test**

Добавить один тест в `verity/testing/integration_test.go`, где используется новый import path и новые символы API.

**Step 2: Run test to verify it fails**

Run: `go test ./verity/testing -run TestIntegrationUsesVerityImportPath -v`
Expected: FAIL while old imports still present.

**Step 3: Write minimal implementation**

Сделать механическую замену:
- `github.com/nchursin/verity-bdd` -> `github.com/nchursin/verity-bdd`

**Step 4: Run test to verify it passes**

Run: `go test ./verity/testing -run TestIntegrationUsesVerityImportPath -v`
Expected: PASS.

**Step 5: Commit**

```bash
git add serenity examples README.md docs
git commit -m "refactor(rename): update internal imports to verity-bdd module path"
```

### Task 4: Переименовать корневую директорию пакетов serenity -> verity

**Files:**
- Move: `serenity/` -> `verity/`
- Modify: `.github/workflows/ci.yml`
- Test: `verity/doc.go`

**Step 1: Write the failing test**

Добавить тест, который запускает пакетный импорт только по `./verity/...`.

**Step 2: Run test to verify it fails**

Run: `go test ./verity/...`
Expected: FAIL because directory/package path still old.

**Step 3: Write minimal implementation**

Переименовать директорию `serenity` в `verity`, затем поправить все пути в import и CI (`./verity/...` -> `./verity/...`).

**Step 4: Run test to verify it passes**

Run: `go test ./verity/...`
Expected: PASS.

**Step 5: Commit**

```bash
git add verity .github/workflows/ci.yml
git commit -m "refactor(rename)!: move root package directory from serenity to verity"
```

### Task 5: Переименовать package identity и типы API без backward compatibility

**Files:**
- Modify: `verity/doc.go`
- Modify: `verity/testing/serenity_test_manager.go`
- Modify: `verity/testing/serenity_test_manager_test.go`
- Test: `verity/testing/serenity_test_manager_test.go`

**Step 1: Write the failing test**

Добавить один тест, который проверяет отсутствие `VerityTest` и работоспособность только `VerityTest` API.

**Step 2: Run test to verify it fails**

Run: `go test ./verity/testing -run TestNoSerenityAPINamesRemain -v`
Expected: FAIL while old symbols still exported.

**Step 3: Write minimal implementation**

Переименовать:
- `package serenity` -> `package verity` в `verity/doc.go`
- `type VerityTest` -> `type VerityTest`
- `NewVerityTest*` -> `NewVerityTest*`
- удалить экспорт старых имен (без alias и deprecations).

**Step 4: Run test to verify it passes**

Run: `go test ./verity/testing -run TestNoSerenityAPINamesRemain -v`
Expected: PASS.

**Step 5: Commit**

```bash
git add verity/testing/serenity_test_manager.go verity/testing/serenity_test_manager_test.go verity/doc.go
git commit -m "feat(rename)!: rename public testing API from Serenity* to Verity*"
```

### Task 6: Обновить CI, линтер и release-конфигурацию под новый бренд

**Files:**
- Modify: `.github/workflows/ci.yml`
- Modify: `.github/workflows/release.yml`
- Modify: `.golangci.yml`
- Modify: `.semrelrc`
- Modify: `cliff-github.toml`
- Modify: `scripts/release.sh`
- Modify: `docs/RELEASING.md`

**Step 1: Write the failing test**

Добавить script-level проверку (grep-based) в локальный verify command: CI/release файлы не содержат `serenity-go`.

**Step 2: Run test to verify it fails**

Run: `rg -n "serenity-go|Verity-BDD|serenity-coverage|test-serenity" .github scripts .semrelrc .golangci.yml cliff-github.toml docs/RELEASING.md`
Expected: FAIL (найдены старые строки).

**Step 3: Write minimal implementation**

Обновить все хардкоды под `verity-bdd` и `verity-*` naming.

**Step 4: Run test to verify it passes**

Run: `rg -n "serenity-go|Verity-BDD|serenity-coverage|test-serenity" .github scripts .semrelrc .golangci.yml cliff-github.toml docs/RELEASING.md`
Expected: no matches.

**Step 5: Commit**

```bash
git add .github/workflows/ci.yml .github/workflows/release.yml .golangci.yml .semrelrc cliff-github.toml scripts/release.sh docs/RELEASING.md
git commit -m "chore(rename): align ci and release tooling with verity-bdd naming"
```

### Task 7: Полностью обновить README/docs/examples под Verity-BDD

**Files:**
- Modify: `README.md`
- Modify: `docs/index.md`
- Modify: `docs/abilities.md`
- Modify: `docs/reporting.md`
- Modify: `docs/SATISFIES_EXAMPLES.md`
- Modify: `docs/examples/**/*.md`
- Modify: `docs/templates/ability.md`
- Modify: `examples/**/*.go`
- Modify: `examples/**/*.md`

**Step 1: Write the failing test**

Добавить docs-check команду в local verification: в публичной документации не осталось `Verity-BDD`.

**Step 2: Run test to verify it fails**

Run: `rg -n "Verity-BDD|github.com/nchursin/verity-bdd|serenity/" README.md docs examples`
Expected: FAIL with matches.

**Step 3: Write minimal implementation**

Обновить бренд, install snippets, import snippets, package aliases, разделы сравнения и примеры API.

**Step 4: Run test to verify it passes**

Run: `rg -n "Verity-BDD|github.com/nchursin/verity-bdd|serenity/" README.md docs examples`
Expected: no matches (кроме отдельного migration файла).

**Step 5: Commit**

```bash
git add README.md docs examples
git commit -m "docs(rename): migrate all docs and examples to Verity-BDD"
```

### Task 8: Добавить migration guide для пользователей

**Files:**
- Create: `docs/MIGRATION_SERENITY_TO_VERITY.md`

**Step 1: Write the failing test**

Добавить ссылку на migration doc из `README.md`; пока файла нет, линк невалиден (проверяется вручную).

**Step 2: Run test to verify it fails**

Run: `test -f docs/MIGRATION_SERENITY_TO_VERITY.md`
Expected: FAIL.

**Step 3: Write minimal implementation**

Создать `docs/MIGRATION_SERENITY_TO_VERITY.md` с:
- before/after imports,
- таблицей rename символов,
- шагами массовой замены,
- явным предупреждением о breaking change.

**Step 4: Run test to verify it passes**

Run: `test -f docs/MIGRATION_SERENITY_TO_VERITY.md`
Expected: PASS.

**Step 5: Commit**

```bash
git add docs/MIGRATION_SERENITY_TO_VERITY.md README.md
git commit -m "docs!: add migration guide from Verity-BDD to Verity-BDD"
```

### Task 9: Финальная полная верификация

**Files:**
- Modify: none
- Test: full repo

**Step 1: Write the failing test**

Не требуется (verification task).

**Step 2: Run test to verify it fails**

Run: `rg -n "Verity-BDD|github.com/nchursin/verity-bdd|NewVerityTest|VerityTest|/verity/" .`
Expected: matches only in `docs/MIGRATION_SERENITY_TO_VERITY.md`.

**Step 3: Write minimal implementation**

Если есть лишние совпадения вне migration doc — убрать.

**Step 4: Run test to verify it passes**

Run:
- `go mod tidy`
- `go test ./... -race`
- `golangci-lint run`

Expected: all PASS.

**Step 5: Commit**

```bash
git add .
git commit -m "chore(rename)!: complete hard transition from Verity-BDD to Verity-BDD"
```

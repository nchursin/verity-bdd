# Console Reporting

Serenity-Go предоставляет мощную систему консольного репортинга для визуализации результатов тестирования в реальном времени.

## Overview

ConsoleReporter автоматически отображает информацию о выполнении тестов, включая:
- Статусы тестов с emoji индикаторами
- Время выполнения
- Детали ошибок и проваленных ожиданий
- Возможность записи вывода в файл

## Basic Usage

### Автоматическая интеграция с TestContext

```go
func TestAPITesting(t *testing.T) {
    test := serenity.NewSerenityTest(t, serenity.Scene{})

    actor := test.ActorCalled("APITester").WhoCan(api.CallAnApiAt("https://jsonplaceholder.typicode.com"))

    actor.AttemptsTo(
        api.SendGetRequest("/posts"),
        ensure.That(api.LastResponseStatus{}, expectations.Equals(200)),
    )
}
```

Вывод в консоль:
```
🚀 Starting: TestAPITesting
  🔄 Sends GET request to /posts
  ✅ Sends GET request to /posts (0.21s)
  🔄 Ensures that the last response status code equals 200
  ✅ Ensures that the last response status code equals 200 (0.00s)
✅ TestAPITesting: PASSED (0.26s)
```

### Ручная настройка репортера

```go
import (
    "github.com/nchursin/serenity-go/serenity/reporting/console_reporter"
    serenity "github.com/nchursin/serenity-go/serenity/testing"
)

func TestCustomReporting(t *testing.T) {
    reporter := console_reporter.NewConsoleReporter()

    test := serenity.NewSerenityTestWithReporter(t, reporter)

    // ... тестовый код
}
```

## Custom Reporter Configuration

### Настройка вывода в файл

```go
import (
    "os"
    "github.com/nchursin/serenity-go/serenity/reporting/console_reporter"
    serenity "github.com/nchursin/serenity-go/serenity/testing"
)

reporter := console_reporter.NewConsoleReporter()

// Создаем файл для вывода
file, err := os.Create("test-results.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// Настраиваем репортер на запись в файл
reporter.SetOutput(file)

test := serenity.NewSerenityTestWithReporter(t, reporter)

// ... тестовый код
```

### Методы управления

```go
import (
    "os"
    "github.com/nchursin/serenity-go/serenity/reporting/console_reporter"
)

reporter := console_reporter.NewConsoleReporter()

// Установка вывода (файл или консоль)
reporter.SetOutput(os.Stdout)  // Консольный вывод
reporter.SetOutput(file)      // Вывод в файл
```

## File Output

ConsoleReporter может записывать вывод в файл для последующего анализа:

```go
import (
    "os"
    "github.com/nchursin/serenity-go/serenity/reporting/console_reporter"
    serenity "github.com/nchursin/serenity/testing"
)

// Создаем файл для вывода
file, err := os.Create("test-results.txt")
if err != nil {
    t.Fatalf("Failed to create output file: %v", err)
}
defer file.Close()

// Создаем репортер с выводом в файл
reporter := console_reporter.NewConsoleReporter()
reporter.SetOutput(file)

test := serenity.NewSerenityTestWithReporter(t, reporter)

// ... тестовый код
```

Файл будет содержать полный вывод тестов в том же формате, что и консоль.

## Output Format

### Статусы тестов

| Статус | Emoji | Описание |
|--------|-------|----------|
| ✅ | ✅ | Тест успешно пройден |
| ❌ | ❌ | Тест провален |
| ⚠️ | ⚠️ | Предупреждение (неиспользованный actor) |

### Формат вывода

```
✅ TestName (duration)
❌ TestName (duration)
   Error: error message
   Stack trace: stack information
⚠️ TestName (duration)
   Warning: warning message
```

### Пример полного вывода

```
🚀 Starting: TestAPITesting
  🔄 Sends GET request to /posts
  ✅ Sends GET request to /posts (0.21s)
  🔄 Ensures that the last response status code equals 200
  ✅ Ensures that the last response status code equals 200 (0.00s)
✅ TestAPITesting: PASSED (0.26s)

🚀 Starting: TestFailedExpectation
  🔄 Sends GET request to /posts
  ❌ Sends GET request to /posts (0.15s)
     Error: Expected status code to equal 200, but got 404
❌ TestFailedExpectation: FAILED (0.15s)
```

## Integration Information

### Совместимость с TestContext API

ConsoleReporter автоматически интегрирован с TestContext API:

```go
test := serenity.NewSerenityTest(t, serenity.Scene{})  // Автоматически использует ConsoleReporter
```

### Интеграция с SerenityTest

```go
import (
    "github.com/nchursin/serenity-go/serenity/reporting/console_reporter"
    serenity "github.com/nchursin/serenity-go/serenity/testing"
)

test := serenity.NewSerenityTestWithReporter(t, customReporter)
```

### Обработка ошибок

Репортер автоматически:
- Логирует ошибки записи в файл
- Обрабатывает проблемы с доступом к файловой системе
- Предоставляет информативные сообщения об ошибках

### Потокобезопасность

ConsoleReporter потокобезопасен и может использоваться в параллельных тестах. Каждая тестовая сессия создает изолированный репортер.

## Migration from Legacy Testing

### Старый подход (ручная обработка ошибок)

```go
func TestOldStyle(t *testing.T) {
    test := serenity.NewSerenityTest(t, serenity.Scene{})

    actor := test.ActorCalled("Tester").WhoCan(api.CallAnApiAt("https://api.example.com"))

    err := actor.AttemptsTo(
        api.SendGetRequest("/users"),
        ensure.That(api.LastResponseStatus{}, expectations.Equals(200)),
    )
    if err != nil {
        t.Errorf("Test failed: %v", err)
    }
}
```

### Новый подход (автоматический репортинг)

```go
func TestNewStyle(t *testing.T) {
    test := serenity.NewSerenityTest(t, serenity.Scene{})

    actor := test.ActorCalled("Tester").WhoCan(api.CallAnApiAt("https://api.example.com"))

    actor.AttemptsTo(
        api.SendGetRequest("/users"),
        ensure.That(api.LastResponseStatus{}, expectations.Equals(200)),
    )
    // Статус и ошибки автоматически отображаются в консоли
}
```

## Best Practices

1. **Используйте TestContext API** для автоматического репортинга
2. **Настраивайте файловый вывод** для CI/CD пайплайнов
3. **Используйте descripting имена** для акторов для лучшей читаемости
5. **Настройте quiet mode** для CI сред, где важен только файловый вывод

## Troubleshooting

### Файл не создается

```go
err := reporter.EnableFileOutput("results.txt")
if err != nil {
    // Проверьте права доступа и директорию
    log.Printf("Failed to create file: %v", err)
}
```

### Нет вывода в консоли

Убедитесь, что репортер настроен на вывод в консоль:
```go
import (
    "os"
    "github.com/nchursin/serenity-go/serenity/reporting/console_reporter"
)

reporter := console_reporter.NewConsoleReporter()
reporter.SetOutput(os.Stdout)  // Явный вывод в консоль
```

### Проблемы с параллельными тестами

Каждый тест должен создавать собственный TestContext:
```go
import (
    serenity "github.com/nchursin/serenity-go/serenity/testing"
)

func TestParallel1(t *testing.T) {
    test := serenity.NewSerenityTest(t, serenity.Scene{})
    // ... тестовый код
}
```

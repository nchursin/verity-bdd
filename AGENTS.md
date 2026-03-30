# AGENTS.md - Guide for AI Coding Agents

This document provides essential information for AI agents working with the Verity-BDD codebase.

## Screenplay Pattern
This lib is implementing a [Screenplay pattern](https://serenity-js.org/handbook/design/screenplay-pattern/) for Go.

Core objects are in the `./verity` folder.

## Development workflow
For all functionlaity **ALWAYS** use TDD approach.

1. Create a set of test cases as a plan
2. Go with the TDD loop
3. **DO NOT WRITE** new functionality without a **RED FAILING TEST**

### TDD loop
1. Write **exactly one** **RED** failing test
2. Write code to make it **GREEN**
3. Refactor. Do not add new functionality, **refactor only**
4. Make a commit with prefix "TDD WIP:"
5. Go to 1.

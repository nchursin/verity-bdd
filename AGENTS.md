# AGENTS.md - Guide for AI Coding Agents

This document provides essential information for AI agents working with the Serenity-Go codebase.

## Screenplay Pattern
This lib is implementing a [Screenplay pattern](https://serenity-js.org/handbook/design/screenplay-pattern/) for Go.

Core objects are in the `./serenity` folder.

## Development workflow
For all functionlaity **ALWAYS** use TDD approach.

1. Create a set of test cases as a plan
2. Go with the TDD loop

### TDD loop
1. Write one **RED** failing test
2. Write code to make it **GREEN**
3. Refactor. Do not add new functionality, **refactor only**
4. Make a commit with prefix "TDD WIP:"
5. Go to 1.

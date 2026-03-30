#!/bin/bash

# Verity-BDD Release Script
# Automates the release process with proper validation

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🚀 Verity-BDD Release Script${NC}"
echo "================================"

# Check if we're in the right directory
if [[ ! -f "go.mod" || ! -f "Makefile" ]]; then
    echo -e "${RED}❌ Error: Must be run from project root directory${NC}"
    exit 1
fi

# Check required tools
TOOLS=("git" "go")
if [[ ! -f "$HOME/bin/git-cliff" ]]; then
    echo -e "${YELLOW}⚠️  git-cliff not found in ~/bin/ - some features may not work${NC}"
fi

for tool in "${TOOLS[@]}"; do
    if ! command -v "$tool" &> /dev/null; then
        echo -e "${RED}❌ Error: Required tool '$tool' not found${NC}"
        exit 1
    fi
done

# Ensure we're on main/master branch
current_branch=$(git --no-pager branch --show-current)
if [[ "$current_branch" != "main" && "$current_branch" != "master" ]]; then
    echo -e "${RED}❌ Must be on main or master branch. Current: $current_branch${NC}"
    exit 1
fi

# Check if working directory is clean
if [[ -n $(git --no-pager status --porcelain) ]]; then
    echo -e "${RED}❌ Working directory not clean. Commit or stash changes.${NC}"
    git --no-pager status --short
    exit 1
fi

# Function to check if git-cliff is available
check_git_cliff() {
    if [[ ! -f "$HOME/bin/git-cliff" ]]; then
        echo -e "${YELLOW}⚠️  git-cliff not found. Using basic git --no-pager log for changelog.${NC}"
        return 1
    fi
    return 0
}

# Function to show preview
show_preview() {
    echo -e "${BLUE}📋 Preview of changes since last release:${NC}"

    LAST_TAG=$(git --no-pager describe --tags --abbrev=0 2>/dev/null || echo "")

    if [[ "$LAST_TAG" == "" ]]; then
        echo -e "${YELLOW}📝 No tags found - this will be the first release (v0.1.0)${NC}"
        echo
        git --no-pager log --oneline -10
    else
        echo -e "${BLUE}🏷️  Last tag: $LAST_TAG${NC}"
        COMMITS_SINCE_TAG=$(git --no-pager rev-list --count "$LAST_TAG..HEAD")
        echo -e "${BLUE}📊 Commits since last tag: $COMMITS_SINCE_TAG${NC}"
        echo

        if [[ $COMMITS_SINCE_TAG -gt 0 ]]; then
            if check_git_cliff; then
                "$HOME/bin/git-cliff" --config cliff.toml --unreleased
            else
                git --no-pager log --format="%h %s" "$LAST_TAG..HEAD" | head -20
            fi
        else
            echo -e "${YELLOW}ℹ️  No new commits since last tag${NC}"
        fi
    fi
}

# Function to determine next version
determine_next_version() {
    LAST_TAG=$(git --no-pager describe --tags --abbrev=0 2>/dev/null || echo "")

    if [[ "$LAST_TAG" == "" ]]; then
        echo "v0.1.0"
        return
    fi

    COMMITS_SINCE_TAG=$(git --no-pager rev-list --count "$LAST_TAG..HEAD")
    if [[ "$COMMITS_SINCE_TAG" == "0" ]]; then
        echo ""
        return
    fi

    # Simple semantic versioning based on commit messages
    if git --no-pager log --format="%s" "$LAST_TAG..HEAD" | grep -q "^feat"; then
        BUMP="minor"
    elif git --no-pager log --format="%s" "$LAST_TAG..HEAD" | grep -q "^fix"; then
        BUMP="patch"
    else
        BUMP="patch"
    fi

    CURRENT_VERSION=${LAST_TAG#v}
    IFS='.' read -ra VERSION_PARTS <<< "$CURRENT_VERSION"
    MAJOR=${VERSION_PARTS[0]}
    MINOR=${VERSION_PARTS[1]}
    PATCH=${VERSION_PARTS[2]}

    case $BUMP in
        "minor")
            ((MINOR++))
            PATCH=0
            ;;
        "patch")
            ((PATCH++))
            ;;
        "major")
            ((MAJOR++))
            MINOR=0
            PATCH=0
            ;;
    esac

    echo "v${MAJOR}.${MINOR}.${PATCH}"
}

# Parse command line arguments
COMMAND=${1:-"preview"}

case "$COMMAND" in
    "preview"|"dry")
        echo -e "${BLUE}🔍 Dry run mode - no changes will be made${NC}"
        echo
        show_preview

        NEXT_VERSION=$(determine_next_version)
        if [[ "$NEXT_VERSION" != "" ]]; then
            echo -e "${GREEN}🎯 Next version will be: $NEXT_VERSION${NC}"
        else
            echo -e "${YELLOW}ℹ️  No version bump needed${NC}"
        fi
        ;;

    "prepare")
        echo -e "${BLUE}🚀 Preparing release...${NC}"
        echo

        # Run tests
        echo -e "${YELLOW}🧪 Running tests...${NC}"
        if ! make test; then
            echo -e "${RED}❌ Tests failed!${NC}"
            exit 1
        fi
        echo -e "${GREEN}✅ Tests passed${NC}"

        # Run linting
        echo -e "${YELLOW}🔍 Running linting...${NC}"
        if ! make lint; then
            echo -e "${RED}❌ Linting failed!${NC}"
            exit 1
        fi
        echo -e "${GREEN}✅ Linting passed${NC}"

        # Generate changelog
        echo -e "${YELLOW}📝 Generating changelog...${NC}"
        if check_git_cliff; then
            "$HOME/bin/git-cliff" --config cliff.toml --latest --output CHANGELOG.md
        else
            echo "# Changelog (generated)" > CHANGELOG.md
            echo "" >> CHANGELOG.md
            LAST_TAG=$(git --no-pager describe --tags --abbrev=0 2>/dev/null || echo "")
            if [[ "$LAST_TAG" == "" ]]; then
                git --no-pager log --format="- %s" -20 >> CHANGELOG.md
            else
                git --no-pager log --format="- %s" "$LAST_TAG..HEAD" >> CHANGELOG.md
            fi
        fi
        echo -e "${GREEN}✅ Changelog generated${NC}"

        # Show what will be released
        echo
        show_preview

        NEXT_VERSION=$(determine_next_version)
        if [[ "$NEXT_VERSION" != "" ]]; then
            echo -e "${GREEN}🎯 Ready to release: $NEXT_VERSION${NC}"
            echo
            echo -e "${BLUE}Next steps:${NC}"
            echo "1. Review the changelog in CHANGELOG.md"
            echo "2. Run: $0 release"
            echo "3. Push changes to trigger automatic release"
        else
            echo -e "${YELLOW}ℹ️  No changes to release${NC}"
        fi
        ;;

    "release")
        echo -e "${BLUE}🎯 Creating release...${NC}"
        echo

        # Check if changelog exists
        if [[ ! -f "CHANGELOG.md" ]]; then
            echo -e "${YELLOW}📝 CHANGELOG.md not found. Generating...${NC}"
            "$0" prepare
            echo
        fi

        NEXT_VERSION=$(determine_next_version)
        if [[ "$NEXT_VERSION" == "" ]]; then
            echo -e "${YELLOW}ℹ️  No changes to release${NC}"
            exit 0
        fi

        echo -e "${GREEN}🏷️  Releasing version: $NEXT_VERSION${NC}"
        echo

        # Commit changes
        if [[ -n "$(git --no-pager status --porcelain CHANGELOG.md)" ]]; then
            echo -e "${YELLOW}📝 Committing changelog...${NC}"
            git --no-pager add CHANGELOG.md
            git --no-pager commit -m "chore: update changelog for $NEXT_VERSION [skip ci]"
        fi

        # Create tag
        echo -e "${YELLOW}🏷️  Creating tag...${NC}"
        git --no-pager tag -a "$NEXT_VERSION" -m "Release $NEXT_VERSION"

        # Push changes and tag
        echo -e "${YELLOW}🚀 Pushing to remote...${NC}"
        git --no-pager push origin main
        git --no-pager push origin "$NEXT_VERSION"

        echo
        echo -e "${GREEN}✅ Release $NEXT_VERSION created successfully!${NC}"
        echo -e "${BLUE}🌐 GitHub Actions will now create the automated release.${NC}"
        echo -e "${BLUE}🔗 Watch the progress at: https://github.com/nchursin/verity-bdd/actions${NC}"
        ;;

    "help"|"-h"|"--help")
        echo "Usage: $0 [command]"
        echo
        echo "Commands:"
        echo "  preview, dry   - Preview changes without modifying anything"
        echo "  prepare        - Run tests, lint, and generate changelog"
        echo "  release        - Create and push the release"
        echo "  help           - Show this help"
        echo
        echo "Release process:"
        echo "1. Run '$0 preview' to see what will be released"
        echo "2. Run '$0 prepare' to prepare the release"
        echo "3. Run '$0 release' to create the release"
        ;;

    *)
        echo -e "${RED}❌ Unknown command: $COMMAND${NC}"
        echo
        $0 help
        exit 1
        ;;
esac

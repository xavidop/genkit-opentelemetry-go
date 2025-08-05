---
name: Bug Report
about: Report a bug in the Opentelemetry Plugin
title: 'fix: '
labels: ['bug', 'needs-triage']
assignees: ['xavidop']
---

## Bug Description

A clear and concise description of what the bug is.

## Steps to Reproduce

1. Go to '...'
2. Run command '...'
3. See error

## Expected Behavior

A clear and concise description of what you expected to happen.

## Actual Behavior

A clear and concise description of what actually happened.

## Error Output

```
Paste any error messages or stack traces here
```

## Environment

- **Go version**: 
- **Plugin version**: 
- **Operating system**: 

## Minimal Reproduction

If possible, provide a minimal code example that reproduces the issue:

```go
// Your minimal reproduction code here
```

## Additional Context

Add any other context about the problem here, such as:
- Screenshots
- Configuration files
- Related issues

## Checklist

- [ ] I have searched existing issues to ensure this is not a duplicate
- [ ] I have provided all the information requested above
- [ ] I have tested with the latest version of the plugin

## Fix Commit Message Preview

When fixing this bug, the commit message should follow conventional commits format:

```
fix(scope): brief description of the fix

Longer description of what was broken and how it's fixed.

Fixes #[issue-number]
```

---

**Note**: This project uses [Conventional Commits](https://conventionalcommits.org/) for automated semantic versioning and changelog generation.

# Pull Request Template

## Checklist

- [ ] Branch has been rebased with `master`
- [ ] Commits
    - [ ] Have been properly set-up
    - [ ] Merge strategy has been specified in the corresponding section
- [ ] Reviewers have been set
- [ ] New or changing code is properly tested
    - [ ] TDD has been applied
    - [ ] Coverage of at least the happy path has been checked
- [ ] Linters pass
- [ ] Blockers are specified
- [ ] After merging `master` can be deployed
- [ ] Documentation has been added, in godoc and/or README.md

## Goal

Why are you doing this. Relevant motivation and context.

## Merge strategy

The preferred merge strategy is:

- [ ] Create a merge commit
- [ ] Squash and merge
- [ ] Rebase and merge

## Blocked by

Add links to this project or other projects PRs, issues, reasons and/or people
who's preventing this from beeing merged.

## How are you doing this

Ideally this should be empty, the code should be clear and with godoc.

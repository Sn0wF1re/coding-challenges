
## Purpose — Practice & Learning

This repository is primarily a personal practice workspace for learning and improving skills in new or familiar programming languages. Use it to experiment, try different idiomatic approaches, keep short implementation notes, and track small projects you use to practice a language.

Each top-level language folder (for example `golang/`, `python/`, `javascript/`) contains self-contained challenge folders. Each challenge should include implementation, tests (when useful), and a short `README.md` describing how you ran or tested the solution and any lessons learned.

Repository overview

```text
coding-challenges/
├─ README.md                # This file — repo purpose and contribution notes
├─ golang/                  # Go solutions and instructions
│  ├─ README.md
│  ├─ go.mod
│  └─ <challenge-name>/
│     ├─ source files
│     ├─ test files
│     ├─ testdata/
│     └─ README.md
├─ python/                  # (example) Python solutions
└─ javascript/              # (example) JavaScript solutions
```


Naming conventions

- Use lowercase, kebab-case for folders (e.g. `prime-sieve`).
- Keep challenge folder names descriptive and short.
- Put implementation, tests, and fixtures inside the challenge folder.

Practice notes and expectations

- Treat these folders like a personal lab: drafts, experiments, and incomplete solutions are fine.
- Add a short `NOTES.md` or add a `Practice:` section in the challenge `README.md` describing what you practiced and what you learned.
- Keep runnable examples so you can quickly re-run solutions later.

Current languages

- Go — see the `golang` folder for instructions and existing challenges.

How to add a new language

1. Create a language folder at the repository root (for example `python/`).
2. Add a language-level `README.md` describing setup, common commands, and how to run tests for that language.
3. Add challenge folders inside the language folder following the challenge template below.

How to add a new challenge

1. In the chosen language folder, create a new folder in kebab-case: `<challenge-name>/`.
2. Add implementation files, tests, and an optional `testdata/` or `fixtures/` folder.
3. Add a short `README.md` in the challenge folder with usage and run instructions.

Suggested challenge template

```text
<language>/<challenge-name>/
├─ README.md            # short usage + run examples
├─ main/source file(s)
├─ test file(s)
└─ testdata/ (optional)
```


Contributing (for collaborators)

- This repo is focused on personal practice. If you accept contributions, prefer small, clearly labeled PRs and include a short explanation of the learning goal for the change.
- Include runnable examples in each challenge's `README.md` and a `Practice:` note explaining what the change practices or demonstrates.

Roadmap

- Implement the same challenges across multiple languages.
- Add consistent tests and CI for each language.
- Maintain a short matrix of implemented challenges by language.

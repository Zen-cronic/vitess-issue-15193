## Research Repo for the Issue [#15193](https://github.com/vitessio/vitess/issues/15193) from [Vitess](https://github.com/vitessio/vitess)

### Task: Find the occurrences of the old range loop format (e.g., `for i := 0; i < n; i++`) 

### Approach: using [ripgrep](https://github.com/BurntSushi/ripgrep?tab=readme-ov-file#ripgrep-rg):

`rg -n --pcre2 'for i :=(?!.*range)' > ../rg-for-range.txt`

- The `rg` command is run from the root of the cloned local `vitess` repo. (`--pcre2` flag needed for negative lookahead)
- It's then redirected into a file outside the repo for parsing. 

### Findings* :

- `815` occurrences 
- Across `350` unique files

### ***Note: As of merged PR [#16371](https://github.com/vitessio/vitess/pull/16371)**




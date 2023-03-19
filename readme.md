### Substring test
Read in the example files and create substrings based on the ranges under the "---" line by line.
- Constraints
 - All ranges are not overlapped
  - If two ranges have numbers overlapped, then they should be merged into one range
 - each range only contains x and y and x <= y

Example 1
```
按行政執行事件，行政處分經撤銷或變更確定者，執行機關應依職權或因義務人、利害關係人之申請終止執行。
---
1,5 // 按行政執行
6,7 // 事件
```

Example 2
```
按行政執行事件，行政處分經撤銷或變更確定者，執行機關應依職權或因義務人、利害關係人之申請終止執行。
---
1,5 -> 1,7 按行政執行事件
3,7
```

### How to use?
1. `git clone` the repo to local
2. `cd substring_test` then run `go run .`

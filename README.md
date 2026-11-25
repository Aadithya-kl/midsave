# midsave

**Crash-test your app's data. One command.**

midsave starts your app, kills it suddenly at a random moment (like a crash or a battery dying), restarts the check, and tells you if any data was lost or half-written. It repeats this many times so you find the bugs before your users do.

```
midsave run 1:  ok
midsave run 2:  BROKEN (killed after 1.803s)
...
7 of 10 runs lost or broke data
```

## Why

Most apps that save data (databases, queues, payment apps, game saves) are never tested for "what if it crashes in the middle of saving?" Testing it by hand is slow and boring, so people skip it.

Tools for this already exist (CrashMonkey, ALICE, LazyFS, Jepsen), but they are research or expert tools, mostly Linux-only, and hard to set up.

midsave aims to be the simple one:

- one command
- works on Windows, Mac and Linux
- plain-English results

## How it works

```
for each run:
  1. delete old data (fresh start)
  2. start your app
  3. wait a random time
  4. kill it
  5. run your checker
  6. record: ok or BROKEN
```

You write the **checker**, because only you know what "correct data" means for your app. The checker is any program that:

- exits with code `0` if the data is fine
- exits with any other code if the data is broken

## Quick start

Requires Go 1.22+.

```
git clone https://github.com/sharanyashetty013-droid/midsave.git
cd midsave

go build -o bin/writer.exe ./examples/writer
go build -o bin/badwriter.exe ./examples/badwriter
go build -o bin/check.exe ./examples/check
```

Test the buggy app:

```
go run ./cmd/midsave --app bin/badwriter.exe --check bin/check.exe --clean data.txt --runs 10
```

Test the safe app:

```
go run ./cmd/midsave --app bin/writer.exe --check bin/check.exe --clean data.txt --runs 10
```

On Mac/Linux, drop the `.exe`.

## Flags

| Flag | What it does | Default |
|---|---|---|
| `--app` | the program to crash-test | required |
| `--check` | the program that checks the data | required |
| `--clean` | file to delete before each run | none |
| `--runs` | how many times to crash it | 10 |

## Examples

| Example | What it does | Result |
|---|---|---|
| `examples/writer` | writes numbers 1 to 1000, one full line at a time | 0 of 10 broken |
| `examples/badwriter` | writes each number in two parts (a "torn write") | 7 of 10 broken |
| `examples/check` | checks the file has 1, 2, 3... with no broken lines | the checker |

## Limitations

Killing a process is **not** the same as a power cut. When a process is killed, data it already handed to the operating system still reaches the disk. So midsave v0.1 catches **app-level** bugs:

- data written in more than one step (torn writes)
- data still sitting in the app's memory
- unsafe file replacing

Simulating a real power cut (data the OS never wrote to disk) is planned for a later version.

## Roadmap

- [ ] `--seed` flag to replay a crash exactly
- [ ] better report: show where and why the data broke
- [ ] JSON output
- [ ] GitHub Action
- [ ] SQLite example
- [ ] power-cut mode

## Contributing

Issues and PRs are welcome. Look for issues labelled `good first issue`.

## License

MIT

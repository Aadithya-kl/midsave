<div align="center">

# 💥 midsave

### Your app saves data. What happens if it dies *mid-save*?

**midsave pulls the plug on your app hundreds of times and tells you exactly when your data breaks.**

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go&logoColor=white)
![Windows](https://img.shields.io/badge/Windows-✓-0078D6?logo=windows&logoColor=white)
![macOS](https://img.shields.io/badge/macOS-✓-000000?logo=apple&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-✓-FCC624?logo=linux&logoColor=black)
![License](https://img.shields.io/badge/license-MIT-green)
![PRs welcome](https://img.shields.io/badge/PRs-welcome-ff69b4)

</div>

---

## 🔌 The 10-second pitch

You pay ₹500. Your phone dies halfway through.
The money left your account. The shop never got it.

That's a **crash in the middle of a save**, and almost nobody tests for it, because testing it by hand means crashing your app over and over and checking the data every time.

**midsave does that for you. Automatically. In one command.**

---

## ⚡ See it catch a bug

Same checker. Two apps. One is safe, one has a hidden bug.

<table>
<tr>
<th>✅ Safe app</th>
<th>❌ Buggy app</th>
</tr>
<tr>
<td>

```
run 1:  ok
run 2:  ok
run 3:  ok
run 4:  ok
run 5:  ok
run 6:  ok
run 7:  ok
run 8:  ok
run 9:  ok
run 10: ok

0 of 10 runs lost or broke data
```

</td>
<td>

```
run 1:  ok
run 2:  BROKEN (killed after 1.803s)
run 3:  BROKEN (killed after 1.169s)
run 4:  ok
run 5:  BROKEN (killed after 2.157s)
run 6:  BROKEN (killed after 1.006s)
run 7:  ok
run 8:  BROKEN (killed after 1.945s)
run 9:  BROKEN (killed after 1.207s)
run 10: BROKEN (killed after 1.378s)

7 of 10 runs lost or broke data
```

</td>
</tr>
</table>

The buggy app writes each number in two steps. Kill it in between and you get **half a number** on disk: a *torn write*. midsave found it 7 times out of 10.

---

## 🧠 How it works

```
        ┌──────────────┐
        │ start app    │
        └──────┬───────┘
               ▼
        ┌──────────────┐
        │ wait random  │   ← could be 0.2s, could be 2s
        └──────┬───────┘
               ▼
        ┌──────────────┐
        │ 💥 KILL IT   │   ← no warning, no cleanup
        └──────┬───────┘
               ▼
        ┌──────────────┐
        │ run checker  │   ← is the data still okay?
        └──────┬───────┘
               ▼
          ok or BROKEN
               │
               └── repeat × N
```

You bring a **checker**: any program that exits `0` if the data is fine, and anything else if it's broken. Only you know what "correct" means for your app, so you decide.

---

## 🚀 Try it in 60 seconds

```bash
git clone https://github.com/sharanyashetty013-droid/midsave.git
cd midsave

go build -o bin/badwriter.exe ./examples/badwriter
go build -o bin/check.exe ./examples/check

go run ./cmd/midsave --app bin/badwriter.exe --check bin/check.exe --clean data.txt --runs 10
```

> On Mac/Linux, drop the `.exe`.

| Flag | What it does | Default |
|---|---|---|
| `--app` | the program to crash-test | required |
| `--check` | the program that checks the data | required |
| `--clean` | file to delete before each run | none |
| `--runs` | how many times to crash it | `10` |

---

## 🤔 Doesn't this already exist?

Sort of, but not for everyday developers.

| | Who it's for | Windows? | Setup |
|---|---|---|---|
| CrashMonkey | file system researchers | ❌ | heavy |
| ALICE / Pathfinder | academic research | ❌ | heavy |
| LazyFS | storage experts | ❌ | FUSE + config |
| Jepsen | distributed systems experts | ❌ | heavy |
| **midsave** | **any developer** | ✅ | **one command** |

---

## ⚠️ Honest limits

Killing a process ≠ a real power cut. Data your app already handed to the OS still reaches the disk when only the process dies. So **v0.1 catches app-level bugs**:

- 🧩 data written in more than one step (torn writes)
- 🧠 data still sitting in your app's memory
- 📄 unsafe file replacing

Simulating a true power cut is on the roadmap.

---

## 🗺️ Roadmap

- [ ] `--seed` to replay any crash exactly
- [ ] report *which* line broke and why
- [ ] JSON output
- [ ] GitHub Action: crash-test on every push
- [ ] SQLite + Postgres examples
- [ ] 🔌 power-cut mode

---

## 🤝 Contribute

New to open source? Perfect. Issues labelled **`good first issue`** are small, clear and ready to pick up. Every PR gets a reply within a day or two.

---

<div align="center">

**If midsave saved your data, give it a ⭐**

Made with 💥 by [Sharanya L Shetty](https://github.com/sharanyashetty013-droid)

</div>

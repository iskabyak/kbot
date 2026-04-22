DMR@DESKTOP-7B4KJLP MINGW64 ~/kbot (main)
$ git pull origin main --rebase
remote: Enumerating objects: 8, done.
remote: Counting objects: 100% (8/8), done.
remote: Compressing objects: 100% (6/6), done.
remote: Total 6 (delta 2), reused 0 (delta 0), pack-reused 0 (from 0)
Unpacking objects: 100% (6/6), 1.85 KiB | 75.00 KiB/s, done.
From https://github.com/iskabyak/kbot
 * branch            main       -> FETCH_HEAD
   c1df9b0..5848219  main       -> origin/main
Updating c1df9b0..5848219
Fast-forward
 README.md | 5 ++++-
 1 file changed, 4 insertions(+), 1 deletion(-)

DMR@DESKTOP-7B4KJLP MINGW64 ~/kbot (main)
$ cat README.md
## Setup

export TELE_TOKEN=your_token
go run main.go

DMR@DESKTOP-7B4KJLP MINGW64 ~/kbot (main)
$ git status
On branch main
Your branch is up to date with 'origin/main'.

nothing to commit, working tree clean

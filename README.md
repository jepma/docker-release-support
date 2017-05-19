# git-semtags

git-semtags is a git plugin command that shows, creates and bumps semantic version tags in your git repository.  It can be used 
to track the release of a single git repository, or to releases of different components in a single git repository.

see [docs/git-semtags.txt](git-semtags) manual page for details
 
## Examples

```bash
$ git init
$ touch README.md
$ git add README.md
$ git semtags --init version-
$ git commit -m 'initial import'
$ git semtags  --patch-release
Bumped to version 0.0.1 with tag version-0.0.1

$ touch README.more
$ git semtags
0.0.1-dirty

$ git add .
$ git commit -m 'more'
$ git semtags
0.0.1-bd7b174

$ git semtags --minor-release
Bumped to version 0.1.0 with tag version-0.1.0
```

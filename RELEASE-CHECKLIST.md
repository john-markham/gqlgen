# When gqlgen gets released, the following things need to happen
Assuming the next version is $NEW_VERSION=v0.16.0 or something like that.

1. Run the https://github.com/john-markham/gqlgen/blob/master/bin/release:
```
./bin/release $NEW_VERSION
```
2. git-chglog -o CHANGELOG.md
3. go generate ./...
4. git commit and push the CHANGELOG.md
5. Go to https://github.com/john-markham/gqlgen/releases and draft new release, autogenerate the release notes, and Create a discussion for this release
6. Comment on the release discussion with any really important notes (breaking changes)

I used https://github.com/git-chglog/git-chglog to automate the changelog maintenance process for now. We could just as easily use go releaser to make the whole thing automated.

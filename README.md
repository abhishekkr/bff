
## Big File Finder (BFF)

> there is just one purpose of this simple program, to find big files under a given target path
>
> and tries to do it as simply as it can

---

### Usage

* syntax `bff -path <PATH TO SCAN> -minsize <FILE ABOVE SIZE IN MB>`

#### Example

* check for all file size under `~/Downloads` with size greater than 750MB

```
bff -path ~/Downloads -minsize 750
```

* check for all file size under `~/Downloads` with size greater than 1GB

```
bff -path ~/Downloads -minsize 1024
```

* list directories also which have total size above limit, even if individual file escape the check

```
bff -dir -path ~/Downloads -minsize 1500
```

* if listing is erroneous, get to see more info using `-debug` switch as

```
bff -dir -path ~/Downloads -minsize 1500 -debug
```

---

### Install/Download

* [latest version](https://github.com/abhishekkr/bff/releases/latest)

* [v0.0.2 Release Page](https://github.com/abhishekkr/bff/releases/tag/v0.0.2), [v0.0.1 Release Page](https://github.com/abhishekkr/bff/releases/tag/v0.0.1)

> * [linux](https://github.com/abhishekkr/bff/releases/download/v0.0.2/bff-linux-amd64)
>
> * [macos](https://github.com/abhishekkr/bff/releases/download/v0.0.2/bff-darwin-amd64)
>
> * [windows](https://github.com/abhishekkr/bff/releases/download/v0.0.2/bff-windows-amd64)

---

# Reddit-image-downloader

######Written in go.

Downloads images from a subreddit into the project's root directory.

### Build

```
$ wire           # this will recreate the dependency injection container with any changes
$ mockery --all  # this will create mocks for any new interfaces you make
$ go build
```

### Run Tests

```
$ go test ./...
```

### Usage

```
$ ./reddit-downloader --subr [name] --limit [number of images to download]
```

* This project will need to sit in ~/Documents/ for it to run successfully 

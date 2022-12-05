# Objectives

You must follow the same instructions as in the first subject but the second argument must be the name of the template. I know some templates may be hard to read, just do not obsess about it.

# Instructions

+ Your project must be written in **Go**.
+ The code must respect the [**good practices**](https://01.kood.tech/git/root/public/src/branch/master/subjects/good-practices/README.md).
+ It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).
+ You can see all about the **banners** [here](https://01.kood.tech/git/root/public/src/branch/master/subjects/ascii-art/).
+ The usage must respect this format `go run . [STRING] [BANNER]`, any other formats must return the following usage message:

```
Usage: go run . [STRING] [BANNER]

EX: go run . something standard
```

If there are other `ascii-art` optional projects implemented, the program should accept other correctly formatted `[OPTION]` and/or `[BANNER]`.
Additionally, the program must still be able to run with a single `[STRING]` argument.

```
Usage

$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $

$ go run . "Hello There!" shadow | cat -e
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $

$ go run . "Hello There!" thinkertoy | cat -e
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $
```

# Allowed packages

+ Only the standard Go packages are allowed
This project will help you learn about :

+ The Go file system(**fs**) API
+ Data manipulation

## Testing

Make `test.sh` file executable 
```
chmod +x test.sh
```
Run `test.sh` file for tests
```
./test.sh
```
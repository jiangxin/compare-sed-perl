Compare consumed resource for sed and perl command.

## Environment

* sed (builtin): macOS 11.3.1
* gsed (GNU sed): 4.8
* perl 5.32.0

## Result

    $ go build && ./compare-sed-perl 
    Command: sed  ..., MaxRSS: 905216
    Command: gsed ..., MaxRSS: 2043904
    Command: perl ..., MaxRSS: 2338816

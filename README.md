# action-test

Testing tool for actions.

## Usage

Test actions

```sh
$ action-test
Usage of action-test:
  Application action-test runs tests from JSON or YAML files against actions from different sources.

Flags:
  -debug=false: Log debug info
  -sources="": Actions sources (comma separated directories & urls)
  -tests="": Files or directory containing YAML or JSON tests (can be glob pattern)
  -v=false: Verbose output: log all tests

Example YAML tests:

  -
    name: "filmweb.find"
    description: "Should find movie by title"
    arguments:
      title: "Pulp Fiction"
    expect:
      writers: "Quentin Tarantino"
      directors: "Quentin Tarantino"
      title: "Pulp Fiction"
      year: "1994"

$ action-test -tests=./example/tests.yaml -sources=./example/actions
=== RUN filmweb.find

  Should find Pulp Fiction

    √ genres => Gangsterski
    √ description => Przemoc i odkupienie w opowieści o dwóch płatnych mordercach pracujących na zlecenie mafii, żon...
    √ writers => Quentin Tarantino
    √ directors => Quentin Tarantino
    √ year => 1994
    √ poster => http://1.fwcdn.pl/po/10/39/1039/7517880.3.jpg
    √ countries => USA
    √ title => Pulp Fiction
    √ rating => 8,5

--- PASS: filmweb.find (234.0134ms)
...
```

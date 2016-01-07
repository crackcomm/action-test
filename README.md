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

$ action-test -tests=./example/tests.yaml -sources=file://example/actions
=== RUN filmweb.find

  Should find Pulp Fiction

    √ title => Pulp Fiction
    √ year => 1994
    √ countries => USA
    √ genres => Gangsterski
    ? description => Przemoc i odkupienie w opowieści o dwóch płatnych mordercach pracujących na zlecenie mafii, żon...
    ? poster => http://1.fwcdn.pl/po/10/39/1039/7517880.3.jpg
    √ writers => Quentin Tarantino
    √ directors => Quentin Tarantino

--- PASS filmweb.find (132.529264ms)
=== RUN filmweb.find

  Should find Apollo 13

    √ year => 1995
    √ countries => USA
    √ genres => Dramat
    ? poster => http://1.fwcdn.pl/po/10/53/1053/7473756.3.jpg
    ? description => Podczas rutynowego lotu w kosmos na pokładzie statku Apollo 13 następuje wybuch.
    √ directors => Ron Howard
    √ writers => [William Broyles Jr. Al Reinert]
    √ title => Apollo 13

--- PASS filmweb.find (153.756246ms)
=== RUN filmweb.find

  Should find Avatar

    √ directors => James Cameron
    √ writers => James Cameron
    √ title => Avatar
    √ year => 2009
    ? poster => http://1.fwcdn.pl/po/91/13/299113/7332755.3.jpg
    ? description => Jake, sparaliżowany były komandos, otrzymuje misję i zostaje wysłany na planetę Pandora, gdzie ...
    √ genres => Sci-Fi
    √ countries => [USA Wielka Brytania]

--- PASS filmweb.find (171.49564ms)
--- OK   457.927317ms
```

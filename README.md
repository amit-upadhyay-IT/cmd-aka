# cmd-alias
> A command line application with features like cmd alias with parameters, etc

## Design:
- The app is divided into different modules, each module is responsible for
  performing a specific type of problem in this app

### Levelized diagram:

                  main
                /  | \  \
               /   |  \  \
              \/   |   \  \
        operations |   \/  \
           /\      |   aka  \
           |       \/        \
           |---tokenparser    \
                 |            \/
                 |          utils
                 |          /   \
                 |         /     \
                 |        \/     \/
                 ----->aka_util   io

- **main** : is the entry point of app
- **operations** : Each possible operation that can be performed in this app is defined
  here as enum
- **tokenparser** : this is a parser, it takes token as input and returns the
  operation that needs to be performed along with the operators also. (A token
  is an input given on command line).
- **aka** : aka stands for `Also Known As`, its the module which has actual
  business logic in its method.
- **utils** : this module containis a set of utility functions.
- **utils/io** : conatins input/output related utility functions.
- **utils/aka_util** : contains utility functions for `aka` module.

### The work is in progress

### TODOs:
- write unit testcases for each module
- complete the logics in `tokenparser`
- Think on how would you go upon writing unit testcases for method which are
  private (or which have package level access):
    - Do we want segregate testcases to different modules?
    - Or, consider them writing in a different module? then, maybe we will need
      to grab the std output from console and assert on them?:wq

### archived stuff
- aku set_-alias key value: this will just store the key value in a file

- aku set-alias some__complex_command {p1} some_things_else {p2}... | some_simple_command {val1}, {val2} : this will store the key with the place holders(which are {p1}, {p2}, etc) and value will be the command after the pipe(|), NOTE while setting up the alais, we don't need to put the values, but while using the alias we will need to input the values for the place holders


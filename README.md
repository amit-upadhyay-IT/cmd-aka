#cmd-alias
>> A command line application with features like cmd alias with parameters, etc

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


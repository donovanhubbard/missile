# missile
An interactive CLI for interacting with memcached

## TODO

- Create a backend object
- Have the backend object create a client for every host so you can ping them individually. These will need to use go routines
- When you perform operations you should list what server was written to
- create the interface. It should have
  - A text area to accept new commands. It should have autocomplete the various commands and arguments
  - A server list that tells if servers are online or not
  - A viewport that lists the output of commands and the command history
  - Maybe a disconnect button? Or maybe they will just type quit.
  - See if you can add a title pane or icon or something


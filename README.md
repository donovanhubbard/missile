# missile
An interactive CLI for interacting with memcached

## TODO

- Delete commandhistory text after a scrollback number has been passed
- consider moving the model/model.go object out of the model directory becuase it will be the core logic for the whole app.
- Create a backend object
- Have the backend object create a client for every host so you can ping them individually. These will need to use go routines
- When you perform operations you should list what server was written to

## Interseting unicode characters
The following are some interesting unicode characters you may want to use
:  ⃝ U+20DD
⏺ U+23FA
● U+25CF
⃠ U+20E0
∅ U+2205
⊘ U+2298
⚠ U+26A0
⚡U+26A1


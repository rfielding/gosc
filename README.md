gosc
====

A simple project for doing audio generation with Go, starting under Linux.
The first steps are simple.  This current code is just smoke testing the idea for feasability.
 
Build the binary, and it will emit CD standard audio on stdout.  Redirect it into a command-line program.  So with the built gosc program, do this to emit audio while listening on UDP port 9999 from anywhere:

./gosc 0.0.0.0:9999 | aplay -f cd --buffer-size=256

For now, it's a simple stereo tone.  It is using 5ms buffers in aplay, so the code should probably follow suit and use either 128 or 256 buffer size for maximum performance. 

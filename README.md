gosc
====

A simple project for doing audio generation with Go, starting under Linux.
The first steps are simple.

Build the binary, and it will emit CD standard audio on stdout.  Redirect it into a command-line program.  So with the built gosc program, do this to emit audio while listening on UDP port 9999 from anywhere:

./gosc 0.0.0.0:9999 | aplay -f cd

For now, it's a simple stereo tone to smoke-test the idea.  Todo is to add in listening on UDP port (preferably the SuperCollider standard port).  From there I will listen for OSC packets in a well-defined format that's emitted by my Windows 8 project (CSharpAttempt).

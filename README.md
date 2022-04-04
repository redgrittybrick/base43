Base43
======

This is a simple command line tool that can encode to, or decode from, the
Base43 encoding that is used by the Electrum wallet software for Bitcoin.

Usage
=====

Show command line options

    base43 -h

Decode a base43 file and display in hexadecimal

    base43 -decode -hex < mydata-b43.txt

Read some hexadecimal data from a file, encode into Base43 and write to a file

    base43 -hex < mydata-hex.txt > mydata-b43.txt

Note that the base43 tool is designed to be able to be used in a pipeline 
(just like many Unix/Linux/GNU tools can be).

    somecommand | base43 ... | othercommand

So the way to specifiy input from a file is to use your command-shell's
input-redirection operator `< filename`. The way to get output  into a file is
to use your command-shell's output redirection operator `> filename`.
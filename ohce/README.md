ohce Kata
=================

Source
https://kata-log.rocks/ohce-kata

This kata has been slightly adjusted from the original.

Requirements
------------
## Your task

Write a console application that mimics a conversation with a friend named Ohce (which is just echo spelled backwards).

The application must read a string from the console, reverse it, and then print it back with some specific behavior:

- Between 6 and 12 hours, ohce will greet you saying: ¡Good morning < your name >!
- Between 12 and 18 hours, ohce will greet you saying: ¡Good afternoon < your name >!
- Between 18 and 21 hours, ohce will greet you saying: ¡Good evening < your name >!
- Between 21 and 6 hours, ohce will greet you saying: ¡Good night < your name >!
- When you introduce a palindrome, ohce likes it and after reverse-echoing it, it adds ¡Nice palindrome!
- ohce knows when to stop, you just have to write Stop! and it'll answer Goodbye < your name > and end.

## Example
This is an example of using ohce during the morning:
```
$ Hi, I'm Ohce! What's your name?
$ < your name >
> ¡Good morning <your name>!
$ hello
> olleh
$ oto
> oto
> ¡Nice palindrome!
$ stop
> pots
$ Stop!
> Goodbye <your name>
```


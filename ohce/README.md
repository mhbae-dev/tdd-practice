ohce Kata
=================

Source
https://kata-log.rocks/ohce-kata

This kata has been altered for english speakers.

Requirements
------------
## Your task

ohce is a console application that echoes the reverse of what you input through the console.

Even though it seems a silly application, ohce knows a thing or two.

When you start ohce, it greets you differently depending on the current time in English:
- Between 6 and 12 hours, ohce will greet you saying: ¡Good morning < your name >!
- Between 12 and 18 hours, ohce will greet you saying: ¡Good afternoon < your name >!
- Between 18 and 21 hours, ohce will greet you saying: ¡Good evening < your name >!
- Between 21 and 6 hours, ohce will greet you saying: ¡Good night < your name >!
- When you introduce a palindrome, ohce likes it and after reverse-echoing it, it adds ¡Nice palindrome!
- ohce knows when to stop, you just have to write Stop! and it'll answer Goodbye < your name > and end.

## Example
This is an example of using ohce during the morning:
```
$ ohce <your name>
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



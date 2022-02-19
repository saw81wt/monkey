# 概要

goで作るインタプリタの写経。

https://www.oreilly.co.jp/books/9784873118222/


```
$ go run main.go
Hello XXX! This is the Monkey programming language!
Feel free to type in commands
>> let hoge = 1;
>> hoge
1
>> let fuga = "fuga";
>> fuga
fuga
>> let foo = 1 + 2;
>> foo
3
>> let add = fn(x, y) { return x + y; };
>> add(1, 2)
3
>> let arr = [1, 2, 3];
>> arr[0]
1
>> let dic = {"hoge": 1};
>> dic["hoge"]
1
```
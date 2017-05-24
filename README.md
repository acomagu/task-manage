# Task Manager on Command line
![25314.png](https://qiita-image-store.s3.amazonaws.com/0/89142/0b14626e-84e5-b7cb-108c-d0dcbe0cc27a.png "25314.png")
This is tasks manager on command line
I create this easy task manager.

## Install

You type these commands.
You get my easy task manager.
 
```
$ go get github.com/Yamashou/task-manage
$ cd $GOPATH/src/github.com/Yamashou/task-manage
$ go build
$ task-manage init
```

## Sub Commands

### init
Create a place to stock the task

```
$task-manage init
```

### add

Add a task.

```
$task-manage add <task title(it is file name)> <Task content> <
Number of days to take the task>
```




#### example

```
$task-manage add example 例えばのタスク 4
Title :  example
Content :  例えばのタスク
Dead Line :  2017-05-28
---------------------------------------------------------------

```

## edit

Edit task

```
$task-manage edit <Old task name> <New task name> <New task content> <New Number of days to take the task>
```

#### example

```
$task-manage edit example example1 例えばのタスクです！ 5
Title :  example1
Content :  例えばのタスクです！
Dead Line :  2017-05-29
---------------------------------------------------------------
```

## done

Completion of task

```
$task-manage done <Task name>
```
### example

```
$task-manage done example1
Title :  example1
Content :  例えばのタスクです！
Dead Line :  2017-05-29
Done Time :  2017-05-24
---------------------------------------------------------------
```
## show
View incomplete tasks

```
$task-manage show <Task name>
$task-manage show
```

## example
```
$task-manage show example1 
Title :  example1
Content :  例えばのタスクです！
Dead Line :  2017-05-29
---------------------------------------------------------------
$task-manage show 
Title :  api.token
Content :  apiコール用のトークン
Dead Line :  2017-05-25
---------------------------------------------------------------
Title :  example1
Content :  例えばのタスクです！
Dead Line :  2017-05-29
---------------------------------------------------------------
Title :  fix.task-manage.api.delete
Content :  デリート機能の追加
Dead Line :  2017-05-22
---------------------------------------------------------------
Title :  sporters
Content :  講座記入
Dead Line :  2017-06-18
---------------------------------------------------------------
Title :  track-money
Content :  管理お金
Dead Line :  2017-05-23
---------------------------------------------------------------
```

## list 
All task display

```
$task-manage list
```
### example
```
$task-manage list
Title :  self.add
Content :  編集機能追加
Dead Line :  2017-05-15
Done Time :  0001-01-01
---------------------------------------------------------------
Title :  tester
Content :  ee
Dead Line :  2017-05-14
Done Time :  2017-05-14
---------------------------------------------------------------
Title :  ttttt
Content :  44444
Dead Line :  2017-05-15
Done Time :  2017-05-11
---------------------------------------------------------------
Title :  ubic
Content :  勉強会2
Dead Line :  2017-05-24
Done Time :  2017-05-23
---------------------------------------------------------------
Title :  example1
Content :  例えばのタスクです！
Dead Line :  2017-05-29
---------------------------------------------------------------
Title :  Supporter
Content :  講座記入
Dead Line :  2017-06-18
---------------------------------------------------------------
Title :  track-money
Content :  管理お金
Dead Line :  2017-05-23
---------------------------------------------------------------
```

### delete
Delete task

```
$task-manage delete <task name> <-o have, -f finished>
```
### example
```
$task-manage delete tester -f       
Delete  tester
$task-manage delete tester -o
Delete  tester
```

### recode
Save the current situation

```
$task-manage recode
```

### push
Send record to api

The api code is [here](https://github.com/Yamashou/task-manage-api)

```
$task-manage push
```









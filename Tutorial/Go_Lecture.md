<h1 style="text-align:center">Go</h1>
<p style="text-align:right">Provided by <b>Tae Geun Kim</b></p>

## 1. What is the **Go**?

> Go (often referred to as golang) is a free and open source programming language created at **Google**.   
> ... It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

## 2. Why do we use **Go**?

### - Programming Language Hall of Fame
Year | Winner
-----|-------
2016 |  Go
2015 |  Java
2014 |  JavaScript
2013 |  Transact-SQL
2012 |  Objective-C
2011 |  Objective-C
2010 |  Python
2009 |  Go
2008 |  C
2007 |  Python
2006 |  Ruby
2005 |  Java
2004 |  PHP
2003 |  C++

* Of course, Go is not the most popular language.

### - TIOBE Index for June 2017

Jun 2017 | Jun 2016 | Change | Programming Language | Ratings | Change
---------|----------|--------|----------------------|---------|-------
1 | 1 |  | Java | 14.493% | -6.30%
2 | 2 |  | C | 6.848% | -5.53%
3 | 3 |  | C++ | 5.723% | -0.48%
4 | 4 |  | Python | 4.333% | +0.43%
5 | 5 |  | C# | 3.530% | -0.26%
6 | 9 | Up | Visual Basic .NET | 3.111% | +0.76%
7 | 7 |  | JavaScript | 3.025% | +0.44%
8 | 6 | Down | PHP | 2.774% | -0.45%
9 | 8 | Down | Perl | 2.309% | -0.09%
10 | 12 | Up | Assembly language | 2.252% | +0.13%
11 | 10 | Down | Ruby | 2.222% | -0.11%
12 | 14 | Up | Swift | 2.209% | +0.38%
13 | 13 |  | Delphi/Object Pascal | 2.158% | +0.22%
14 | 16 | Up | R | 2.150% | +0.61%
15 | 48 | Great Up | Go | 2.044% | +1.83%
16 | 11 | Down | Visual Basic | 2.011% | -0.24%
17 | 17 |  | MATLAB | 1.996% | +0.55%

### - RedMonk Ranking

<img src="http://sogrady-media.redmonk.com/sogrady/files/2017/06/lang.rank_.617.wm_.png"></img>

### - Additional References

> <a href="https://insights.stackoverflow.com/survey/2016#technology-most-loved-dreaded-and-wanted" target="blank">Most Loved and Dreaded Languages of Programmers 2017</a>

## 3. Advantages of Go

1. Simple
    > **- Example Code**
    > - C (Hello World)
    > ```C
    > #include<stdio.h>
    > int main(int argc, char * argv[])
    > {
    >  printf("Hello World!\n");
    >  return 0;
    > }
    >```
    >
    > - Python (Hello World)
    > ```Python
    > print("Hello World!\n")
    > ```
    > - Go (Hello World)
    > ```Go
    > package main
    > 
    > import "fmt"
    >
    > func main() {
    >   fmt.Println("Hello World!\n")    
    >}
    >```
     - Simplicity : Python > Go >> C >> Java
2. Fast Compile
    > Go 언어의 특징은 컴파일 언어이지만 컴파일러의 컴파일 속도가 매우 빨라 인터프리터 언어처럼 쓸 수 있다는 점에 있다.[13] 이는 언어의 문법 구조를 개선함으로써 달성하였다. 컴파일러가 소스 코드를 해석하는 pass 수를 줄여서 달성한 것으로 보인다. 접근하기 어렵지 않고, 코드 역시 간결하면서도 컴파일 언어답게 높은 성능을 낼 수 있다는 점이 호평을 받는다.  
    -출처: 나무위키-

3. Just Fast

source | secs | mem | gz | cpu |
-------|------|-----|----|-----|
Go | 1.98 | 3,044 | 1344 | 5.65 |
Python3 | 110.91 | 8,024 | 977 | 110.87 |

- Reference (<a href="http://benchmarksgame.alioth.debian.org/u64q/compare.php?lang=go&lang2=python3" target="blank">Benchmark Game</a>) 
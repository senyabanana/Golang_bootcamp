# Day 08 — Go Boot camp

## Adventure, Danger and Cocoa

## Contents

1. [Chapter I](#chapter-i) \
    1.1. [General rules](#general-rules)
2. [Chapter II](#chapter-ii) \
    2.1. [Rules of the day](#rules-of-the-day)
3. [Chapter III](#chapter-iii) \
    3.1. [Intro](#intro)
4. [Chapter IV](#chapter-iv) \
    4.1. [Exercise 00: Arithmetic](#exercise-00-arithmetic)
5. [Chapter V](#chapter-v) \
    5.1. [Exercise 01: Botany](#exercise-01-botany)
6. [Chapter VI](#chapter-vi) \
    6.1. [Exercise 02: Hot Chocolate](#exercise-02-hot-chocolate)


<h2 id="chapter-i" >Chapter I</h2>
<h2 id="general-rules" >General rules</h2>

- Your programs should not exit unexpectedly (give an error on valid input). If this happens, your project will be considered non-functional and will receive a 0 in the evaluation.
- We encourage you to create test programs for your project, even though this work doesn't have to be submitted and won't be graded. This will allow you to easily test your work and the work of your peers. You will find these tests particularly useful during your defense. In fact, you are free to use your tests and/or the tests of the peer you are evaluating during your defense.
- Submit your work to your assigned git repository. Only the work in the git repository will be evaluated.
- If your code uses external dependencies, it should use [Go Modules](https://go.dev/blog/using-go-modules) to manage them.

<h2 id="chapter-ii" >Chapter II</h2>
<h2 id="rules-of-the-day" >Rules of the Day</h2>

- You should only submit `*.go` files and (in case of external dependencies) `go.mod` + `go.sum`.
- Your code for this task should be buildable with just `go build`.

<h2 id="chapter-iii" >Chapter III</h2>
<h2 id="intro" >Intro</h2>

People tend to say that one of the main differences between Go and C is a pointer safety. That's partly true, and Go will try really hard not to let you shoot yourself in the foot if you tango with pointers. But we're already far enough into the jungle to be able to play a little with danger, don't you think?

<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="ex00">Exercise 00: Arithmetic</h3>

Here in a jungle you can find some strange creatures that you have to treat in an unusual way. For this task you need to write a function `getElement(arr []int, idx int) (int, error)` that accepts and an index and returns the element with that index. Seems simple enough, right? But there is one condition — you can't use lookup by this index (like `arr[idx]`), the only lookup allowed is a first element (`arr[0]`). You may need to remember some C to complete this exercise.

On any invalid input (empty slice, negative index, index is out of bounds), the function should return an error with a textual explanation of the problem.

<h2 id="chapter-v" >Chapter V</h2>
<h3 id="ex01">Exercise 01: Botany</h3>

You're in luck! You've found some very rare plants:

```
type UnknownPlant struct {
    FlowerType string
    LeafType string
    Color int `color_scheme:"rgb"`
}

type AnotherUnknownPlant struct {
    FlowerColor int
    LeafType string
    Height int `unit:"inches"`
}
```

Well, yes, the current representation is a bit of a mess. Your goal would be to write a single function `describePlant` that accepts any type of plant (yes, it should work with structures of different types) and then prints all the fields as key-value pairs, separated by commas (note the tags), like this

```
FlowerColor:10
LeafType:lanceolate
Height(unit=inches):15
```

<h2 id="chapter-vi" >Chapter VI</h2>
<h3 id="ex02">Exercise 02: Hot Chocolate</h3>

Okay, now it's time to relax and have some cocoa. Cocoa usually comes in packages (see the provided zip archive). You don't need to modify the code in the packaged files in any way, the only thing you need to do is to write code (including the Cocoa files as part of your project) that creates a standard empty Mac OS GUI window (size 300x200) with the title "School 21". It's easier than you think!


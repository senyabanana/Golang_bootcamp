# Day 07 — Go Boot camp

## Moneybag

## Contents

1. [Chapter I](#chapter-i) \
    1.1. [General rules](#general-rules)
2. [Chapter II](#chapter-ii) \
    2.1. [Rules of the day](#rules-of-the-day)
3. [Chapter III](#chapter-iii) \
    3.1. [Intro](#intro)
4. [Chapter IV](#chapter-iv) \
    4.1. [Exercise 00: King's Bounty](#exercise-00-kings-bounty)
5. [Chapter V](#chapter-v) \
    5.1. [Exercise 01: Need for Speed](#exercise-01-need-for-speed)
6. [Chapter VI](#chapter-vi) \
    6.1. [Exercise 02: Elder Scrolls](#exercise-02-elder-scrolls)


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
- All your tests should be executable with the standard `go test ./...` call.

<h2 id="chapter-iii" >Chapter III</h2>
<h2 id="intro" >Intro</h2>

"There are several areas where we consider reliability and speed to be critical. Areas that directly affect people's lives — medicine, air safety, finance. Of course, this means that we thoroughly examine every detail of our product before releasing it to the public. Ladies and gentlemen, I give you... the Moneybag!"

<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="ex00">Exercise 00: King's Bounty</h3>

You keep listening to the CEO's voice, but your eyes are looking at the code on your laptop.

Sometimes it seems that people always use coins to pay for things. In laundromats, vending machines or jukeboxes, it is still normal to accept only pieces of metal as payment. But sometimes people hate standing in line and waiting for someone else to collect coins and put them in one at a time. Why can't people just always use a minimal amount of coins to avoid slowing everyone down?

This is a fairly common problem, and your colleague has already written some code and uploaded it to you for review:

```
func minCoins(val int, coins []int) []int {
    res := make([]int, 0)
    i := len(coins) - 1
    for i >= 0 {
        for val >= coins[i] {
            val -= coins[i]
            res = append(res, coins[i])
        }
        i -= 1
    }
    return res
}
```

It accepts a required amount and a sorted slice of unique coin denominations. It can be something like [1,5,10,50,100,500,1000] or something exotic like [1,3,4,7,13,15]. The output is supposed to be a slice of coins of minimal size that can be used to express the value (e.g., for 13 and [1,5,10], you should get [10,1,1,1]).

The problem is that you have a gut feeling that something is wrong with this code. Your goal here is to write several tests (in `*_test.go` files) for this code that will show it is producing wrong results. You will also need to write a separate function (you should call it `minCoins2`) that will have the same parameters, but will successfully handle these cases. If there are duplicates in a denomination slice, or it is not sorted, the function should still return a correct result. If it is empty, it should return an empty slice. 

<h2 id="chapter-v" >Chapter V</h2>
<h3 id="ex01">Exercise 01: Need for Speed</h3>

Now that you have a new version of the code from EX00, let's test it for performance. Your goals here are

 - Get a list of the top 10 functions in your code (call your function with some test data) that your CPU spends the most time executing (you should use Go's built-in tools for this). Submit this list as a `top10.txt` file.
 
 - Write a benchmark version of your tests that compares the performance of your new code to the old one, especially when using relatively large denomination slices. If you find more optimizations during this phase, feel free to submit a newer version of your `minCoins2` function, calling it `minCoins2Optimized` (not a required step).

<h2 id="chapter-vi" >Chapter VI</h2>
<h3 id="ex02">Exercise 02: Elder Scrolls</h3>

Now that you've fixed the bug and written some tests for your code, it's time to create some documentation for it. Use comments in your code to describe how your solution differs from the standard one, and what optimizations you used. Then use any tool you can find to generate HTML documentation based on those comments.

Instructions on how to reproduce the documentation generation should also be included in the comments. Saving HTML pages from the web browser is considered cheating (though not strictly prohibited, so if you couldn't do it any other way, just say so explicitly in the comments).

Submit generated documentation (HTML files + static stuff like images, js and css) packed into a `docs.zip` archive.


# Day 09 — Go Boot camp

## Daily Routine

## Contents

1. [Chapter I](#chapter-i) \
    1.1. [General rules](#general-rules)
2. [Chapter II](#chapter-ii) \
    2.1. [Rules of the day](#rules-of-the-day)
3. [Chapter III](#chapter-iii) \
    3.1. [Intro](#intro)
4. [Chapter IV](#chapter-iv) \
    4.1. [Exercise 00: Sleepsort](#exercise-00-sleepsort)
5. [Chapter V](#chapter-v) \
    5.1. [Exercise 01: Spider-Sens](#exercise-01-spider-sense)
6. [Chapter VI](#chapter-vi) \
    6.1. [Exercise 02: Dr. Octopus](#exercise-02-dr-octopus)


<h2 id="chapter-i" >Chapter I</h2>
<h2 id="general-rules" >General rules</h2>

- Your programs should not exit unexpectedly (give an error on valid input). If this happens, your project will be considered non-functional and will receive a 0 in the evaluation.
- We encourage you to create test programs for your project, even though this work doesn't have to be submitted and won't be graded. This will allow you to easily test your work and the work of your peers. You will find these tests particularly useful during your defense. In fact, you are free to use your tests and/or the tests of the peer you are evaluating during your defense.
- Submit your work to your assigned git repository. Only the work in the git repository will be evaluated.
- If your code uses external dependencies, it should use [Go Modules](https://go.dev/blog/using-go-modules) to manage them.

<h2 id="chapter-ii" >Chapter II</h2>
<h2 id="rules-of-the-day" >Rules of the Day</h2>

- You should only print `\*.go`, `\*_test.go` and (in case of external dependencies) `go.mod` + `go.sum` files.
- Your code for this task should be buildable with just `go build`.
- All your tests should be executable with the standard `go test ./...` call.

<h2 id="chapter-iii" >Chapter III</h2>
<h2 id="intro" >Intro</h2>

Sometimes we hear that there are some people who can "do several things at once". While it is theoretically possible to do completely different things (e.g., with different hands), this phrase usually refers to people who do multiple tasks at the same time, but NOT in parallel. What do I mean by that?

In computer science, "parallel" usually means making progress on more than one task at a time. But with humans, it is a bit different - the real trick is to maintain context and switch from one task to another quickly. It may even look like "parallelism" from the side, but it's not — it's *concurrency*, which is a slightly broader concept. And yes, it means that concurrency can be implemented with parallelism, but it can also work without it (as most people do).

When we program things to run in parallel, we generally think of explicitly creating multiple separate threads and giving each of them a target function. But that's not how Golang works — it works in terms of *concurrency*, which means we don't really need to think about if and how actual parallelization is happening under the hood.

This gives us a lot of power, but with great power comes great responsibility...

<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="ex00">Exercise 00: Sleepsort</h3>

So let's write a toy algorithm as an example. It is pretty useless for production, but it helps to understand the concept.

You need to write a function called `sleepSort` which takes an unsorted slice of integers and returns an integer channel where these numbers are written one by one in sorted order. To test it, you should read and print output values from a returned channel in the main goroutine and gracefully exit the application at the end.

The idea of Sleepsort (what makes it a "toy") is that we'll spawn a number of goroutines equal to the size of an input slice. Then each goroutine sleeps for a number of seconds equal to the number it received. Then it wakes up and sends that number to the output channel. It's easy to see that this way the numbers are returned in a sorted order.

<h2 id="chapter-v" >Chapter V</h2>
<h3 id="ex01">Exercise 01: Spider-Sense</h3>

You probably remember how Peter Parker realized he had superpowers when he woke up in the morning. Well, let's write our own spider (or crawler) to parse the web. You need to implement a function `crawlWeb` that accepts an input channel (for sending URLs) and returns another channel for crawling results (pointers to web page bodies as strings). Also, there shouldn't be more than 8 goroutines querying pages in parallel at any given time.

But we want to be fast and flexible, so another requirement is to be able to stop the crawling process at any time by pressing Ctrl+C (and your code should perform a graceful shutdown). For this, you can add more input arguments to the `crawlWeb` function, which should be either `context.Context` for abort or `done` channel. If not interrupted, the program should gracefully exit after processing all given URLs.

<h2 id="chapter-vi" >Chapter VI</h2>
<h3 id="ex02">Exercise 02: Dr. Octopus</h3>

Okay, now we have to kill the villain! The main problem with Dr. Octopus is that he has a lot of tech tentacles, and it's hard to keep track of them all. Let's tie them together!

For this exercise, you need to write a function called `multiplex`, which should be *variadic* (taking a variable number of arguments). It should take channels (`chan interface{}`) as arguments and return a single channel of the same type. The idea is to redirect all incoming messages from these input channels to a single output channel, effectively implementing a "fan-in" pattern.

As proof of work, you should write a test on sample data that explicitly shows that all values randomly sent to any input channels will be received on the same output channel.

And... you've just defeated a villain!

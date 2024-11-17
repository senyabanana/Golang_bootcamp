# Team 01 — Go Boot camp

## Warehouse 13

## Contents

1. [Chapter I](#chapter-i) \
    1.1. [General rules](#general-rules)
2. [Chapter II](#chapter-ii) \
    2.1. [Rules of the day](#rules-of-the-day)
3. [Chapter III](#chapter-iii) \
    3.1. [Intro](#intro)
4. [Chapter IV](#chapter-iv) \
    4.1. [Task 00: Scalability](#exercise-00-scalability)
5. [Chapter V](#chapter-v) \
    5.1. [Task 01: Balancing and Queries](#exercise-01-anomaly-balancing-and-queries)
6. [Chapter VI](#chapter-vi) \
    6.1. [Task 02: Long Live the King](#exercise-02-long-live-the-king)
7. [Chapter VII](#chapter-vii) \
    7.1. [Task 03: Consensus](#exercise-03-consensus)

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

"Oh, come on, Artie, this is ancient!" Lattimer was about to pull his own hair out. "It's the 21st century, nobody uses pen and paper to catalog things anymore!"

"What do you want me to do, Pete? It's been this way forever!"

"Well, we have a computer, right? It's pretty old, but we can install..."

"No, we can't! You know the Warehouse must remain top secret, don't you? We're not downloading or installing any software here."

"Okay, so you want us to write our own database implementation? Will that work?"

"Hmm, it MIGHT work..."

"Perfect! So I'll ask Myka to implement one for us!"

"Wait, I thought you were talking about implementing it yourself..."

"Nah, I'm not good at coding. Anyway, let's design it! What information should we store?"

"Each artifact will have its own unique ID, and then we need to store some metadata about it in a structured format. Also, everything should be accessible via a command line interface, because I don't understand these modern GUIs..."

One hour later...

"What? You want me to write a fully functional key-value store for working with JSON documents? From scratch?"

"I know, I know! But you're not alone! Here, I've made you some coffee in an Andy Warhol coffee mug, so it's pretty much and unlimited resource of caffeine superpower."

"But, Pete! We have several problems to solve! What if the data is corrupted? What if we can't access some artifact data when we need it most?"

"Don't worry, Myka, you can do this! I'll sit here and help, too. Let's just go through the problems one at a time."

<h2 id="chapter-iv" >Chapter IV</h2>
<h3 id="ex00">Task 00: Scalability</h3>

After some time, the board was covered with writings.

"Command line access — there should be a separate application that provides the REPL and connects to a running instance over the network, even if it's just a local host and port."

"We should be able to kill any instance (process) of the database and it should continue to run and provide answers to queries. This means that one of the configurable parameters, for example, should be a replication factor, i.e. how many copies of the same document we store. For testing purposes, 2 is probably enough."

"The client should perform heartbeats to check if the current database instance is available. If it stops responding, it should automatically switch to another instance."

"Also, for simplicity, let's assume for now that being scalable means that the client should be aware of all other nodes. Any heartbeat response from a current node should include all currently known instance addresses and ports along with the current replication factor."

So here we need to implement two programs — one is the client and one is an instance of a database. Whenever you start a new instance, you should be able to point it to an existing instance, so that after receiving a heartbeat it will send its host and port to all other running nodes, and everyone will know the new guy.

If the instance node is started with a different replication factor than existing nodes, it should detect this and automatically fail without joining the cluster. This means that the replication factor should probably be included in the heartbeat as well.

You can use any network protocol for this — HTTP, gRPC, etc.

Whenever the replication factor is more than a number of running nodes, information about this problem should be included in a heartbeat and explicitly displayed in each connected client. You can see an example of a user session in Task 01.

Actual work with documents will be implemented in the next task.

<h2 id="chapter-v" >Chapter V</h2>
<h3 id="ex01">Task 01: Balancing and Queries</h3>

"Okay, so let's use UUID4 strings as artifact keys. We also need to implement some balancing to provide fault tolerance..."

Our simple database should only support three operations — GET, SET and DELETE. 

Here's what a typical session should look like, with comments (starting with #):

```
~$ ./warehouse-cli -H 127.0.0.1 -P 8765
Connected to a database of Warehouse 13 at 127.0.0.1:8765
Known nodes:
127.0.0.1:8765
127.0.0.1:9876
127.0.0.1:8697
> SET 12345 '{"name": "Chapayev's Mustache comb"}'
Error: Key is not a proper UUID4
> SET 0d5d3807-5fbf-4228-a657-5a091c4e497f '{"name": "Chapayev's Mustache comb"}'
Created (2 replicas)
> GET 0d5d3807-5fbf-4228-a657-5a091c4e497f
'{"name": "Chapayev's Mustache comb"}'
> DELETE 0d5d3807-5fbf-4228-a657-5a091c4e497f
Deleted (2 replicas)
> GET 0d5d3807-5fbf-4228-a657-5a091c4e497f
Not found
>
# if current instance is stopped in the background
Reconnected to a database of Warehouse 13 at 127.0.0.1:8697
Known nodes:
127.0.0.1:9876
127.0.0.1:8697
> 
# if another current instance is stopped in the background
Reconnected to a database of Warehouse 13 at 127.0.0.1:9876
Known nodes:
127.0.0.1:9876
WARNING: cluster size (1) is smaller than a replication factor (2)!
>
```

If a key specified in SET already exists in a database, the value should be overwritten. If it doesn't, then the SET operation should provide read-after-write consistency, meaning that an immediate read should return the correct value.

When updating or deleting an existing value, eventual consistency should be implemented, meaning that immediate (dirty) reads may (but not "should"!) give you old results, but after a few seconds the data should be updated to a proper new state.

You can implement key-hash-based balancing, so your client can explicitly compute for each entry the list of nodes where it should be stored according to a replication factor. This is also useful for deletion.

If a current node is killed while writing, your client should automatically re-request to another available node. The only case where the user should see an error like "Failed to write/read an entry" is when ALL instances are dead.

<h2 id="chapter-vi" >Chapter VI</h2>
<h3 id="ex02">Task 02: Long Live the King</h3>

Let's update the logic from Tasks 00/01. We now introduce the concepts of a Leader and a Follower node. This leads to a list of important changes:

* From now on, the client ONLY interacts with a Leader node. The hash function to determine where to write replicas is now *on* the Leader, *not* in the client.
* All nodes (Leader and Follower) keep sending each other heartbeats with a full list of nodes. If a node doesn't respond to heartbeats for a certain configurable timeout (for testing purposes you should set it to 10 seconds by default).
* If the Leader is stopped, the remaining Followers should be able to choose a new Leader from among them. For simplicity, each of them can just sort the list of nodes by some other unique identifier (numeric id, port, etc.) and pick the top one. From that moment on, all heartbeats will contain a new elected Leader.
* If a client is unable to connect to a known Leader, they should try to connect to Followers to get a heartbeat from them. If a Leader is killed, this heartbeat will contain a new elected Leader.

<h2 id="chapter-vii" >Chapter VII</h2>
<h3 id="ex03">Task 03: Consensus</h3>

**NOTE: this task is completely optional. It is only graded as a bonus part**

You may have noticed that a lot of things can go wrong in a schema provided above, specifically race conditions and the possibility of losing some data because replicas are not automatically resynchronized between instances after some of them die.

You can try to fix this for some extra credit, either by using an existing solution or by writing a workaround yourself. Here are some options:

* Use an existing Raft implementation (https://github.com/hashicorp/raft) or write a minimal implementation yourself.
* Use external tools, such as Zookeeper (https://zookeeper.apache.org/) or Etcd (https://etcd.io/)
* Go some other way, like Paxos (https://github.com/kkdai/paxos), some blockchain implementations (like https://tendermint.com/), or your own hacks.

...Hopefully now Pete and Myka won't have to go through a pile of paperwork every time they need to find something. Probably Artie will do it anyway, because sometimes it's really hard to challenge the power of habit.

But I think it's been an interesting journey, and we've found some cool artifacts along the way. Do you?

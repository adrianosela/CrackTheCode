## Crack The Code

This purpose of this project was to experiment with go routines, channels, and select statements (on the client), mutex locks (for the resource on the server), and automating a build process using a Makefile, shell scripts and Docker containers.

### Build & Run!
Clone the repo:

```
git clone https://github.com/adrianosela/CrackTheCode
```

Spin up server container and build client:

```
make up
```

To stop the container and clean up Docker image:

```
make down
```

### CLI Commands:

When the server is started, a random code is set. To set a new one at any point and reset the number of tries recorded, use the command:

```
ctc code generate
```

To have a go at trying to guess the code:

```
ctc code crack --num [value]
```

Or to brute force the server's code:

```
ctc code crack --all
```

To see how many times you've tried to guess the code or after how many requests the brute force crack found the code:

```
ctc code tries
```

Finally, to cheat and see the secret code:

```
ctc code cheat
```


Usage Summary:

```
NAME:
   ctc code - manages the current code

USAGE:
   ctc code command [command options] [arguments...]

COMMANDS:
     generate  generates and sets a new code on the server
     crack     queries the server with a guess for the code
     cheat     gets the value of the current code
     tries     gets current number of tries to crack the code

OPTIONS:
   --help, -h  show help
```

   

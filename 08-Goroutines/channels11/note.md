# Note:
- OS threads has fixed stack size (64KB to 2MB)
- Goroutines starts with 2KB stack, it grows and shrinks as needed.

- Goroutines contains its own schedular, uses a technique known as m:n scheduling. 'm' goroutines scheduled on 'n' OS threads.

- Goroutines does not need a full kernel context switch for rescheduling.
package main

import (
	"embed"
	"fmt"
	"github.com/joho/godotenv"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

const PORT = "8080"

var quotes = []string{
	"Programs must be written for people to read, and only incidentally for machines to execute. — Harold Abelson",
	"Simplicity is the soul of efficiency. — Austin Freeman",
	"First, solve the problem. Then, write the code. — John Johnson",
	"Design is not just what it looks like and feels like. Design is how it works. — Steve Jobs",
	"Good software, like wine, takes time. — Joel Spolsky",
	"Make it work, make it right, make it fast. — Kent Beck",
	"Any fool can write code that a computer can understand. Good programmers write code that humans can understand. — Martin Fowler",
	"The best way to get a project done faster is to start sooner. — Jim Highsmith",
	"Walking on water and developing software from a specification are easy if both are frozen. — Edward V. Berard",
	"You can’t have great software without a great team, and most software teams behave like dysfunctional families. — Jim McCarthy",
	"A good programmer is someone who always looks both ways before crossing a one-way street. — Doug Linder",
	"Talk is cheap. Show me the code. — Linus Torvalds",
	"In software, the good stuff happens after the first draft is written. — Unknown",
	"The best error message is the one that never shows up. — Thomas Fuchs",
	"Fail often so you can succeed sooner. — Tom Kelley",
	"Simplicity is the ultimate sophistication. — Leonardo da Vinci",
	"It’s not a bug—it’s an undocumented feature. — Unknown",
	"A language that doesn’t affect the way you think about programming is not worth knowing. — Alan Perlis",
	"If you don’t have time to do it right, when will you have time to do it over? — John Wooden",
	"There are only two hard things in computer science: cache invalidation and naming things. — Phil Karlton",
	"When in doubt, use brute force. — Ken Thompson",
	"Software is a great combination of artistry and engineering. — Bill Gates",
	"Programming is breaking a problem down into small, manageable parts. — Unknown",
	"The computer is mightier than the pen. — Unknown",
	"An expert is a person who has made all the mistakes that can be made in a very narrow field. — Niels Bohr",
	"Code is like humor. When you have to explain it, it’s bad. — Cory House",
	"If you want to go fast, go alone. If you want to go far, go together. — African Proverb",
	"The most effective debugging tool is still careful thought, coupled with judiciously placed print statements. — Brian Kernighan",
	"Design is intelligence made visible. — Alina Wheeler",
	"It's not the tools you have faith in, tools are just tools. — Claude Monet",
	"Before software can be reusable it first has to be usable. — Ralph Johnson",
	"The best way to predict the future is to invent it. — Alan Kay",
	"Code is like a joke; if you have to explain it, it’s bad. — Unknown",
	"A problem well stated is a problem half solved. — Charles Kettering",
	"Fail fast, fail often. — Unknown",
	"Good code is its own best documentation. — Steve McConnell",
	"Life is too short to write bad code. — Unknown",
	"The only way to learn a new programming language is by writing programs in it. — Dennis Ritchie",
	"Software engineering is the art of balancing conflicting goals. — Unknown",
	"Any sufficiently advanced technology is indistinguishable from magic. — Arthur C. Clarke",
	"It’s not about ideas. It’s about making ideas happen. — Scott Belsky",
	"When the only tool you have is a hammer, everything looks like a nail. — Abraham Maslow",
	"The problem is not that we have problems. The problem is expecting otherwise and thinking that having problems is a problem. — Theodore Rubin",
	"First, do no harm. — Hippocrates",
	"We’ve all heard that in order to solve a problem, you need to understand it. But sometimes, you need to just do something and see what happens. — Unknown",
	"There’s always a better way. — Unknown",
	"The most important part of writing a program is not how it is written, but what it does. — Unknown",
	"Good programmers know what to write. Great programmers know what to rewrite (and reuse). — Unknown",
	"The best thing about a boolean is even if you are wrong, you are only off by a bit. — Unknown",
	"The only difference between a problem and a solution is that people understand the solution. — Unknown",
	"Time is money, and in programming, time equals code quality. — Unknown",
	"Success is walking from failure to failure with no loss of enthusiasm. — Winston Churchill",
	"Good design is as little design as possible. — Dieter Rams",
	"It’s not about the size of your code, it’s how you use it. — Unknown",
	"Great things are not done by impulse, but by a series of small things brought together. — Vincent van Gogh",
	"Learn the rules like a pro, so you can break them like an artist. — Pablo Picasso",
	"Don't worry about failure; you only have to be right once. — Drew Houston",
	"Simplicity is the ultimate sophistication. — Leonardo da Vinci",
	"The best way to predict the future is to create it. — Peter Drucker",
	"The problem is not that we have problems, it’s expecting otherwise. — Unknown",
	"Keep it simple, stupid. — Kelly Johnson",
	"The only way to make a great software product is by failing fast and often. — Unknown",
	"Code should be easy to understand, even for someone who has never written it before. — Unknown",
	"A designer knows he has achieved perfection not when there is nothing left to add, but when there is nothing left to take away. — Antoine de Saint-Exupery",
	"The key to successful programming is consistency. — Unknown",
	"Software isn’t finished when it’s perfect, it’s finished when it’s usable. — Unknown",
	"The function of good software is to make the complex appear to be simple. — Grady Booch",
	"The best way to design something is to figure out what the user wants before they even know it themselves. — Unknown",
	"To solve big problems, you need to tackle the small ones first. — Unknown",
	"If you’re not failing, you’re not trying. — Unknown",
	"A code without tests is like a car without brakes. — Unknown",
	"Failing is just another step on the road to success. — Unknown",
	"To simplify complexity, we must first understand it. — Unknown",
	"Good software design is about making trade-offs. — Unknown",
	"Developing software is like gardening. It’s not just about planting the seeds; it’s about nurturing them over time. — Unknown",
	"Everything should be made as simple as possible, but not simpler. — Albert Einstein",
	"Software is a tool to help people solve problems. — Unknown",
	"The art of programming is to organize complexity. — Unknown",
	"Code is poetry, not because it’s beautiful, but because it communicates with clarity and precision. — Unknown",
	"To design is to plan for the unknown. — Unknown",
	"You don’t write code for a machine; you write code for people. — Unknown",
	"Debugging is like being a detective in a criminal movie where you are also the murderer. — Filipe Fortes",
	"The best way to get something done is to begin. — Unknown",
	"Programming isn't about what you know; it's about what you can figure out. — Chris Pine",
	"An optimized code is never as readable as it should be. — Unknown",
	"The value of a great team is much higher than the value of a single expert. — Unknown",
	"A good system has more to do with how well the team collaborates than how well it is designed. — Unknown",
	"Every line of code you write should have a reason behind it. — Unknown",
	"Design is thinking made visual. — Saul Bass",
	"In software development, you’re never truly finished. — Unknown",
	"A programmer is a person who solves a problem you didn’t know you had in a way you don’t understand. — Unknown",
	"The key to solving any problem is understanding the problem. — Unknown",
	"It’s not about how much you know, it’s about how much you apply what you know. — Unknown",
	"You only need to know one thing to be successful in programming: the solution. — Unknown",
	"When code is clean, it’s beautiful. — Unknown",
	"A system that doesn’t work is worse than a system that’s broken. — Unknown",
	"The best software design makes the difficult look easy. — Unknown",
	"Innovation distinguishes between a leader and a follower. — Steve Jobs",
}

//go:embed index.html
var index embed.FS

func main() {
	_ = godotenv.Load()
	port := os.Getenv("PORT_GO_SSE_TUTORIAL")
	if port == "" {
		log.Fatal("PORT_GO_SSE_TUTORIAL environment variable is required.")
	}
	addr := "localhost:" + port

	indexFS, _ := fs.Sub(index, ".")

	http.Handle("/", http.FileServer(http.FS(indexFS)))
	http.HandleFunc("/quotes", eventsHandler)

	log.Println("Server listening on http://" + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("error starting server at " + addr)
	}
}

func eventsHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	for _, quote := range quotes {
		_, err := fmt.Fprintf(w, "data: %s\n\n", quote)
		if err != nil {
			continue
		}
		w.(http.Flusher).Flush()
		time.Sleep(5 * time.Second)
	}
}

package main

import "fmt"

type Message struct {
	From    string
	Payload string
}

type server struct {
	servChan chan Message
}

func (s *server) listenAndServe() {
	for msg := range s.servChan {
		fmt.Printf("from: %s, payload: %s\n", msg.From, msg.Payload)
	}

}
func sendMessageToTheServer(message string, s chan Message) {
	m := Message{
		From:    "haleem",
		Payload: message,
	}
	s <- m

}

func main() {
	serverChan := make(chan Message)
	s := server{
		servChan: serverChan,
	}
	go s.listenAndServe()
	sendMessageToTheServer("halemo is here", serverChan)

	select {}
}

// func getUserData(userId int, resChan chan string, resWg *sync.WaitGroup) {
// 	time.Sleep(80 * time.Millisecond)
// 	fmt.Printf("getting user data of %v \n", userId)
// 	resChan <- "mahmoud"
// 	resWg.Done()
// }
// func getAiRecommendations(userId int, resChan chan string, resWg *sync.WaitGroup) {
// 	time.Sleep(120 * time.Millisecond)
// 	fmt.Printf("getting user data of %v \n", userId)
// 	resChan <- "Google"
// 	resWg.Done()
// }
// func getuserLikes(userId int, resChan chan string, resWg *sync.WaitGroup) {
// 	time.Sleep(60 * time.Millisecond)
// 	fmt.Printf("get user data ...,%v \n", userId)
// 	resChan <- "likeeeeeeeees"
// 	resWg.Done()
// }
// func main() {
// 	now := time.Now()
// 	userId := 10
// 	resWg := sync.WaitGroup{}
// 	resChan := make(chan string, 120)
// 	resWg.Add(3)
// 	go getUserData(userId, resChan, &resWg)
// 	go getAiRecommendations(userId, resChan, &resWg)
// 	go getuserLikes(userId, resChan, &resWg)

// 	go func() {
// 		resWg.Wait()
// 		close(resChan)
// 	}()

// 	for res := range resChan {
// 		fmt.Println("Result:", res)
// 	}

// 	fmt.Println("Total time:", time.Since(now))
// }

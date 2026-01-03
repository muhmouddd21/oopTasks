package main

import (
	"context"
	"fmt"
)

// ------------------------------------------------------------------
// cancel a request using context
// ------------------------------------------------------------------
// type Respose struct {
// 	value int
// 	err   error
// }

// func fetchUserData(ctx context.Context, userid int) (int, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
// 	defer cancel()
// 	respch := make(chan Respose)
// 	go func() {
// 		val, err := fetchThirdPartyApi()
// 		respch <- Respose{
// 			value: val,
// 			err:   err,
// 		}
// 	}()
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return 0, fmt.Errorf("fetching data from third party took too long")
// 		case res := <-respch:
// 			return res.value, res.err
// 		}
// 	}
// }

// func fetchThirdPartyApi() (int, error) {
// 	time.Sleep(time.Millisecond * 500)
// 	return 666, nil
// }

// func main() {
// 	start := time.Now()
// 	ctx := context.Background()
// 	userId := 10
// 	val, err := fetchUserData(ctx, userId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(val)
// 	fmt.Println(time.Since(start))
// }

// --------------------------------
type Message struct {
	From    string
	Payload string
}

type server struct {
	servChan chan Message
}

func (s *server) listenAndServe(ctx context.Context) {
	for {
		select {
		case msg := <-s.servChan:
			fmt.Println(msg)
		case <-ctx.Done():
			return
		}
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
	ctx := context.Background()
	serverChan := make(chan Message)
	s := server{
		servChan: serverChan,
	}
	go s.listenAndServe(ctx)
	sendMessageToTheServer("halemo is here", serverChan)

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

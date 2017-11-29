package ai

func Start(comChan chan string) {
	comChan <- "YES"
}

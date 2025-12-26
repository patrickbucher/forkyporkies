package main

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	fmt.Println(token)
}

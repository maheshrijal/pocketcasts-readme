package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

type AuthResponse struct {
	Token string `json:"token"`
}

type Episode struct {
	UUID          string   `json:"uuid"`
	URL           string   `json:"url"`
	Published     string   `json:"published"`
	Duration      int      `json:"duration"`
	FileType      string   `json:"fileType"`
	Title         string   `json:"title"`
	Size          string   `json:"size"`
	PlayingStatus int      `json:"playingStatus"`
	PlayedUpTo    int      `json:"playedUpTo"`
	Starred       bool     `json:"starred"`
	PodcastUUID   string   `json:"podcastUuid"`
	PodcastTitle  string   `json:"podcastTitle"`
	EpisodeType   string   `json:"episodeType"`
	EpisodeSeason int      `json:"episodeSeason"`
	EpisodeNumber int      `json:"episodeNumber"`
	IsDeleted     bool     `json:"isDeleted"`
	Author        string   `json:"author"`
	Bookmarks     []string `json:"bookmarks"`
}

type Data struct {
	Total    int       `json:"total"`
	Episodes []Episode `json:"episodes"`
}

func main() {
	emailFlag := flag.String("email", "", "Pocket Casts login email")
	passwordFlag := flag.String("password", "", "Pocket Casts login password")
	flag.Parse()

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	if *emailFlag == "" || *passwordFlag == "" {
		fmt.Println("Email & Password are mandatory!")
		flag.PrintDefaults()
		return
	}

	if !isValidEmail(*emailFlag, emailRegex) {
		fmt.Println("Invalid email format.")
		return
	}

	token, err := authenticate(*emailFlag, *passwordFlag)
	if err != nil {
		log.Fatalf("Authentication error: %s", err)
		return
	}

	fmt.Println("Fetching listening history.")

	historyData, err := fetchUserHistory(token)
	if err != nil {
		log.Fatalf("Error fetching listen history: %s", err)
		return
	}

	err = ParseJson(string(historyData))
	if err != nil {
		fmt.Println(string(historyData))
		log.Fatalf("Failed parsing history data: %s", err)
	}
}

func isValidEmail(email, regex string) bool {
	match, err := regexp.MatchString(regex, email)
	if err != nil {
		log.Fatal("Failed to match Email regex.")
	}
	return match
}

func authenticate(email, password string) (string, error) {
	authUrl := "https://api.pocketcasts.com/user/login"
	payload := map[string]string{
		"email":    email,
		"password": password,
		"scope":    "webplayer",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "Failed creating payload!", err
	}

	req, err := http.NewRequest("POST", authUrl, bytes.NewReader(payloadBytes))
	if err != nil {
		return "Faled to make request!", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Error sending request:", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("authentication request failed with status code: %s", resp.Status)
	}

	var authResponse AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return "Failed decode response: ", err
	}
	return authResponse.Token, nil
}

func fetchUserHistory(token string) ([]byte, error) {
	historyUrl := "https://api.pocketcasts.com/user/history"

	historyReq, err := http.NewRequest("POST", historyUrl, nil)
	if err != nil {
		return nil, err
	}
	historyReq.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	historyResp, err := client.Do(historyReq)
	if err != nil {
		return nil, err
	}
	defer historyResp.Body.Close()

	if historyResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("user history request failed with status code: %s", historyResp.Status)
	}

	historyData, err := io.ReadAll(historyResp.Body)
	if err != nil {
		return nil, err
	}

	return historyData, nil
}

func ParseJson(historydata string) error {
	var data Data
	err := json.Unmarshal([]byte(historydata), &data)
	if err != nil {
		log.Fatal("Unable to marshall json")
	}

	// Get the top 5 episodes
	topEpisodes := data.Episodes[:5]

	var top5content string

	// Write the top 5 episodes to file
	for _, episode := range topEpisodes {
		top5content += fmt.Sprintf("%s: [%s](%s)\n\n", episode.PodcastTitle, episode.Title, episode.URL)
	}

	fmt.Printf("Printing top5 Conent: \n%s\n", top5content)

	readmePath := "README.md"
	readmeContent, err := os.ReadFile(readmePath)
	if err != nil {
		log.Fatalf("Error reading README.md: %v", err)
		return err
	}

	// Define the section markers
	startMarker := "<!-- PODCAST-LIST:START -->"
	endMarker := "<!-- PODCAST-LIST:END -->"

	// Create a regular expression pattern that matches the content between the markers
	pattern := regexp.MustCompile(startMarker + `[\s\S]*?` + endMarker)

	//Replace content
	// updatedContent := strings.Replace(string(readmeContent), startMarker+"\n"+"<!-- PODCAST-LIST:END -->", startMarker+"\n"+top5content+"\n"+endMarker, -1)

	// Replace the content between markers with the new content
	updatedContent := pattern.ReplaceAllString(string(readmeContent), startMarker+"\n"+top5content+"\n"+endMarker)

	// Write the updated content back to the README.md file
	err = os.WriteFile(readmePath, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Println("Error writing to README.md:", err)
		return err
	}

	fmt.Println("README.md updated successfully.")
	return nil
}

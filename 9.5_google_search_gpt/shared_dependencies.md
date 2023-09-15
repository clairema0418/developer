the app is: Golang package integrating Google Search API with ChatGPT

the files we have decided to generate are: main.go, google_search.go, chatgpt.go, config.go

Shared dependencies:
1. Package imports:
   - "net/http"
   - "encoding/json"
   - "io/ioutil"
   - "log"
   - "os"
   - "github.com/joho/godotenv"

2. API keys and configuration:
   - GoogleSearchAPIKey (string)
   - ChatGPTAPIKey (string)
   - GoogleSearchAPIEndpoint (string)
   - ChatGPTAPIEndpoint (string)

3. Data schemas:
   - GoogleSearchResult (struct)
   - ChatGPTRequest (struct)
   - ChatGPTResponse (struct)

4. Function names:
   - main()
   - loadConfig()
   - googleSearch(query string) ([]GoogleSearchResult, error)
   - chatGPTQuery(prompt string) (string, error)
   - handleUserQuestion(w http.ResponseWriter, r *http.Request)
package azure_tts

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/texttospeech"
	"github.com/Azure/go-autorest/autorest"
)

const (
	azureTTSKey      = "your_azure_tts_key"
	azureTTSEndpoint = "https://your_region.tts.speech.microsoft.com/cognitiveservices/v1"
)

func ConvertAudioToText(audioFilePath string) (string, error) {
	client := texttospeech.NewSpeechSynthesizerClient(azureTTSEndpoint)
	client.Authorizer = autorest.NewCognitiveServicesAuthorizer(azureTTSKey)

	audioBytes, err := ioutil.ReadFile(audioFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read audio file: %v", err)
	}

	ctx := context.Background()
	response, err := client.SynthesizeSpeech(ctx, "application/ssml+xml", "audio/wav", "en-US", "Microsoft.ServerSpeech.TextToSpeech.Voice.en-US.Guy24KRUS", string(audioBytes))
	if err != nil {
		return "", fmt.Errorf("failed to synthesize speech: %v", err)
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	text := result["DisplayText"].(string)
	return text, nil
}

func GetToken() (string, error) {
	req, err := http.NewRequest("POST", "https://your_region.api.cognitive.microsoft.com/sts/v1.0/issueToken", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %v", err)
	}

	req.Header.Set("Ocp-Apim-Subscription-Key", azureTTSKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get token: %v", err)
	}

	defer resp.Body.Close()
	tokenBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read token response: %v", err)
	}

	token := strings.TrimSpace(string(tokenBytes))
	return token, nil
}
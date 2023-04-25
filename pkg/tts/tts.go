package tts

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const (
	ttsApi    = "https://southeastasia.tts.speech.microsoft.com/cognitiveservices/v1"
	accessKey = "5e2b9452dc09416981e5b9c45b0124b6"
)

func TextToSpeech(plaintext, path string, part int, langKaz bool) (string, error) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}

	var lang string
	var gender string
	var voice string

	if langKaz {
		lang = "kk-KZ"
		gender = "Female"
		voice = "kk-KZ-AigulNeural"
	} else {
		lang = "ru-RU"
		gender = "Male"
		voice = "ru-RU-DmitryNeural"
	}
	client := http.DefaultClient

	ssml := fmt.Sprintf(`<speak version='1.0' xml:lang='%v'>
	<voice xml:lang='%v' xml:gender='%v' name='%v'>
	%v
	</voice></speak>`, lang, lang, gender, voice, plaintext)

	body := bytes.NewBuffer([]byte(ssml))
	req, err := http.NewRequest(http.MethodPost, ttsApi, body)
	if err != nil {
		return "", fmt.Errorf("cannot create tts request: %v", err)
	}
	req.Header.Add("Ocp-Apim-Subscription-Key", accessKey)
	req.Header.Add("Content-Type", "application/ssml+xml")
	req.Header.Add("X-Microsoft-OutputFormat", "audio-16khz-32kbitrate-mono-mp3")
	req.Header.Add("User-Agent", "tts-test")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send requet: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to request tts, status: %v", resp.Status)
	}

	filename := path + lang + "." + strconv.Itoa(part) + ".mp3"

	out, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("cannot create output file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("cant copy: %w", err)
	}

	return filename, nil
}

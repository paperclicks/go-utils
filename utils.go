package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//GetFileWriter gets a filename as input, and if the file exists returns a writer to the file.
//If the file does not exist creates a new file and returns the writer
func GetFileWriter(fileName string) (io.Writer, error) {

	//check if file exists
	var _, err = os.Stat(fileName)

	// create file if not exists
	if os.IsNotExist(err) {
		out, err := os.Create(fileName)
		if err != nil {

			return os.Stdout, err
		}
		defer out.Close()
	}

	out, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0664)

	if err != nil {
		return os.Stdout, err
	}

	return out, err
}

//IsValidWidget checks if a widget value is empty or equal to a traffic source token
//(meaning that the widget was not replaced correctly)
func IsValidWidget(widgetID string) bool {

	if widgetID == "" {
		return false
	}

	if widgetID == " " {
		return false
	}

	//if widget starts with { or ends with } means that the token was not replaced successfully
	if strings.HasPrefix(widgetID, "{") {
		return false
	}

	if strings.HasSuffix(widgetID, "}") {
		return false
	}

	//Outbrain: if widget starts with $ or ends with $ means that the token was not replaced successfully
	if strings.HasPrefix(widgetID, "$") {
		return false
	}

	if strings.HasSuffix(widgetID, "$") {
		return false
	}

	//ContentAd:if widget starts with [] or ends with ] means that the token was not replaced successfully
	if strings.HasPrefix(widgetID, "[") {
		return false
	}

	if strings.HasSuffix(widgetID, "]") {
		return false
	}
	if strings.HasSuffix(widgetID, "s0") {
		return false
	}

	return true
}

//sanitizeString takes in input a string and removes possible starting/ending whitespaces and special chars as /n, /r, /t
func SanitizeString(str string) string {
	//remove whitespaces
	s := strings.TrimSpace(str)

	//remove /t /r /n
	s = strings.TrimSuffix(s, "/t")
	s = strings.TrimSuffix(s, "/n")
	s = strings.TrimSuffix(s, "/r")

	s = strings.TrimPrefix(s, "/t")
	s = strings.TrimPrefix(s, "/n")
	s = strings.TrimPrefix(s, "/r")

	return s
}

//GetIntegerEnv is a helper to get an int ENV value, or fallback to a default value if that env is not set
func GetIntegerEnv(env string, fallback int) int {

	//try to parse the value from the given env variable
	value, err := strconv.Atoi(os.Getenv(env))

	if err != nil {
		value = fallback
	}

	return value
}

//GetStringEnv is a helper to get a string ENV value, or fallback to a default value if that ENV is not set
func GetStringEnv(env string, fallback string) string {

	//try to parse the value from the given env variable
	value, ok := os.LookupEnv(env)

	if !ok {

		value = fallback
	}

	return value
}

//FileContainsString takes in input a string and file path and checks if the string is present if the file content
func FileContainsString(s string, file string) bool {

	// read the file
	buff, err := ioutil.ReadFile(file)
	if err != nil {
		return false
	}

	content := string(buff)

	// //check whether s contains substring text
	if strings.Contains(content, s) {
		return true
	}

	return false
}

func QueueName(name string) string {

	//if ENVIRONMENT=local then use "local" as a postfix for queues
	postfix := ""
	if os.Getenv("ENVIRONMENT") == "local" {
		postfix = ".local"
	}

	queueName := fmt.Sprintf("%s%s", name, postfix)

	return queueName
}

package tension

import "net/http"
import "net/url"
import "fmt"
import "bytes"
import "encoding/json"

type Slack struct {
	token         string
	userNameCache map[UserId]string
}

func New(token string) (s *Slack) {
	s = &Slack{token: token}
	return
}

func (s *Slack) call(endpoint string, params *map[string]string, d interface{}) (err error) {
	var u url.URL
	u.Scheme = "https"
	u.Host = "slack.com"
	u.Path = fmt.Sprintf("api/%s", endpoint)

	q := url.Values{"token": []string{s.token}}
	if params != nil {
		for paramKey, paramValue := range *params {
			q.Set(paramKey, paramValue)
		}
	}
	u.RawQuery = q.Encode()

	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	rb := bytes.Buffer{}
	rb.ReadFrom(response.Body)
	err = json.Unmarshal(rb.Bytes(), &d)
	return
}

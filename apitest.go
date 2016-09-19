package tension

import "fmt"

// APITest issues a call against the Slack API, which you can optionally force
// to return unsuccessfully.
func (s *Slack) APITest(fail bool) (res *APITestResult, err error) {
	res = &APITestResult{}

	p := map[string]string{}

	if fail {
		p["error"] = "tension_forced_api_error"
	}

	err = s.call("api.test", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

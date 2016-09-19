package tension

import "fmt"

// AuthTest will ask the API whether our authentication token is valid, and
// if so, give us some boilerplate information about the token
func (s *Slack) AuthTest() (res *AuthTestResult, err error) {
	res = &AuthTestResult{}
	err = s.call("auth.test", nil, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

// AuthRevoke will cause the current authentication token associated with `s`
// to be revoked, for instance if the user is logging out. Don't use this with
// bot tokens. Set `test` to exercise the API without really revoking anything.
func (s *Slack) AuthRevoke(test bool) (res *AuthRevokeResult, err error) {
	res = &AuthRevokeResult{}
	p := map[string]string{}

	if test {
		p["test"] = "1"
	}

	err = s.call("auth.revoke", &p, res)
	if err != nil {
		return
	}
	if res.Ok == false {
		err = fmt.Errorf("Slack API error: %s", res.ErrorMsg)
	}
	return
}

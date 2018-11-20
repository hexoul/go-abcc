package abcc

import (
	"encoding/json"
	"fmt"

	"github.com/hexoul/go-abcc/types"
)

// Me obtains your own personal asset information
//   arg: -
//   src: https://api.abcc.com/api/v1/members/me
//   doc: -
func (s *Client) Me(options *types.Options) (*types.UserInfo, error) {
	url := fmt.Sprintf("%s/members/me?%s", baseURL, s.parseOptions("/api/v1/members/me", options))

	body, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	var result = new(types.UserInfo)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}
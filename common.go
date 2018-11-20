package abcc

import (
	"encoding/json"
	"fmt"

	"github.com/hexoul/go-abcc/types"
)

// Timestamp returns server time
//   arg: -
//   src: https://api.abcc.com/api/v1/common/timestamp
//   doc: -
func (s *Client) Timestamp(options *types.Options) (*types.Timestamp, error) {
	url := fmt.Sprintf("%s/common/timestamp", baseURL)

	body, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	var result = new(types.Timestamp)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

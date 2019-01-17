package abcc

import (
	"encoding/json"
	"fmt"

	"github.com/hexoul/go-abcc/types"
)

// Orders list
//   arg: -
//   src: https://api.abcc.com/api/v1/exchange/orders
//   doc: -
func (s *Client) Orders(options *types.Options) (*types.Orders, error) {
	url := fmt.Sprintf("%s/exchange/orders?%s", baseURL, s.parseOptions("/api/v1/exchange/orders", s.fillTimestampFromServer(options)))

	body, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	var result = new(types.Orders)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

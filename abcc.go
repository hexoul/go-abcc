// Package abcc is an API Client for ABCC
package abcc

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hexoul/go-abcc/types"
	"github.com/hexoul/go-abcc/util"
)

// Interface for APIs
type Interface interface {
	Me(options *types.Options) (*types.UserInfo, error)
}

// Client for ABCC API
type Client struct {
	accessKey string
	secretKey string
}

var (
	instance  *Client
	once      sync.Once
	accessKey string
	secretKey string
)

const (
	baseURL = "https://api.abcc.com/api/v1"
)

func init() {
	for _, val := range os.Args {
		arg := strings.Split(val, "=")
		if len(arg) < 2 {
			continue
		} else if arg[0] == "-abcc:accesskey" {
			accessKey = arg[1]
		} else if arg[0] == "-abcc:secretkey" {
			secretKey = arg[1]
		}
	}
}

// GetInstance returns singleton
func GetInstance() *Client {
	once.Do(func() {
		if accessKey == "" || secretKey == "" {
			panic("KEYS FOR BOTH ACCESS AND SECRET REQUIRED")
		}
		instance = &Client{
			accessKey: accessKey,
			secretKey: secretKey,
		}
	})
	return instance
}

// GetInstanceWithKey returns singleton
func GetInstanceWithKey(accessKey, secretKey string) *Client {
	once.Do(func() {
		if accessKey == "" || secretKey == "" {
			panic("KEYS FOR BOTH ACCESS AND SECRET REQUIRED")
		}
		instance = &Client{
			accessKey: accessKey,
			secretKey: secretKey,
		}
	})
	return instance
}

func (s *Client) parseOptions(endpoint string, options *types.Options) string {
	// Make params
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	// NOTE: params should be sorted alphabetically
	params := "access_key=" + s.accessKey + "&tonce=" + timestamp
	if options == nil {
		options = &types.Options{}
	}
	if options.Symbol != "" {
		params = params + "&symbol=" + options.Symbol
	}

	// Sign
	msg := "GET" + endpoint + "?" + params
	h := hmac.New(sha256.New, []byte(s.secretKey))
	h.Write([]byte(msg))
	sign := hex.EncodeToString(h.Sum(nil))

	return params + "&signature=" + sign
}

func (s *Client) getResponse(url string) ([]byte, error) {
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	body, err := util.DoReq(httpReq)
	if err != nil {
		return nil, err
	}
	resp := new(types.Response)
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, fmt.Errorf("[%d] %s:%s", resp.Error.Code, resp.Error.Message, resp.Error.Reason)
	}
	return body, nil
}

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
// Package abcc is an API Client for ABCC
package abcc

import (
	"bytes"
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
	baseURL = "https://api.abcc.com/v1"
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

func (s *Client) parseOptions(endpoint string, options *types.Options) *types.Request {
	// Make params
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	params := "accesskey=" + s.accessKey + "&tonce=" + timestamp
	if options == nil {
		options = &types.Options{}
	}
	if options.Symbol != "" {
		params = params + "&symbol=" + options.Symbol
	}

	// Sign
	msg := "POST" + endpoint + "?" + params
	h := hmac.New(sha256.New, []byte(s.secretKey))
	h.Write([]byte(msg))
	sign := hex.EncodeToString(h.Sum(nil))

	// Make request
	return &types.Request{
		AccessKey: s.accessKey,
		Timestamp: timestamp,
		Params:    params,
		Sign:      sign,
	}
}

func (s *Client) getResponse(url string, req *types.Request) ([]byte, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBufferString(string(reqBody)))
	if err != nil {
		return nil, err
	}
	body, err := util.DoReq(httpReq)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Me obtains your own personal asset information
//   arg: -
//   src: https://api.abcc.com/v1/members/me
//   doc: -
func (s *Client) Me(options *types.Options) (*types.UserInfo, error) {
	url := fmt.Sprintf("%s/members/me", baseURL)

	body, err := s.getResponse(url, s.parseOptions("/api/v1/members/me", options))
	if err != nil {
		return nil, err
	}

	var result = new(types.UserInfo)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

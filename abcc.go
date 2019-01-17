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
	Timestamp(options *types.Options) (*types.Timestamp, error)
	Markets(options *types.Options) (*types.Markets, error)

	Orders(options *types.Options) (*types.Orders, error)

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

func (s *Client) fillTimestampFromServer(options *types.Options) *types.Options {
	timestamp, err := s.Timestamp(nil)
	if err != nil {
		return nil
	}
	if options == nil {
		options = &types.Options{}
	}
	options.Timestamp = strconv.FormatFloat(timestamp.Timestamp, 'f', -1, 64)
	return options
}

func (s *Client) parseOptions(endpoint string, options *types.Options) string {
	// Make params
	params := []string{}
	if options == nil {
		options = &types.Options{}
	}
	options.AccessKey = s.accessKey
	if options.Timestamp == "" {
		options.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	}
	if bOption, err := json.Marshal(options); err == nil {
		mOption := new(map[string]interface{})
		if err := json.Unmarshal(bOption, &mOption); err == nil {
			for _, k := range types.OptionOrder {
				if (*mOption)[k] != nil {
					params = append(params, fmt.Sprintf("%s=%v", k, (*mOption)[k]))
				}
			}
		}
	}
	options.AccessKey = ""
	options.Timestamp = ""
	sParams := strings.Join(params, "&")

	// Sign
	msg := "GET|" + endpoint + "|" + sParams
	h := hmac.New(sha256.New, []byte(s.secretKey))
	h.Write([]byte(msg))
	sign := hex.EncodeToString(h.Sum(nil))

	return sParams + "&signature=" + sign
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

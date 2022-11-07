package gt3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/nilorg/geetest/pkg/util"
)

// Client gt3客户端
type Client struct {
	geetestID   string // 公钥
	geetestKey  string // 私钥
	SdkVersion  string
	EesponseStr string
	opts        *ClientOptions
	httpClient  *http.Client
}

// NewClient 创建客户端
func NewClient(geetestID string, geetestKey string, opts ...Option) *Client {
	o := newOptions(opts...)
	return &Client{
		geetestID:   geetestID,
		geetestKey:  geetestKey,
		SdkVersion:  VERSION,
		opts:        &o,
		EesponseStr: "",
		httpClient:  &http.Client{Timeout: time.Duration(o.HTTPTimeout) * time.Second},
	}
}

func (g *Client) PreProcess(user_id int) (int, string) {
	new_captcha := 1
	JSON_FORMAT := 1
	client_type := ClientTypeWeb
	ip_address := ""
	status, challenge := g.register(user_id, new_captcha, JSON_FORMAT, client_type, ip_address)
	return status, challenge
}

func (g *Client) buildRequestComm(userID string) *RequestComm {
	return &RequestComm{
		UserID:     userID,
		ClientType: g.opts.ClientType,
		IPAddress:  g.opts.IPAddress,
		JSONFormat: g.opts.JSONFormat,
		Sdk:        g.opts.Version,
	}
}

// httpGet 发送GET请求，获取服务器返回结果
func (g *Client) httpGet(uri string, params map[string]string) (body []byte, err error) {
	q := url.Values{}
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
	}
	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s?%s", g.opts.APIURL, uri, q.Encode()), nil)
	if err != nil {
		return
	}
	var res *http.Response
	res, err = g.httpClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		err = fmt.Errorf("response status code: %d", res.StatusCode)
		return
	}
	body, err = ioutil.ReadAll(res.Body)
	return
}

// httpPost 发送POST请求，获取服务器返回结果
func (g *Client) httpPost(uri string, params map[string]string) (body []byte, err error) {
	q := url.Values{}
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
	}
	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s", g.opts.APIURL, uri), strings.NewReader(q.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	var res *http.Response
	res, err = g.httpClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("response status code: %d", res.StatusCode)
		return
	}
	body, err = ioutil.ReadAll(res.Body)
	return
}

// Register register
func (g *Client) register(UserID int, new_captcha int, JSON_FORMAT int, client_type string, ip_address string) (int, string) {
	challenge1, err := g.register_challenge(UserID, new_captcha, JSON_FORMAT, client_type, ip_address)
	if err != nil {
		return 1, ""
	}
	var challenge string
	if len(challenge1) == 32 {
		challenge = util.MD5Encode(challenge1 + g.geetestKey)
		print("challenge1", challenge)
		return 0, challenge
	} else {
		return 1, ""
	}
}

func (g *Client) register_challenge(user_id int, new_captcha int, JSON_FORMAT int, client_type string, ip_address string) (string, error) {
	var register_url string
	if user_id != 0 {
		register_url = fmt.Sprintf("%s%s?gt=%s&user_id=%d&json_format=%d&client_type=%s&ip_address=%s", g.opts.APIURL, g.opts.RegisterURL, g.geetestID, user_id, JSON_FORMAT, client_type, ip_address)
	} else {
		return "", fmt.Errorf("参数错误")
	}
	req, err := http.NewRequest("GET", register_url, nil)

	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	data, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	body, err := ioutil.ReadAll(data.Body)
	if err != nil {
		return "", err
	}

	resp := make(map[string]interface{})

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return "", err
	}
	fmt.Println("resp", resp)
	res_string := resp["challenge"].(string)
	fmt.Println("res_string", res_string)
	return res_string, nil
}

// Validate validate
func (g *Client) Validate(challenge, seccode string, userID ...string) (res *ValidateResponse, err error) {
	var reqComm *RequestComm
	if len(userID) > 0 {
		reqComm = g.buildRequestComm(userID[0])
	} else {
		reqComm = g.buildRequestComm("")
	}
	req := &ValidateRequest{
		RequestComm: reqComm,
		CaptchaID:   g.geetestID,
		Seccode:     seccode,
		Challenge:   challenge,
	}
	params := util.StructToMap(req)
	var body []byte
	body, err = g.httpPost(g.opts.ValidteURL, params)
	if err != nil {
		return
	}
	res = new(ValidateResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		res = nil
	}
	return
}

// BuildChallenge 构建验证初始化返回数据
func (g *Client) BuildChallenge(challenge string, digestmod string) (enchallenge string) {
	// challenge 为空或者值为0代表失败
	if challenge == "" || challenge == "0" {
		// 本地随机生成32位字符串
		characters := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
		challenge := []byte{}
		for i := 0; i < 32; i++ {
			challenge = append(challenge, characters[rand.Intn(len(characters))])
		}
		enchallenge = string(challenge)
	} else {
		if digestmod == "md5" {
			enchallenge = util.MD5Encode(challenge + g.geetestKey)
		} else if digestmod == "sha256" {
			enchallenge = util.Sha256Encode(challenge + g.geetestKey)
		} else if digestmod == "hmac-sha256" {
			enchallenge = util.HmacSha256Encode(challenge, g.geetestKey)
		} else {
			enchallenge = util.MD5Encode(challenge + g.geetestKey)
		}
	}
	return
}

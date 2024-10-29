package link

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

// ClientConfig 定义客户端配置
type ClientConfig struct {
	BaseURL             string
	MaxIdleConns        int           // 连接池中的最大空闲连接数
	MaxIdleConnsPerHost int           // 每个主机的最大空闲连接数
	MaxConnsPerHost     int           // 每个主机的最大连接数
	IdleConnTimeout     time.Duration // 空闲连接的超时时间
	DialTimeout         time.Duration // 建立连接的超时时间
	RequestTimeout      time.Duration // 请求的超时时间
	KeepAlive           time.Duration // TCP KeepAlive 时间
}

// Response 定义API响应的结构
type Response struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Data    string `json:"data"`
	Msg     string `json:"msg"`
}

// Client 定义API客户端
type Client struct {
	config     *ClientConfig
	httpClient *http.Client
}

// DefaultConfig 返回默认配置
func DefaultConfig() *ClientConfig {
	return &ClientConfig{
		BaseURL:             "https://link-system.rockorca.com",
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		MaxConnsPerHost:     100,
		IdleConnTimeout:     90 * time.Second,
		DialTimeout:         5 * time.Second,
		RequestTimeout:      10 * time.Second,
		KeepAlive:           30 * time.Second,
	}
}

// NewClient 创建新的客户端实例
func NewClient(config *ClientConfig) *Client {
	if config == nil {
		config = DefaultConfig()
	}

	// 创建自定义传输层
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   config.DialTimeout,
			KeepAlive: config.KeepAlive,
		}).DialContext,
		MaxIdleConns:        config.MaxIdleConns,
		MaxIdleConnsPerHost: config.MaxIdleConnsPerHost,
		MaxConnsPerHost:     config.MaxConnsPerHost,
		IdleConnTimeout:     config.IdleConnTimeout,
	}

	// 创建 HTTP 客户端
	httpClient := &http.Client{
		Transport: transport,
		Timeout:   config.RequestTimeout,
	}

	return &Client{
		config:     config,
		httpClient: httpClient,
	}
}

// GetClickLog 获取点击链接
func (c *Client) GetClickLog(channelID, userAgent, action string) (*Response, error) {
	// 构建请求URL
	endpoint := "/api/yl/link/click-log"
	params := url.Values{}
	params.Add("channelId", channelID)
	params.Add("userAgent", userAgent)
	params.Add("action", action)

	reqURL := fmt.Sprintf("%s%s?%s", c.config.BaseURL, endpoint, params.Encode())

	// 创建请求
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	// 添加请求头
	req.Header.Set("User-Agent", "LinkSystem-Client/1.0")
	req.Header.Set("Accept", "application/json")

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析JSON响应
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	// 检查响应状态
	if !response.Success || response.Code != 200 {
		return nil, fmt.Errorf("API返回错误: %s", response.Msg)
	}

	return &response, nil
}

// Close 关闭客户端，清理资源
func (c *Client) Close() {
	if transport, ok := c.httpClient.Transport.(*http.Transport); ok {
		transport.CloseIdleConnections()
	}
}

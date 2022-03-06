package rest

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"net/url"
	"time"
)

// used for JSON unmarshaling
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (c *Client) request(req Requester, results interface{}) error {
	res, reqErr := c.do(req)
	if reqErr != nil {
		return reqErr
	}

	decErr := decode(res, results)
	if decErr != nil {
		return decErr
	}
	return nil
}

func (c *Client) do(r Requester) (*fasthttp.Response, error) {
	req := c.newRequest(r)

	// fasthttp for http2.0
	res := fasthttp.AcquireResponse()
	err := c.HTTPC.DoTimeout(req, res, c.HTTPTimeout)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%+v\n", string(res.Body()))
	// no usefull headers

	if res.StatusCode() != 200 {
		var resp = new(Response)
		if jsonErr := json.Unmarshal(res.Body(), resp); jsonErr != nil {
			return nil, &APIError{
				Status:  res.StatusCode(),
				Message: jsonErr.Error(),
			}
		}

		if !resp.Success {
			return nil, &APIError{
				Status:  res.StatusCode(),
				Message: resp.Error,
			}
		}
	}

	return res, nil
}

func (c *Client) newRequest(r Requester) *fasthttp.Request {
	// avoid Pointer's butting
	u, _ := url.ParseRequestURI(c.Endpoint)
	u.Path = u.Path + r.Path()

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(r.Method())
	req.SetRequestURI(u.String())
	body := r.Payload()
	req.SetBody(body)

	nonce := fmt.Sprintf("%d", int64(time.Now().UTC().UnixNano()/int64(time.Millisecond)))
	payload := nonce + r.Method() + u.Path

	u.RawQuery = r.Query()
	if u.RawQuery != "" {
		payload += "?" + u.RawQuery
	}

	payload += string(body)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req
}

func decode(res *fasthttp.Response, out interface{}) error {
	var r = new(Response)
	r.Result = out
	err := json.Unmarshal(res.Body(), r.Result)
	if err != nil {
		r.Success = false
		return err
	} else {
		r.Success = true
	}

	if !r.Success {
		return fmt.Errorf("decode error")
	}

	return nil
}

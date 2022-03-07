package rest

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"net/url"
	"time"
)

// shutdownRequest doesn't get an actual response from the server and thus
// doesn't process a server reply. It only returns an error if there was
// 1) A network transmission error i.e. offline
// 2) An incorrect server URI
// 3) The server was already taken offline before
// https://support.objectivity.com/sites/default/files/docs/ig/latest/index.html#page/topics%2Frest%2FrestVersionShutdownPOST.html%23
func (c *Client) shutdownRequest(r Requester) error {
	req := c.newRequest(r)
	res := fasthttp.AcquireResponse()
	err := c.HTTPC.DoTimeout(req, res, c.HTTPTimeout)
	return checkError(err)
}

// used for JSON unmarshaling
var json = jsoniter.ConfigCompatibleWithStandardLibrary

//
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

// do build & executes the actual request from the rquester
// requester - implementation
// targetStatusCode the expected http status code i.e. 200
func (c *Client) do(r Requester) (*fasthttp.Response, error) {
	req := c.newRequest(r)
	fmt.Printf("Path: %+v\n", string(r.Path()))

	// fasthttp for http2.0
	res := fasthttp.AcquireResponse()
	err := c.HTTPC.DoTimeout(req, res, c.HTTPTimeout)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%+v\n", string(res.Body()))
	// no usefull headers

	if res.StatusCode() != r.ResponseCode() {
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

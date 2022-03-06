package rest

import "go-infinity/rest/server"

// Uptime Retrieves information about the InfiniteGraph REST server, including how long it has been running and the available versions of the API.
func (c *Client) Uptime(req *server.RequestForUptime) (*server.ResponseForUptime, error) {
	results := server.NewResponseForUptime()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// Resources Retrieves information about the available resources for this version of the API.
func (c *Client) Resources(req *server.RequestForResources) (*server.ResponseForResources, error) {
	results := server.NewResponseForResources()
	if err := c.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

// ShutdownServer Terminates the InfiniteGraph REST server running at the provided URL.
func (c *Client) ShutdownServer(req *server.RequestForShutdown) error {
	err := c.shutdownRequest(req)
	return checkError(err)

}

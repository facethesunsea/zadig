package aslan

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/koderover/zadig/pkg/tool/httpclient"
	"github.com/koderover/zadig/pkg/tool/log"
)

func (c *Client) ListTestWorkflows(testName string, header http.Header, qs url.Values) ([]byte, error) {
	url := fmt.Sprintf("/workflow/workflow/testName/%s", testName)

	res, err := c.Get(url, httpclient.SetHeadersFromHTTPHeader(header), httpclient.SetQueryParamsFromValues(qs))
	if err != nil {
		return nil, err
	}

	return res.Body(), nil
}

func (c *Client) CreateWorkflowTask(header http.Header, qs url.Values, body []byte) ([]byte, error) {
	url := "/workflow/workflowtask"

	res, err := c.Post(url, httpclient.SetHeadersFromHTTPHeader(header), httpclient.SetQueryParamsFromValues(qs), httpclient.SetBody(body))
	if err != nil {
		return nil, err
	}

	return res.Body(), nil
}

func (c *Client) CancelWorkflowTask(header http.Header, qs url.Values, id string, name string) (statusCode int, err error) {
	url := fmt.Sprintf("/workflow/workflowtask/id/%s/pipelines/%s", id, name)

	res, err := c.Delete(url, httpclient.SetHeadersFromHTTPHeader(header), httpclient.SetQueryParamsFromValues(qs))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return res.StatusCode(), nil
}

func (c *Client) RestartWorkflowTask(header http.Header, qs url.Values, id string, name string) (statusCode int, err error) {
	url := fmt.Sprintf("/workflow/workflowtask/id/%s/pipelines/%s/restart", id, name)

	res, err := c.Post(url, httpclient.SetHeadersFromHTTPHeader(header), httpclient.SetQueryParamsFromValues(qs))
	if err != nil {
		log.Errorf("RestartWorkflowTask err: %s,res: %s", err, string(res.Body()))
		return http.StatusInternalServerError, err
	}

	return res.StatusCode(), nil
}
func (c *Client) ListWorkflowTask(header http.Header, qs url.Values, commitId string) ([]byte, error) {
	url := fmt.Sprintf("/workflow/v2/status/task/info?commitId=%s", commitId)

	res, err := c.Get(url, httpclient.SetHeadersFromHTTPHeader(header), httpclient.SetQueryParamsFromValues(qs))
	if err != nil {
		return nil, err
	}

	return res.Body(), nil
}

func (c *Client) ListDelivery(header http.Header, qs url.Values, productName, workflowName, taskId, perPage, page string) ([]byte, error) {
	url := fmt.Sprintf("/delivery/releases?orgId=1&productName=%s&workflowName=%s&taskId=%s&per_page=%s&page=%s", productName, workflowName, taskId, perPage, page)

	res, err := c.Get(url, httpclient.SetHeadersFromHTTPHeader(header), httpclient.SetQueryParamsFromValues(qs))
	if err != nil {
		return nil, err
	}

	return res.Body(), nil
}

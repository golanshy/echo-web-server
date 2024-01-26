package api_echo_handler

type LambdaResponse struct {
	ReqMethod     string            `json:"req_method"`
	ReqBody       string            `json:"req_body"`
	ReqDomainName string            `json:"req_domain_name"`
	ReqHeaders    map[string]string `json:"req_headers"`
}

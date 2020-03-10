package models

type Metric []struct {
	/*	CreatedTime string      `json:"createdTime"`
		DataSource  string      `json:"dataSource"`
		Duration    interface{} `json:"duration"`
		ErrorMsg    interface{} `json:"errorMsg"`
		GroupID     string      `json:"groupId"`
		ID          string      `json:"id"`
		Location    struct {
			Host    string `json:"host"`
			Port    int64  `json:"port"`
			TLSPort int64  `json:"tlsPort"`
		} `json:"location"`
		QueueInsertionTime string `json:"queueInsertionTime"`
	*/
	//RunnerStatusCode string `json:"runnerStatusCode"`
	//	Status           string `json:"status"`
	StatusCode string `json:"statusCode"`
	//	Type             string `json:"type"`
}

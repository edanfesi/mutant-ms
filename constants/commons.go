package constants

type headers struct {
	ApplicationJSON string
	ContentType     string
	XApplicationID  string
	APIKey          string
}

type commons struct {
	ProjectName string

	Headers headers
}

var Commons = commons{
	ProjectName: "mutant-ms",

	Headers: headers{
		ContentType:     "Content-Type",
		XApplicationID:  "x-application-id",
		ApplicationJSON: "application/json",
		APIKey:          "API_KEY",
	},
}

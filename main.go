package main

import (
	"bytes"
	"html/template"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type conversion struct {
	Domain string
	Vanity string
	Prefix string
	Real   string
}

var catchAll = conversion{
	Domain: "go.kelfa.io",
	Real:   "https://github.com/kelfa/kelfa",
}

var tpl = template.Must(template.New("main").Parse(`<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
        <meta name="go-import" content="{{ .Domain }} git {{ .Real }}">
    </head>
</html>
`))

// Handler function Using AWS Lambda Proxy Request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var tplOutput bytes.Buffer
	if err := tpl.Execute(&tplOutput, catchAll); err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: tplOutput.String(), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}

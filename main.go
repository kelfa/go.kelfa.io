package main

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type conversion struct {
	Vanity string
	Real   string
	Domain string
}

var conversions = []conversion{
	conversion{Domain: "go.kelfa.io", Vanity: "aws-cloudfront-logCompactor", Real: "https://github.com/kelfa/aws-cloudfront-logCompactor"},
	conversion{Domain: "go.kelfa.io", Vanity: "elf", Real: "https://github.com/kelfa/elf"},
	conversion{Domain: "go.kelfa.io", Vanity: "go.kelfa.io", Real: "https://github.com/kelfa/go.kelfa.io"},
}

var tpl = template.Must(template.New("main").Parse(`<!DOCTYPE html>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
        <meta name="go-import" content="{{ .Domain }}/{{ .Vanity }} git {{ .Real }}">
    </head>
</html>
`))

// Handler function Using AWS Lambda Proxy Request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	for _, c := range conversions {
		if request.Path[1:] == c.Vanity || strings.HasPrefix(request.Path[1:], c.Vanity+"/") {
			var tplOutput bytes.Buffer
			if err := tpl.Execute(&tplOutput, c); err != nil {
				return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
			}
			return events.APIGatewayProxyResponse{Body: tplOutput.String(), StatusCode: 200}, nil
		}
	}
	return events.APIGatewayProxyResponse{Body: "Not found", StatusCode: 404}, nil
}

func main() {
	lambda.Start(Handler)
}

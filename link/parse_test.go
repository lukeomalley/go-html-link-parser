package link

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input             string
		expectedLinkCount int
		expectedHref      string
		expectedText      string
	}{
		{
			`
		<html>
			<body>
				<h1>Hello</h1>
				<a href="https://google.com">Go To Google</a>
			</body>
		</html>
	`, 1, "https://google.com", "Go To Google",
		},
		{
			`
			<html>
		 	  <head>
				  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
			  </head>
			  <body>
				  <h1>Social stuffs</h1>
				  <div>
					  <a href="https://www.twitter.com/lukeomalley">
						  Check me out on <strong>twitter</strong>
						  <i class="fa fa-twitter" aria-hidden="true"></i>
					  </a>
					  <a href="https://github.com/lukeomalley">
						  Check me out on <strong>Github</strong>!
					  </a>
				  </div>
			  </body>
			</html>
			`, 2, "https://www.twitter.com/lukeomalley", "Check me out on twitter",
		},
	}

	for _, tt := range tests {
		r := strings.NewReader(tt.input)
		links, err := Parse(r)

		if err != nil {
			t.Fatalf("Error parsing HTML document %T", err)
		}

		if len(links) != tt.expectedLinkCount {
			t.Errorf("Expected %d link, but got %d", tt.expectedLinkCount, len(links))
		}

		if links[0].Href != tt.expectedHref {
			t.Errorf("Expected link href to be %s, but got %s", tt.expectedHref, links[0].Href)
		}

		if links[0].Text != tt.expectedText {
			t.Errorf("Expect link text to be '%s', but got %s", tt.expectedText, links[0].Text)
		}
	}
}

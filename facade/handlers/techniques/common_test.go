package techniques

import (
	"testing"
)

var technique HeadTagsTechnique

func TestHeadTagsTechnique_Extract(t *testing.T) {
	html := `<head>
			 <meta http-equiv="content-type" content="text/html; charset=UTF-8" />
			 <meta name="author" content="Will Larson" />
			 <meta name="description" content="Will Larson&#39;s blog about programming and other things." />
			 <meta name="keywords" content="Blog Will Larson Programming Life" />
			 <link rel="alternate" type="application/rss+xml" title="Page Feed" href="/feeds/" />
			 <link rel="canonical" href="http://lethain.com/digg-v4-architecture-process/">
			 <title>Digg v4&#39;s Architecture and Development Processes - Irrational Exuberance</title>
			 </head>`
	extracted := technique.Extract(html)
	allEmpty := true
	for _, value := range extracted {
		if len(value) > 0 {
			allEmpty = false
		}
	}
	if allEmpty {
		t.Fail()
	}
}

package techniques

import (
	"log"
	"testing"
)

func TestHeadTagsTechnique_Extract(t *testing.T) {
	var technique HeadTagsTechnique

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
	log.Println(extracted)
}

func TestHTML5SemanticTagsTechnique_Extract(t *testing.T) {
	var technique HTML5SemanticTagsTechnique

	html := `<html>
	       <body>
	         <h1>This is not a title to HTML5SemanticTags</h1>
	         <article>
	           <h1>This is a title</h1>
	           <p>This is a description.</p>
	           <p>This is not a description.</p>
	         </article>
	         <video>
	           <source src="this_is_a_video.mp4">
	         </video>
	       </body>
	     </html>`
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
	log.Println(extracted)
}

func TestSemanticTagsTechnique_Extract(t *testing.T) {
	var technique SemanticTagsTechnique
	html := `<head>
	 		 <meta http-equiv="content-type" content="text/html; charset=UTF-8" />
	 		 <meta name="author" content="Will Larson" />
	 		 <meta name="description" content="Will Larson&#39;s blog about programming and other things." />
	 		 <meta name="keywords" content="Blog Will Larson Programming Life" />
	 		 <link rel="alternate" type="application/rss+xml" title="Page Feed" href="/feeds/" />
	 		 <link rel="canonical" href="http://lethain.com/digg-v4-architecture-process/">
	 		 <title>Digg v4&#39;s Architecture and Development Processes - Irrational Exuberance</title>
	 		 <h1 >H11</h1>
	 		 <h2 >H22</h2>
	 		 <h3 >H33</h3>
	 		 <p> PP </p>
			 <image src="http://p3.pstatp.com/large/pgc-image/b2718ff73795485f82a6126d7f4c814f"></image>
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
	log.Println(extracted)
}

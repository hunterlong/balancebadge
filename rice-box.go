package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "svg.xml",
		FileModTime: time.Unix(1529425553, 0),
		Content:     string("<svg xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" width=\"{{ .Width }}\" height=\"20\">\n    <linearGradient id=\"b\" x2=\"0\" y2=\"100%\"><stop offset=\"0\" stop-color=\"#bbb\" stop-opacity=\".1\"/>\n        <stop offset=\"1\" stop-opacity=\".1\"/>\n    </linearGradient>\n    <clipPath id=\"a\">\n        <rect width=\"{{ .Width }}\" height=\"20\" rx=\"3\" fill=\"#fff\"/>\n    </clipPath>\n    <g clip-path=\"url(#a)\"><path fill=\"#{{ .LeftColor }}\" d=\"M0 0h{{ .LeftSize }}v20H0z\"/>\n        <path fill=\"#{{.RightColor}}\" d=\"M{{ .LeftSize }} 0h{{ .Width }}v20H{{ .LeftSize }}z\"/>\n        <path fill=\"url(#b)\" d=\"M0 0h{{ .LeftSize }}{{ .Width }}v20H0z\"/>\n    </g>\n    <g fill=\"#fff\" text-anchor=\"middle\" font-family=\"DejaVu Sans,Verdana,Geneva,sans-serif\" font-size=\"110\">\n        <text x=\"305\" y=\"150\" fill=\"#010101\" fill-opacity=\".3\" transform=\"scale(.1)\" textLength=\"490\">{{ .Label }}</text>\n        <text x=\"305\" y=\"140\" transform=\"scale(.1)\" textLength=\"490\">{{ .Label }}</text>\n        <text x=\"{{ .RightTextX }}\" y=\"150\" fill=\"#010101\" fill-opacity=\".3\" transform=\"scale(.1)\" textLength=\"{{ .RightTextSize }}\">{{ .Balance }} {{ .Coin }}</text>\n        <text x=\"{{ .RightTextX }}\" y=\"140\" transform=\"scale(.1)\" textLength=\"{{ .RightTextSize }}\">{{ .Balance }} {{ .Coin }}</text>\n    </g>\n</svg>"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1529425553, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "svg.xml"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`svg`, &embedded.EmbeddedBox{
		Name: `svg`,
		Time: time.Unix(1529425553, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"svg.xml": file2,
		},
	})
}

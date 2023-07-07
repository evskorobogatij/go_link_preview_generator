package main

import (
	"net/http"

	"github.com/soranoba/googp"
)

type PreviewInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PreviewUrl  string `json:"preview_url"`
}

func GeneratePreview(url string, preview *PreviewInfo) error {
	var ogp1 googp.OGP
	response, err := http.Get(url)
	if err != nil {
		return err
	} else {
		defer response.Body.Close()

		if err := googp.Parse(response, &ogp1); err != nil {
			return err
		}

		preview.Title = ogp1.Title
		preview.Description = ogp1.Description
		if len(ogp1.Images) > 0 {
			tmp_path := ogp1.Images[0].URL
			if tmp_path[0] == '/' {
				preview.PreviewUrl = response.Request.URL.Scheme + "://" + response.Request.URL.Host + ogp1.Images[0].URL
			} else {
				preview.PreviewUrl = ogp1.Images[0].URL
			}

		} else {
			preview.PreviewUrl = ""
		}

	}

	return nil
}

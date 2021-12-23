package main

import "github.com/n0madic/google-play-scraper/pkg/app"

// searchPlayStore() - search app in Play Store using the package name
func (androidapp *AndroidApp) searchPlayStore() error {
	playstoreinfo := app.New(androidapp.apk.PackageName(), app.Options{
		Country:  "us",
		Language: "us",
	})

	if err = playstoreinfo.LoadDetails(); err != nil {
		return err
	}

	androidapp.playStoreInfo = *playstoreinfo
	androidapp.playStoreFound = true
	return nil
}

package main

// downloadFile.go
//
// part of Add-License
//
// Copyright (C) 2022 Th3-S1lenc3
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

import (
	"fmt"
	"time"

	"github.com/cavaliergopher/grab/v3"
)

func DownloadFile(dir string, remoteFileURL string) error {
	// Create Client
	client := grab.NewClient()
	req, _ := grab.NewRequest(dir, remoteFileURL)

	// Start Download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// Start UI Loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf(
				"  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress(),
			)
		case <-resp.Done:
			break Loop
		}
	}

	if err := resp.Err(); err != nil {
		return fmt.Errorf("Download failed %v\n", err)
	}

	fmt.Printf("Download saved to %v \n", resp.Filename)

	return nil
}

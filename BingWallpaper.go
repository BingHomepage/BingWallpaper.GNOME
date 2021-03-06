/*
	MIT License

	Copyright (c) 2018 BingHomepage

	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:

	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.

	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
*/
package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"time"

	"github.com/BingHomepage/BingHomepage.Go"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	if runtime.GOOS == "windows" {
		log.Fatal("Nope... try this one -> https://github.com/BingHomepage/BingWallpaper.Windows")
	} else if runtime.GOOS == "darwin" {
		log.Fatal("Siri-ously?")
	} else {
		var interval string

		flag.StringVar(&interval, "interval", "12h",
			"The duration in which wallpaper is to be updated. Valid units are \"ns\", \"us\", \"ms\", \"s\", \"m\", \"h\".")
		flag.Parse()
		repeat(func() {
			data, err := BingHomepage.Get("US")
			if err != nil {
				log.Println(err)
			}
			fileName := usr.HomeDir + "/BingWallpaper.jpg"
			file, err := os.Create(fileName)
			if err != nil {
				log.Println(err)
			}
			defer file.Close()
			resp, err := http.Get(data.URL)
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()
			io.Copy(file, resp.Body)
			exec.Command("/usr/bin/gsettings", "set", "org.gnome.desktop.background", "picture-uri", "file://"+fileName).Run()
			log.Printf("Response status -> %s.\n", resp.Status)
		}, interval)
	}
}
func repeat(f func(), interval string) {
	f()
	d, err := time.ParseDuration(interval)
	if err != nil {
		log.Println(err)
	}
	for range time.Tick(d) {
		f()
	}
}

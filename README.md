# BingWallpaper

Fetches and applies the image of the day from Bing as the wallpaper for the GNOME environment.

## Building

### Getting BingHomepage.Go

BingWallpaper for GNOME environment uses [BingHomepage.Go](https://github.com/BingHomepage/BingHomepage.Go) library to fetch details.

To install the library, run:

```bash
go get github.com/BingHomepage/BingHomepage.Go
```

### Clone this repository

```bash
mkdir BingWallpaper
cd BingWallpaper
git clone https://github.com/BingHomepage/BingWallpaper.GNOME.git .
```

### Building the app

```bash
go build
```

## Moving executable to `bin`

```bash
mv ./BingWallpaper /usr/bin
```

## Running

To run the application, simply execute `BingWallpaper` with optional flag `-interval` to set the interval in which the wallpaper is to be updated, for example, `-interval 5m` would update wallpaper in every 5 minutes.

Default interval is 12h, meaning the wallpaper will be updated in every 12 hour.

```bash
BingWallpaper -interval 12h30m &; disown;
```

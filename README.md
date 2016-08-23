# madek

**A Go library and command line tool that simplifies accessing the [Madek API](https://medienarchiv.zhdk.ch/api/browser/index.html).**

## Installation

```
$ go get github.com/IAD-ZHDK/madek/cmd/madek
```
 
## Usage

```
madek.

Usage:
  madek fetch <id> [options]
  madek server [options]

Options:
  -h --help                   Show this screen.
  -a --address=<url>          The address of the madek instance [default: https://medienarchiv.zhdk.ch].
  -u --username=<username>    The username used for authentication.
  -p --password=<password>    The password used for authentication.
  -c --cache                  Cache requests in server mode.
```

## Data Format

```js
{
  "id": "82108639-c4a6-412d-b347-341fe5284caa",
  "created_at": "2016-05-25T09:46:40.533Z",
  "meta_data": {
    "subtitle": "MEDUZA is part of the installation that emerged under the theme Strange Garden in the module Spatial Interaction.",
    "title": "Meduza"
  },
  "media_entries": [
    {
      "id": "a9960da5-b72f-43a5-a52d-0ab85637dc0f",
      "meta_data": {
        "title": "Bild Meduza"
      },
      "created_at": "2016-05-25T09:51:56.861Z",
      "file_id": "e401ae20-d5f5-4663-b11b-b4ada3f14252",
      "file_name": "04_strings_collected.jpg",
      "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/e401ae20-d5f5-4663-b11b-b4ada3f14252/data-stream",
      "download_url": "https://medienarchiv.zhdk.ch/files/e401ae20-d5f5-4663-b11b-b4ada3f14252",
      "previews": [
        {
          "id": "ea94e719-d38f-4b58-b0f1-8f2662939859",
          "type": "image",
          "content_type": "image/jpeg",
          "size": "x_large",
          "width": 1024,
          "height": 683,
          "url": "https://medienarchiv.zhdk.ch/media/ea94e719-d38f-4b58-b0f1-8f2662939859"
        },
        // ...
      ]
    },
    // ...
  ]
}
```

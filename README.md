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

```json
{
  "id": "82108639-c4a6-412d-b347-341fe5284caa",
  "created_at": "2016-05-25T09:46:40.533Z",
  "meta_data": {
    "title": "Meduza",
    "subtitle": "MEDUZA is part of the installation that emerged under the theme Strange Garden in the module Spatial Interaction.",
    "description": "The installation invites the viewer to discover a whole new relation between light, space and movement. [...]",
    "authors": [
      {
        "first_name": "Thomas",
        "last_name": "Guthruf"
      },
      {
        "first_name": "Simon",
        "last_name": "Häusler"
      }
    ],
    "keywords": [
      "Spatial Interaction"
    ],
    "genres": [
      "Design"
    ],
    "year": "2016",
    "copyright": {
      "holder": "",
      "usage": "",
      "licenses": null
    },
    "affiliation": [
      {
        "name": "Bachelor Design - Interaction Design",
        "pseudonym": "DDE_FDE_BDE_VIAD.alle"
      },
      {
        "name": "Vertiefung Interaction Design",
        "pseudonym": "DDE_FDE_VIAD.alle"
      }
    ]
  },
  "media_entries": [
    {
      "id": "1e3bada4-5571-4d49-a712-f044cdc745bb",
      "meta_data": {
        "title": "Bild Meduza",
        "genres": [
          "Design"
        ],
        "copyright": {
          "holder": "Interaction Design",
          "usage": "Das Werk darf nur mit Einwilligung des Autors/Rechteinhabers weiter verwendet werden.",
          "licenses": [
            "Alle Rechte vorbehalten"
          ]
        }
      },
      "created_at": "2016-05-25T09:55:01.052Z",
      "file_id": "644a5f64-ebd6-4d75-a10e-391c87312586",
      "file_name": "10_meduza_trough_the_net.jpg",
      "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/644a5f64-ebd6-4d75-a10e-391c87312586/data-stream",
      "download_url": "https://medienarchiv.zhdk.ch/files/644a5f64-ebd6-4d75-a10e-391c87312586",
      "previews": [
        {
          "id": "84865b8b-1c72-4e5f-9c20-0c0e219328e6",
          "type": "image",
          "content_type": "image/jpeg",
          "size": "maximum",
          "width": 4000,
          "height": 6000,
          "url": "https://medienarchiv.zhdk.ch/media/84865b8b-1c72-4e5f-9c20-0c0e219328e6"
        }
        // ...
      ]
    }
    // ...
  ]
}
```

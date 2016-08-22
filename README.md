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
  "title": "Meduza",
  "created_at": "2016-05-25T09:46:40.533Z",
  "media": [
    {
      "id": "31b7f0fe-0eb8-4b52-96e2-8fa28b2807d7",
      "title": "Servo Control",
      "created_at": "2016-05-25T10:48:40.157Z",
      "file_id": "da36ebd4-66aa-444f-9441-c55308f9cb3d",
      "file_name": "ServoControl.zip",
      "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/da36ebd4-66aa-444f-9441-c55308f9cb3d/data-stream",
      "file_url": "https://medienarchiv.zhdk.ch/files/da36ebd4-66aa-444f-9441-c55308f9cb3d",
      "previews": null
    },
    {
      "id": "c39b6dfb-f51f-4bca-9fd0-534462565076",
      "title": "Dokumentation Meduza",
      "created_at": "2016-05-25T09:44:29.808Z",
      "file_id": "5140c3b8-cc6e-42bf-be1d-90abd4fc7854",
      "file_name": "GH_dokumentation.pdf",
      "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/5140c3b8-cc6e-42bf-be1d-90abd4fc7854/data-stream",
      "file_url": "https://medienarchiv.zhdk.ch/files/5140c3b8-cc6e-42bf-be1d-90abd4fc7854",
      "previews": [
        {
          "id": "4e80f623-a3c9-4e1c-b180-c852683f4e71",
          "type": "image",
          "content_type": "image/jpeg",
          "size": "x_large",
          "width": 1024,
          "height": 724,
          "url": "https://medienarchiv.zhdk.ch/media/4e80f623-a3c9-4e1c-b180-c852683f4e71"
        },
        // ...
      ]
    },
    // ...
  ]
}
```

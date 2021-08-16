package madek

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var address = "https://medienarchiv.zhdk.ch"
var username = os.Getenv("USER")
var password = os.Getenv("PASS")

func TestClient(t *testing.T) {
	client := NewClient(address, username, password)

	coll, err := client.CompileCollection("82108639-c4a6-412d-b347-341fe5284caa")
	assert.NoError(t, err)

	coll.MediaEntries = []*MediaEntry{
		coll.MediaEntries[1],
		coll.MediaEntries[2],
		coll.MediaEntries[3],
	}

	bytes, err := json.MarshalIndent(coll, "", "  ")
	assert.NoError(t, err)

	assert.JSONEq(t, `{
	  "id": "82108639-c4a6-412d-b347-341fe5284caa",
	  "created_at": "2016-05-25T09:46:40.533Z",
	  "meta_data": {
		"title": "Meduza",
		"subtitle": "MEDUZA is part of the installation that emerged under the theme Strange Garden in the module Spatial Interaction.",
		"description": "The installation invites the viewer to discover a whole new relation between light, space and movement. It gives a visual, tactile and interactive experience to the visitors. At the beginning, the installation was planned to be a walkable cave, where visitors could perceive a meditative composition of various sounds. During discussions with other groups we decided to create a visual and tactile experience to bring more diversity to the space. Furthermore, we realized that our idea of the dimension was to bold. There was not enough space next to the other installations and not enough time to produce the necessary amount of material. Therefore, we came up with MEDUZA, a curious creature with resemblance to a jellyfish. This creature communicates visual as well as tactile with the visitors. Through movement of its tentacles it changes from transparency to translucency as soon as someone lays beneath it. With the condensing of the textile bands a new sense of space isolates and protects the viewer. Moreover, the easiness of the fabric and the subtle movement suggests a new feeling of safety and warmth.",
		"authors": [
		  {
			"id": "959a242d-5c2a-4d63-a532-a05d7c48d284",
			"first_name": "Thomas",
			"last_name": "Guthruf"
		  },
		  {
			"id": "dea671d9-40b1-4adf-a87e-4399b83e3cc8",
			"first_name": "Simon",
			"last_name": "Häusler"
		  }
		],
		"genres": [
		  "Design"
		],
		"year": "2016",
		"copyright": {},
		"affiliation": [
		  {
			"id": "3c9b2eeb-f6e6-4a02-9507-c4ee9ce5fe2e",
			"name": "Bachelor Design - Interaction Design",
			"pseudonym": "DDE_FDE_BDE_VIAD.alle"
		  },
		  {
			"id": "afe0553e-949a-40e2-ae3c-d7dc2b2ca46f",
			"name": "Fachrichtung Interaction Design (BA & MA)",
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
	      "file_type": "image/jpeg",
	      "file_size": 18830798,
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/644a5f64-ebd6-4d75-a10e-391c87312586/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/644a5f64-ebd6-4d75-a10e-391c87312586",
		  "previews": [
			{
			  "id": "0bf94231-ec89-4859-995e-534bfb6206ad",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 333,
			  "height": 500,
			  "url": "https://medienarchiv.zhdk.ch/media/0bf94231-ec89-4859-995e-534bfb6206ad"
			},
			{
			  "id": "3f072e0f-d13e-418f-b797-824e6f0cf0fc",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 67,
			  "height": 100,
			  "url": "https://medienarchiv.zhdk.ch/media/3f072e0f-d13e-418f-b797-824e6f0cf0fc"
			},
			{
			  "id": "57afb581-746d-41a5-8e4e-d0fdbbeb72a0",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 200,
			  "height": 300,
			  "url": "https://medienarchiv.zhdk.ch/media/57afb581-746d-41a5-8e4e-d0fdbbeb72a0"
			},
			{
			  "id": "84865b8b-1c72-4e5f-9c20-0c0e219328e6",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "maximum",
			  "width": 4000,
			  "height": 6000,
			  "url": "https://medienarchiv.zhdk.ch/media/84865b8b-1c72-4e5f-9c20-0c0e219328e6"
			},
			{
			  "id": "c3f84bea-acb5-4cac-b273-8c13ad10edc7",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 512,
			  "height": 768,
			  "url": "https://medienarchiv.zhdk.ch/media/c3f84bea-acb5-4cac-b273-8c13ad10edc7"
			},
			{
			  "id": "f245209b-2b82-452c-b7d4-3a968db7fce6",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 83,
			  "height": 125,
			  "url": "https://medienarchiv.zhdk.ch/media/f245209b-2b82-452c-b7d4-3a968db7fce6"
			}
		  ]
		},
		{
		  "id": "31b7f0fe-0eb8-4b52-96e2-8fa28b2807d7",
		  "meta_data": {
			"title": "Servo Control",
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
		  "created_at": "2016-05-25T10:48:40.157Z",
		  "file_id": "da36ebd4-66aa-444f-9441-c55308f9cb3d",
		  "file_name": "ServoControl.zip",
	      "file_type": "application/zip",
	      "file_size": 1151,
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/da36ebd4-66aa-444f-9441-c55308f9cb3d/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/da36ebd4-66aa-444f-9441-c55308f9cb3d",
		  "previews": null
		},
		{
		  "id": "44ab4b5b-4b3f-4854-9d5f-66a852a25eab",
		  "meta_data": {
			"title": "Video",
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
		  "created_at": "2016-05-25T10:49:49.487Z",
		  "file_id": "6ca65daf-7ec5-46ab-98a6-d3fff08f18a0",
		  "file_name": "TH_video.mp4",
	      "file_type": "video/mp4",
	      "file_size": 100419381,
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/6ca65daf-7ec5-46ab-98a6-d3fff08f18a0/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/6ca65daf-7ec5-46ab-98a6-d3fff08f18a0",
		  "previews": [
			{
			  "id": "32391c88-b495-4c41-89c8-be5ed5a8b215",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "maximum",
			  "width": 1920,
			  "height": 1078,
			  "url": "https://medienarchiv.zhdk.ch/media/32391c88-b495-4c41-89c8-be5ed5a8b215"
			},
			{
			  "id": "34ce3a55-2c2b-4a4a-aea4-d97fe9be4234",
			  "type": "video",
			  "content_type": "video/webm",
			  "size": "large",
			  "width": 620,
			  "height": 348,
			  "url": "https://medienarchiv.zhdk.ch/media/34ce3a55-2c2b-4a4a-aea4-d97fe9be4234"
			},
			{
			  "id": "62911c6b-8663-43dc-a13a-dcc856e9cd3d",
			  "type": "video",
			  "content_type": "video/mp4",
			  "size": "large",
			  "width": 1920,
			  "height": 1080,
			  "url": "https://medienarchiv.zhdk.ch/media/62911c6b-8663-43dc-a13a-dcc856e9cd3d"
			},
			{
			  "id": "80aa7fa8-fbc9-482b-8802-c61e427de36e",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 100,
			  "height": 56,
			  "url": "https://medienarchiv.zhdk.ch/media/80aa7fa8-fbc9-482b-8802-c61e427de36e"
			},
			{
			  "id": "80e8b5b2-4944-46a5-a5fb-9aa21e318ff9",
			  "type": "video",
			  "content_type": "video/webm",
			  "size": "large",
			  "width": 1920,
			  "height": 1080,
			  "url": "https://medienarchiv.zhdk.ch/media/80e8b5b2-4944-46a5-a5fb-9aa21e318ff9"
			},
			{
			  "id": "a8d0dc69-4acb-4ddc-996b-c6e3ec7b12bd",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 620,
			  "height": 348,
			  "url": "https://medienarchiv.zhdk.ch/media/a8d0dc69-4acb-4ddc-996b-c6e3ec7b12bd"
			},
			{
			  "id": "c83719e2-0cda-4a6b-8476-6298b340b1e1",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 620,
			  "height": 348,
			  "url": "https://medienarchiv.zhdk.ch/media/c83719e2-0cda-4a6b-8476-6298b340b1e1"
			},
			{
			  "id": "cb2705a3-b7fe-4097-98fa-126108876381",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 300,
			  "height": 168,
			  "url": "https://medienarchiv.zhdk.ch/media/cb2705a3-b7fe-4097-98fa-126108876381"
			},
			{
			  "id": "d16696a5-9971-4ab4-9fdc-36549181f390",
			  "type": "video",
			  "content_type": "video/mp4",
			  "size": "large",
			  "width": 620,
			  "height": 348,
			  "url": "https://medienarchiv.zhdk.ch/media/d16696a5-9971-4ab4-9fdc-36549181f390"
			},
			{
			  "id": "eacadcc3-a0b1-4e6f-8165-06794e730314",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 1024,
			  "height": 575,
			  "url": "https://medienarchiv.zhdk.ch/media/eacadcc3-a0b1-4e6f-8165-06794e730314"
			},
			{
			  "id": "ff312114-8733-428c-b117-f6a7280cf31e",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 125,
			  "height": 70,
			  "url": "https://medienarchiv.zhdk.ch/media/ff312114-8733-428c-b117-f6a7280cf31e"
			}
		  ]
		}
	  ]
	}`, string(bytes))
}

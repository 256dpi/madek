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
			"name": "Vertiefung Interaction Design",
			"pseudonym": "DDE_FDE_VIAD.alle"
		  }
		]
	  },
	  "media_entries": [
		{
		  "id": "1a7d6996-771b-4448-a29b-4f52caa8d84c",
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
		  "created_at": "2016-05-25T09:51:36.375Z",
		  "file_id": "bf62e010-a9c7-48cf-9b71-7edfdb50699a",
		  "file_name": "03_proof_of_concept.jpg",
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/bf62e010-a9c7-48cf-9b71-7edfdb50699a/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/bf62e010-a9c7-48cf-9b71-7edfdb50699a",
		  "previews": [
			{
			  "id": "04e67f0c-c5b5-4dcb-bd93-d058d85f34a7",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 512,
			  "height": 768,
			  "url": "https://medienarchiv.zhdk.ch/media/04e67f0c-c5b5-4dcb-bd93-d058d85f34a7"
			},
			{
			  "id": "195686c1-bd15-48fd-a369-d135d14db279",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 333,
			  "height": 500,
			  "url": "https://medienarchiv.zhdk.ch/media/195686c1-bd15-48fd-a369-d135d14db279"
			},
			{
			  "id": "6fa2ef5d-5e18-499c-8d7c-2875acc197d4",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 83,
			  "height": 125,
			  "url": "https://medienarchiv.zhdk.ch/media/6fa2ef5d-5e18-499c-8d7c-2875acc197d4"
			},
			{
			  "id": "961a9a8f-4d8e-4e3e-8fb6-39f3ab22eb4d",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 67,
			  "height": 100,
			  "url": "https://medienarchiv.zhdk.ch/media/961a9a8f-4d8e-4e3e-8fb6-39f3ab22eb4d"
			},
			{
			  "id": "c2c66ed6-c109-4292-bbef-08bf22619d6b",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 200,
			  "height": 300,
			  "url": "https://medienarchiv.zhdk.ch/media/c2c66ed6-c109-4292-bbef-08bf22619d6b"
			},
			{
			  "id": "ec20cbcc-ecad-49a1-8eb9-f115cd2ff7be",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "maximum",
			  "width": 4000,
			  "height": 6000,
			  "url": "https://medienarchiv.zhdk.ch/media/ec20cbcc-ecad-49a1-8eb9-f115cd2ff7be"
			}
		  ]
		},
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
		},
		{
		  "id": "7c7a5f4e-37e7-49fa-8171-68aca503bd31",
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
		  "created_at": "2016-05-25T09:51:13.992Z",
		  "file_id": "f2d90ca2-db5f-4f8f-966c-39c527af2c9d",
		  "file_name": "02_attach_the_fabrics.jpg",
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/f2d90ca2-db5f-4f8f-966c-39c527af2c9d/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/f2d90ca2-db5f-4f8f-966c-39c527af2c9d",
		  "previews": [
			{
			  "id": "0cc0b4c9-1851-42b9-9987-14b9c2d7a16d",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "maximum",
			  "width": 4000,
			  "height": 6000,
			  "url": "https://medienarchiv.zhdk.ch/media/0cc0b4c9-1851-42b9-9987-14b9c2d7a16d"
			},
			{
			  "id": "2e0d3ca0-64c3-4ec0-8090-3fbea530c648",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 200,
			  "height": 300,
			  "url": "https://medienarchiv.zhdk.ch/media/2e0d3ca0-64c3-4ec0-8090-3fbea530c648"
			},
			{
			  "id": "35a8d3b0-40f4-4a07-8b35-869f557131b9",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 83,
			  "height": 125,
			  "url": "https://medienarchiv.zhdk.ch/media/35a8d3b0-40f4-4a07-8b35-869f557131b9"
			},
			{
			  "id": "3f128f65-5aab-45cb-8d09-acfecf952848",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 512,
			  "height": 768,
			  "url": "https://medienarchiv.zhdk.ch/media/3f128f65-5aab-45cb-8d09-acfecf952848"
			},
			{
			  "id": "55dd6d24-6a46-48b7-839e-1d03afbae1a4",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 333,
			  "height": 500,
			  "url": "https://medienarchiv.zhdk.ch/media/55dd6d24-6a46-48b7-839e-1d03afbae1a4"
			},
			{
			  "id": "ffa18f91-15a8-42d2-b346-7c4b7e0fdfc3",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 67,
			  "height": 100,
			  "url": "https://medienarchiv.zhdk.ch/media/ffa18f91-15a8-42d2-b346-7c4b7e0fdfc3"
			}
		  ]
		},
		{
		  "id": "89a14418-1f14-4ced-9c59-33b896630648",
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
		  "created_at": "2016-05-25T09:53:53.827Z",
		  "file_id": "7f649d9e-201d-4505-8e7a-3c6f1654e3a9",
		  "file_name": "08_meduza_from_beneath_2.jpg",
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/7f649d9e-201d-4505-8e7a-3c6f1654e3a9/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/7f649d9e-201d-4505-8e7a-3c6f1654e3a9",
		  "previews": [
			{
			  "id": "1114c4d8-daac-46d5-99f9-e19f0b72d2b3",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "maximum",
			  "width": 4000,
			  "height": 6000,
			  "url": "https://medienarchiv.zhdk.ch/media/1114c4d8-daac-46d5-99f9-e19f0b72d2b3"
			},
			{
			  "id": "1c56b85c-701f-468d-b941-a0799182c712",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 83,
			  "height": 125,
			  "url": "https://medienarchiv.zhdk.ch/media/1c56b85c-701f-468d-b941-a0799182c712"
			},
			{
			  "id": "2d6cbd44-7f86-4059-8c0f-1ca1fd1549c1",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 200,
			  "height": 300,
			  "url": "https://medienarchiv.zhdk.ch/media/2d6cbd44-7f86-4059-8c0f-1ca1fd1549c1"
			},
			{
			  "id": "329e1ea9-bad7-456d-850c-ef2c1f0285e4",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 333,
			  "height": 500,
			  "url": "https://medienarchiv.zhdk.ch/media/329e1ea9-bad7-456d-850c-ef2c1f0285e4"
			},
			{
			  "id": "465ac576-4c1a-4af8-b3ec-df6f8a544dea",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 67,
			  "height": 100,
			  "url": "https://medienarchiv.zhdk.ch/media/465ac576-4c1a-4af8-b3ec-df6f8a544dea"
			},
			{
			  "id": "d40c0bd5-3c45-4f05-9f68-a6fb90fd83ae",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 512,
			  "height": 768,
			  "url": "https://medienarchiv.zhdk.ch/media/d40c0bd5-3c45-4f05-9f68-a6fb90fd83ae"
			}
		  ]
		},
		{
		  "id": "97f9e42f-fb15-4719-b1b9-21c1c60c936c",
		  "meta_data": {
			"title": "Led Control",
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
		  "created_at": "2016-05-25T10:48:39.998Z",
		  "file_id": "4a8ae961-2561-4b20-9de3-46d42ac70820",
		  "file_name": "LedControl.zip",
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/4a8ae961-2561-4b20-9de3-46d42ac70820/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/4a8ae961-2561-4b20-9de3-46d42ac70820",
		  "previews": null
		},
		{
		  "id": "a06d73b7-e94b-478e-8abd-c2bec5959d35",
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
		  "created_at": "2016-05-25T09:50:53.079Z",
		  "file_id": "59df7020-ce15-411e-8a7e-fe9665a86cfa",
		  "file_name": "01_attach_the_fabrics.jpg",
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/59df7020-ce15-411e-8a7e-fe9665a86cfa/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/59df7020-ce15-411e-8a7e-fe9665a86cfa",
		  "previews": [
			{
			  "id": "31a0c964-c7d6-4634-8855-60b827dcb86d",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 300,
			  "height": 200,
			  "url": "https://medienarchiv.zhdk.ch/media/31a0c964-c7d6-4634-8855-60b827dcb86d"
			},
			{
			  "id": "7734cfd6-f248-47ad-a2fa-a1f79f80ab35",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 620,
			  "height": 413,
			  "url": "https://medienarchiv.zhdk.ch/media/7734cfd6-f248-47ad-a2fa-a1f79f80ab35"
			},
			{
			  "id": "a45eb9e4-fc41-4a9e-a762-5595d0dd11dc",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 125,
			  "height": 83,
			  "url": "https://medienarchiv.zhdk.ch/media/a45eb9e4-fc41-4a9e-a762-5595d0dd11dc"
			},
			{
			  "id": "bcfbfd08-2e38-4060-8f5b-da1f27fea454",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 100,
			  "height": 67,
			  "url": "https://medienarchiv.zhdk.ch/media/bcfbfd08-2e38-4060-8f5b-da1f27fea454"
			},
			{
			  "id": "cc1f3fbe-0b1e-4f25-ac80-c85032cc02cd",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "maximum",
			  "width": 6000,
			  "height": 4000,
			  "url": "https://medienarchiv.zhdk.ch/media/cc1f3fbe-0b1e-4f25-ac80-c85032cc02cd"
			},
			{
			  "id": "d3f5d441-22ba-4160-8197-c7ee9ae5620e",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 1024,
			  "height": 683,
			  "url": "https://medienarchiv.zhdk.ch/media/d3f5d441-22ba-4160-8197-c7ee9ae5620e"
			}
		  ]
		},
		{
		  "id": "a9960da5-b72f-43a5-a52d-0ab85637dc0f",
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
		  "created_at": "2016-05-25T09:51:56.861Z",
		  "file_id": "e401ae20-d5f5-4663-b11b-b4ada3f14252",
		  "file_name": "04_strings_collected.jpg",
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/e401ae20-d5f5-4663-b11b-b4ada3f14252/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/e401ae20-d5f5-4663-b11b-b4ada3f14252",
		  "previews": [
			{
			  "id": "154e6460-defb-4077-b1a1-647c5ad42d39",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 125,
			  "height": 83,
			  "url": "https://medienarchiv.zhdk.ch/media/154e6460-defb-4077-b1a1-647c5ad42d39"
			},
			{
			  "id": "8110edc8-9c2f-45d0-81bc-3b0d79aae184",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "maximum",
			  "width": 6000,
			  "height": 4000,
			  "url": "https://medienarchiv.zhdk.ch/media/8110edc8-9c2f-45d0-81bc-3b0d79aae184"
			},
			{
			  "id": "928f5870-a6bd-41d7-807a-babd379a5eff",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 620,
			  "height": 413,
			  "url": "https://medienarchiv.zhdk.ch/media/928f5870-a6bd-41d7-807a-babd379a5eff"
			},
			{
			  "id": "c2aa529e-c81f-40b6-a9ae-c0f4ef7267bc",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 100,
			  "height": 67,
			  "url": "https://medienarchiv.zhdk.ch/media/c2aa529e-c81f-40b6-a9ae-c0f4ef7267bc"
			},
			{
			  "id": "caa86984-2409-4672-b827-1a3b3c298a5b",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 300,
			  "height": 200,
			  "url": "https://medienarchiv.zhdk.ch/media/caa86984-2409-4672-b827-1a3b3c298a5b"
			},
			{
			  "id": "ea94e719-d38f-4b58-b0f1-8f2662939859",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 1024,
			  "height": 683,
			  "url": "https://medienarchiv.zhdk.ch/media/ea94e719-d38f-4b58-b0f1-8f2662939859"
			}
		  ]
		},
		{
		  "id": "af61defb-575d-4e36-ba32-6f777956effe",
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
		  "created_at": "2016-05-25T09:54:29.24Z",
		  "file_id": "63e1c0e9-8703-4128-9415-7fa032846ebd",
		  "file_name": "09_meduza_sideview.jpg",
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/63e1c0e9-8703-4128-9415-7fa032846ebd/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/63e1c0e9-8703-4128-9415-7fa032846ebd",
		  "previews": [
			{
			  "id": "08fc5442-c8d1-4ca7-853e-b5599a7c4b51",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "maximum",
			  "width": 4000,
			  "height": 6000,
			  "url": "https://medienarchiv.zhdk.ch/media/08fc5442-c8d1-4ca7-853e-b5599a7c4b51"
			},
			{
			  "id": "0ea2724a-280f-4c98-8127-6041329a46d4",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 83,
			  "height": 125,
			  "url": "https://medienarchiv.zhdk.ch/media/0ea2724a-280f-4c98-8127-6041329a46d4"
			},
			{
			  "id": "524e77d4-5435-4f9d-8d57-0df6eb2fa020",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 333,
			  "height": 500,
			  "url": "https://medienarchiv.zhdk.ch/media/524e77d4-5435-4f9d-8d57-0df6eb2fa020"
			},
			{
			  "id": "d3d6fc85-5c26-4e44-95cc-5385ef99a084",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 67,
			  "height": 100,
			  "url": "https://medienarchiv.zhdk.ch/media/d3d6fc85-5c26-4e44-95cc-5385ef99a084"
			},
			{
			  "id": "de416ff6-f857-4e8d-b655-082fc6361a3b",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 200,
			  "height": 300,
			  "url": "https://medienarchiv.zhdk.ch/media/de416ff6-f857-4e8d-b655-082fc6361a3b"
			},
			{
			  "id": "fc2e884a-809e-4b43-a861-d951fc478b9b",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 512,
			  "height": 768,
			  "url": "https://medienarchiv.zhdk.ch/media/fc2e884a-809e-4b43-a861-d951fc478b9b"
			}
		  ]
		},
		{
		  "id": "c39b6dfb-f51f-4bca-9fd0-534462565076",
		  "meta_data": {
			"title": "Dokumentation Meduza",
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
		  "created_at": "2016-05-25T09:44:29.808Z",
		  "file_id": "5140c3b8-cc6e-42bf-be1d-90abd4fc7854",
		  "file_name": "GH_dokumentation.pdf",
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/5140c3b8-cc6e-42bf-be1d-90abd4fc7854/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/5140c3b8-cc6e-42bf-be1d-90abd4fc7854",
		  "previews": [
			{
			  "id": "0bb38600-da88-4356-a5e0-6b726a170f92",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 620,
			  "height": 438,
			  "url": "https://medienarchiv.zhdk.ch/media/0bb38600-da88-4356-a5e0-6b726a170f92"
			},
			{
			  "id": "23bd1e05-417f-4aeb-b7d1-7e8c08719427",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 300,
			  "height": 212,
			  "url": "https://medienarchiv.zhdk.ch/media/23bd1e05-417f-4aeb-b7d1-7e8c08719427"
			},
			{
			  "id": "4132fe01-61c8-4c23-b669-dc1c385fb734",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 100,
			  "height": 71,
			  "url": "https://medienarchiv.zhdk.ch/media/4132fe01-61c8-4c23-b669-dc1c385fb734"
			},
			{
			  "id": "4844c8f7-0175-4d32-bb3f-da4c6ad709fc",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 125,
			  "height": 88,
			  "url": "https://medienarchiv.zhdk.ch/media/4844c8f7-0175-4d32-bb3f-da4c6ad709fc"
			},
			{
			  "id": "4e80f623-a3c9-4e1c-b180-c852683f4e71",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 1024,
			  "height": 724,
			  "url": "https://medienarchiv.zhdk.ch/media/4e80f623-a3c9-4e1c-b180-c852683f4e71"
			}
		  ]
		},
		{
		  "id": "c8cabfa4-d4c3-4b15-8c4f-8de1ac240ecb",
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
		  "created_at": "2016-05-25T09:53:26.558Z",
		  "file_id": "c8e6d056-3fff-47c2-abf8-bd159c1c3863",
		  "file_name": "07_meduza_from_beneath_1.jpg",
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/c8e6d056-3fff-47c2-abf8-bd159c1c3863/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/c8e6d056-3fff-47c2-abf8-bd159c1c3863",
		  "previews": [
			{
			  "id": "18b8d8d6-dc2b-442a-b539-c69d09aeac5b",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 333,
			  "height": 500,
			  "url": "https://medienarchiv.zhdk.ch/media/18b8d8d6-dc2b-442a-b539-c69d09aeac5b"
			},
			{
			  "id": "1e35049c-ba24-448e-9862-086d3f23f421",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "maximum",
			  "width": 4000,
			  "height": 6000,
			  "url": "https://medienarchiv.zhdk.ch/media/1e35049c-ba24-448e-9862-086d3f23f421"
			},
			{
			  "id": "49fefe48-6b5b-4dc3-ac13-6781ea94a4bc",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 83,
			  "height": 125,
			  "url": "https://medienarchiv.zhdk.ch/media/49fefe48-6b5b-4dc3-ac13-6781ea94a4bc"
			},
			{
			  "id": "8bc9f028-d747-4f4d-a638-ae1601755e79",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 200,
			  "height": 300,
			  "url": "https://medienarchiv.zhdk.ch/media/8bc9f028-d747-4f4d-a638-ae1601755e79"
			},
			{
			  "id": "a90835e6-b81b-472d-8b17-78daf4c851fa",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 67,
			  "height": 100,
			  "url": "https://medienarchiv.zhdk.ch/media/a90835e6-b81b-472d-8b17-78daf4c851fa"
			},
			{
			  "id": "cdd07ebf-4fd3-4594-9a48-f92ba3e8f107",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 512,
			  "height": 768,
			  "url": "https://medienarchiv.zhdk.ch/media/cdd07ebf-4fd3-4594-9a48-f92ba3e8f107"
			}
		  ]
		},
		{
		  "id": "e281782a-16b1-43d1-a900-9ed8c1cc2b79",
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
		  "created_at": "2016-05-25T09:52:16.998Z",
		  "file_id": "315f85cb-cbaa-4fec-8726-2881d35d4356",
		  "file_name": "05_meduza_attached_to_the_net.jpg",
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/315f85cb-cbaa-4fec-8726-2881d35d4356/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/315f85cb-cbaa-4fec-8726-2881d35d4356",
		  "previews": [
			{
			  "id": "0c4e16e6-4a9d-4801-9a12-bb40733fb8ac",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 1024,
			  "height": 683,
			  "url": "https://medienarchiv.zhdk.ch/media/0c4e16e6-4a9d-4801-9a12-bb40733fb8ac"
			},
			{
			  "id": "2d4e761a-afa4-4d7f-a5dd-aadf86417b29",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 620,
			  "height": 413,
			  "url": "https://medienarchiv.zhdk.ch/media/2d4e761a-afa4-4d7f-a5dd-aadf86417b29"
			},
			{
			  "id": "500d295f-824b-4428-932c-4bda53783572",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 125,
			  "height": 83,
			  "url": "https://medienarchiv.zhdk.ch/media/500d295f-824b-4428-932c-4bda53783572"
			},
			{
			  "id": "78f3d556-5486-4f91-8960-4fc7006f39cf",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 300,
			  "height": 200,
			  "url": "https://medienarchiv.zhdk.ch/media/78f3d556-5486-4f91-8960-4fc7006f39cf"
			},
			{
			  "id": "c85bda52-f3d7-4c54-a05d-1d0f8900388a",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 100,
			  "height": 67,
			  "url": "https://medienarchiv.zhdk.ch/media/c85bda52-f3d7-4c54-a05d-1d0f8900388a"
			},
			{
			  "id": "d04fde81-445a-4e65-b62d-bf3e1b636508",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "maximum",
			  "width": 6000,
			  "height": 4000,
			  "url": "https://medienarchiv.zhdk.ch/media/d04fde81-445a-4e65-b62d-bf3e1b636508"
			}
		  ]
		},
		{
		  "id": "e3e06cbf-db66-45b6-bcd6-2fc78a6d92b1",
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
		  "created_at": "2016-05-25T09:52:50.769Z",
		  "file_id": "9c9ceeba-fdd3-43f6-b7cb-673b4ba010f7",
		  "file_name": "06_strings_tied_to_the_servo.jpg",
		  "stream_url": "https://medienarchiv.zhdk.ch/api/media-files/9c9ceeba-fdd3-43f6-b7cb-673b4ba010f7/data-stream",
		  "download_url": "https://medienarchiv.zhdk.ch/files/9c9ceeba-fdd3-43f6-b7cb-673b4ba010f7",
		  "previews": [
			{
			  "id": "2604f863-7e42-4c09-8ea1-6ac60cf8700b",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "medium",
			  "width": 200,
			  "height": 300,
			  "url": "https://medienarchiv.zhdk.ch/media/2604f863-7e42-4c09-8ea1-6ac60cf8700b"
			},
			{
			  "id": "53194bd8-7ef0-45db-b3df-5ee880f3e867",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "x_large",
			  "width": 512,
			  "height": 768,
			  "url": "https://medienarchiv.zhdk.ch/media/53194bd8-7ef0-45db-b3df-5ee880f3e867"
			},
			{
			  "id": "b343a92a-ed19-42b4-a482-40dd1b8a9efa",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small_125",
			  "width": 83,
			  "height": 125,
			  "url": "https://medienarchiv.zhdk.ch/media/b343a92a-ed19-42b4-a482-40dd1b8a9efa"
			},
			{
			  "id": "db466098-e3bd-47f1-9ca6-9fafc1a070f3",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "small",
			  "width": 67,
			  "height": 100,
			  "url": "https://medienarchiv.zhdk.ch/media/db466098-e3bd-47f1-9ca6-9fafc1a070f3"
			},
			{
			  "id": "dd867dbe-bc6e-43d1-bb18-2f0d9c434427",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "large",
			  "width": 333,
			  "height": 500,
			  "url": "https://medienarchiv.zhdk.ch/media/dd867dbe-bc6e-43d1-bb18-2f0d9c434427"
			},
			{
			  "id": "f0328a6d-198d-4c62-b21a-f19e41ba587c",
			  "type": "image",
			  "content_type": "image/jpeg",
			  "size": "maximum",
			  "width": 4000,
			  "height": 6000,
			  "url": "https://medienarchiv.zhdk.ch/media/f0328a6d-198d-4c62-b21a-f19e41ba587c"
			}
		  ]
		}
	  ]
	}`, string(bytes))
}

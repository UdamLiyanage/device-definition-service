# Device Definition Service
***
This service is reponsible for:
1. Creating device definitions
2. Reading device definitions from database
3. Updating device definitions
4. Delete device definitions

***
## Document structure for Device Definition - Revision 1
```
{
	"_id": "Object ID",
	"version": "Definition Version 1",
	"name": "Device Name",
	"type": "Device Type",
	"organisations": [{
			"organisation_id": "Organisation ID"
		},
		{
			"organisation_id": "Organisation ID"
		}
	],
	"user": [{
			"user_id": "User ID"
		},
		{
			"user_id": "User ID"
		}
	],
	"events": [
		"publish",
		"subscribe"
	],
	"actions": [{
			"type": "Action Type"
		},
		{
			"type": "Action Type"
		}
	],
	"parameters": {
		"param1": "Param 1",
		"param2": "Param 2"
	}
}
```

## Document structure for Device Definition - Revision 2
```
{
	"_id": "Object ID",
	"version": "Definition Version 1",
	"name": "Definition Name",
	"type": "Device Type",
	"public": "true|false",
	"communication": {
		"protocol": "http|mqtt|coap",
		"credentials": {
			"url": "",
			"authentication": {
				"type": "basic",
				"username": "username",
				"password": "password"
			}
		}
	},
	"description": "Description",
	"uid": "User ID",
	"latest_firmware_version": "v1.5",
	"commands": [{
			"name": "command name",
			"type": "Data Type"
		},
		{
			"name": "command name",
			"type": "Data Type"
		}
	],
	"command_format": {
		"format_type": "json|csv",
		"payload": {
			"temp": "Temperature",
			"power": "on"
		}
	},
	"parameters": [{
			"name": "temperature",
			"parameter": "t",
			"data_type": "float"
		},
		{
			"name": "humidity",
			"parameter": "h",
			"data_type": "integer"
		}
	]
}
```

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

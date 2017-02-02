# TownCenter

[![Coverage Status](https://coveralls.io/repos/github/jakelong95/TownCenter/badge.svg?branch=master)](https://coveralls.io/github/jakelong95/TownCenter?branch=master)

TownCenter is the user service for Expresso. It handles registering, updating, listing, and getting users.

## API
### Users
`POST /api/user` creates a new user and adds it to the  database.

Example:
*Request:*
```
POST localhost:8084/api/user
{
	"passHash" : "4hu3qrasdgf34ujasdhfsaf4asdf",
	"firstName" : "First",
	"lastName" : "Last",
	"email" : "email@domain.com",
	"phone" : "124567890",
	"addressLine1" : "Address Line 1",
	"addressLine2" : "Address Line 2",
	"addressCity" : "City",
	"addressState" : "State",
	"addressZip" : "Zip",
	"addrssCountry" : "Country"
}
```

*Response:*
```
{
	"data" : {
		"id" : "86c3d82d-da86-11e6-9d4c-0242ac120004",
		"passHash" : "",
		"firstName" : "First",
		"lastName" : "Last",
		"email" : "email@domain.com",
		"phone" : "124567890",
		"addressLine1" : "Address Line 1",
		"addressLine2" : "Address Line 2",
		"addressCity" : "City",
		"addressState" : "State",
		"addressZip" : "Zip",
		"addrssCountry" : "Country",
		"roasterId" : "",
		"isRoaster" : 0
	}
}
```

#### `GET /api/user?offset=0&limit=20` returns up to `limit` user records starting from `offset` when ordered by userId

Example:

*Request:*
```
GET localhost:8084/api/user?offset=0&limit=20
```

*Response:*
```
{
  "data": [
    {
    	"id" : "86c3d82d-da86-11e6-9d4c-0242ac120004",
		"passHash" : "",
		"firstName" : "First",
		"lastName" : "Last",
		"email" : "email@domain.com",
		"phone" : "124567890",
		"addressLine1" : "Address Line 1",
		"addressLine2" : "Address Line 2",
		"addressCity" : "City",
		"addressState" : "State",
		"addressZip" : "Zip",
		"addrssCountry" : "Country",
		"roasterId" : "",
		"isRoaster" : 0
    }
  ]
}
```

#### `GET /api/user/:userId` returns the user record with the given userID

Example:

*Request:*
```
GET localhost:8084/api/user/86c3d82d-da86-11e6-9d4c-0242ac120004
```

*Response:*
```
{
  "data": {
		"id" : "86c3d82d-da86-11e6-9d4c-0242ac120004",
		"passHash" : "",
		"firstName" : "First",
		"lastName" : "Last",
		"email" : "email@domain.com",
		"phone" : "124567890",
		"addressLine1" : "Address Line 1",
		"addressLine2" : "Address Line 2",
		"addressCity" : "City",
		"addressState" : "State",
		"addressZip" : "Zip",
		"addrssCountry" : "Country",
		"roasterId" : "",
		"isRoaster" : 0
  }
}
```

#### `PUT /api/user/:userId` updates the user record with the given userID to match the provided data. This just overrides values, so anything not present in the request will be set to NULL

Example:

*Request:*
```
PUT localhost:8084/api/user/86c3d82d-da86-11e6-9d4c-0242ac120004
{
    "id" : "86c3d82d-da86-11e6-9d4c-0242ac120004",
	"passHash" : "",
	"firstName" : "First",
	"lastName" : "Last",
	"email" : "email@domain.com",
	"phone" : "124567890",
	"addressLine1" : "Address Line 1",
	"addressLine2" : "Address Line 2",
	"addressCity" : "City",
	"addressState" : "State",
	"addressZip" : "Zip",
	"addrssCountry" : "Country",
	"roasterId" : "",
	"isRoaster" : 0
}
```

*Response:*
```
{
  "data": {
    "id" : "86c3d82d-da86-11e6-9d4c-0242ac120004",
	"passHash" : "",
	"firstName" : "First",
	"lastName" : "Last",
	"email" : "email@domain.com",
	"phone" : "124567890",
	"addressLine1" : "Address Line 1",
	"addressLine2" : "Address Line 2",
	"addressCity" : "City",
	"addressState" : "State",
	"addressZip" : "Zip",
	"addrssCountry" : "Country",
	"roasterId" : "",
	"isRoaster" : 0
  }
}
```

#### `DELETE /api/user/:userId` deletes the user with the given userID
Example:

*Request:*
```
DELETE localhost:8084/api/user/86c3d82d-da86-11e6-9d4c-0242ac120004
```

*Response:*
```
{
  "success": true
}
```

### Roasters

`POST /api/roaster` creates a new roaster and adds it to the  database.

Example:
*Request:*
```
POST localhost:8084/api/roaster
{
	"name" : "Name",
	"email" : "email@domain.com",
	"phone" : "124567890",
	"addressLine1" : "Address Line 1",
	"addressLine2" : "Address Line 2",
	"addressCity" : "City",
	"addressState" : "State",
	"addressZip" : "Zip",
	"addrssCountry" : "Country"
}
```

*Response:*
```
{
	"data" : {
		"id" : "86c3d82d-da86-11e6-9d4c-0242ac120004",
		"name" : "Name",
		"email" : "email@domain.com",
		"phone" : "124567890",
		"addressLine1" : "Address Line 1",
		"addressLine2" : "Address Line 2",
		"addressCity" : "City",
		"addressState" : "State",
		"addressZip" : "Zip",
		"addrssCountry" : "Country"
	}
}
```

#### `GET /api/roaster?offset=0&limit=20` returns up to `limit` roaster records starting from `offset` when ordered by roasterId

Example:
*Request:*
```
GET localhost:8084/api/roaster?offset=0&limit=20
```

*Response:*
```
{
  "data": [
    {
    	"id" : "86c3d82d-da86-11e6-9d4c-0242ac120004",
		"name" : "Name",
		"email" : "email@domain.com",
		"phone" : "124567890",
		"addressLine1" : "Address Line 1",
		"addressLine2" : "Address Line 2",
		"addressCity" : "City",
		"addressState" : "State",
		"addressZip" : "Zip",
		"addrssCountry" : "Country"
    }
  ]
}
```

#### `GET /api/roaster/:roasterId` returns the roaster record with the given roasterId

Example:
*Request:*
```
GET localhost:8084/api/roaster/86c3d82d-da86-11e6-9d4c-0242ac120004
```

*Response:*
```
{
  "data": {
		"id" : "86c3d82d-da86-11e6-9d4c-0242ac120004",
		"name" : "Name",
		"email" : "email@domain.com",
		"phone" : "124567890",
		"addressLine1" : "Address Line 1",
		"addressLine2" : "Address Line 2",
		"addressCity" : "City",
		"addressState" : "State",
		"addressZip" : "Zip",
		"addrssCountry" : "Country"
  }
}
```

#### `PUT /api/roaster/:roasterId` updates the roaster record with the given roasterId to match the provided data. This just overrides values, so anything not present in the request will be set to NULL

Example:
*Request:*
```
PUT localhost:8084/api/roaster/86c3d82d-da86-11e6-9d4c-0242ac120004
{
    "id" : "86c3d82d-da86-11e6-9d4c-0242ac120004",
	"name" : "Name",
	"email" : "email@domain.com",
	"phone" : "124567890",
	"addressLine1" : "Address Line 1",
	"addressLine2" : "Address Line 2",
	"addressCity" : "City",
	"addressState" : "State",
	"addressZip" : "Zip",
	"addrssCountry" : "Country"
}
```

*Response:*
```
{
  "data": {
    "id" : "86c3d82d-da86-11e6-9d4c-0242ac120004",
	"name" : "Name",
	"email" : "email@domain.com",
	"phone" : "124567890",
	"addressLine1" : "Address Line 1",
	"addressLine2" : "Address Line 2",
	"addressCity" : "City",
	"addressState" : "State",
	"addressZip" : "Zip",
	"addrssCountry" : "Country"
  }
}
```

#### `DELETE /api/roaster/:roasterId` deletes the roaster with the given roasterId
Example:
*Request:*
```
DELETE localhost:8084/api/roaster/86c3d82d-da86-11e6-9d4c-0242ac120004
```

*Response:*
```
{
  "success": true
}
```
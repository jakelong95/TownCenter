# TownCenter

TownCenter is the user service for Expresso. It handles registering, updating, listing, and getting users.

## API
####*POST /consumer*

**Parameters:** firstname, lastname, email, phone, password, address_line1, address_line2, address_city, address_zip, address_country

**Action:** Creates a new consumer entry and returns the ID that references the consumer, otherwise returns an error.
___
####*GET /consumer*

**Action:** Returns an array containing all of the registered consumers.
___
####*PATCH /consumer/:consumer_id*

**Parameters:** firstname, lastname, email, phone, password, address_line1, address_line2, address_city, address_zip, address_country, billing, subscription

**Action:** Updates the information for the specified consumer.
___
####*DELETE /consumer/:consumer_id*

**Action:** Deletes the specified consumer.
___
####*GET /consumer/:consumer_id*

**Action:** Returns information about the specified consumer.
___
####*POST /provider*

**Parameters:** name, email, phone, password, address_line1, address_line2, address_city, address_zip, address_country

**Action:** Creates a new provider entry and returns the ID that references that provider, otherwise returns an error.
___
####*GET /provider*

**Action:** Returns an array containing all of the registered providers.
__
####*PATCH /provider/:provider_id*

**Parameters:** name, email, phone, password, address_line1, address_line2, address_city, address_zip, address_country, billing, inventory

**Action:** Updates information for the specified provider.
___
####*DELETE /provider/:provider_id*

**Action:** Deletes the specified provider.
___
####*GET /provider/:provider_id*

**Action:** Returns information about the specified provider.

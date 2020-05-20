# tenprotocol-server

Reference implementation of the Threat Exposure Notification Protocol (10p).

## Creating the Keyspace

```cassandraql
CREATE KEYSPACE tenp WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 1};

CREATE TABLE IF NOT EXISTS tenp.api_keys (
    api_key varchar,
    client_id varchar,
    PRIMARY KEY(api_key)
);

CREATE TABLE IF NOT EXISTS tenp.contacts (
    code varchar,
    ts timestamp,
	reporting_client_id varchar,
	verification_code varchar,
    PRIMARY KEY(code, ts)
);

INSERT INTO tenp.api_keys (api_key, client_id) VALUES ('foo', 'bar');
```

## Uploading your contacts

1. Collect contacts broadcast by other clients
2. Collect a verification_code (for example BX8714198SF872D)
3. Upload contacts:

```shell script
curl -vk -H 'Authorization: APIKEY foo' --data '{"verification_code": "BX8714198SF872D", "contacts": [{"code": "8937498237492", "timestamp": 1586647079}, {"code": "7826487623847", "timestamp": 1586647080}]}' 'https://localhost:5000/contacttracing/contacts'
```

## Checking your contacts

```shell script
curl -vk -H 'Authorization: APIKEY foo' 'https://localhost:5000/contacttracing/contacts?codes=8937498237492&code=78264876238476'
```

response:
```json
{
    "contacts": "2",
    "first_contact": "1586647079",
    "last_contact": "1586647080"
}
```

<!-- DO NOT EDIT: this file is automatically generated using scw-doc-gen -->
# Documentation for `scw baremetal`
Baremetal API
  
- [Operating System (OS) management commands](#operating-system-(os)-management-commands)
  - [Get an OS with a given ID](#get-an-os-with-a-given-id)
  - [List all available OS that can be install on a baremetal server](#list-all-available-os-that-can-be-install-on-a-baremetal-server)
- [Server management commands](#server-management-commands)
  - [Create a baremetal server](#create-a-baremetal-server)
  - [Delete a baremetal server](#delete-a-baremetal-server)
  - [Get a specific baremetal server](#get-a-specific-baremetal-server)
  - [Install a baremetal server](#install-a-baremetal-server)
  - [List baremetal servers](#list-baremetal-servers)
  - [Reboot a baremetal server](#reboot-a-baremetal-server)
  - [Start a baremetal server](#start-a-baremetal-server)
  - [Stop a baremetal server](#stop-a-baremetal-server)
  - [Update a baremetal server](#update-a-baremetal-server)
  - [Wait for a server to reach a stable state (delivery and installation)](#wait-for-a-server-to-reach-a-stable-state-(delivery-and-installation))

  
## Operating System (OS) management commands

An Operating System (OS) is the underlying software installed on your server


### Get an OS with a given ID

Return specific OS for the given ID.

**Usage:**

```
scw baremetal os get <os-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| os-id | Required | ID of the OS |
| zone | Default: `fr-par-1`<br />One of: `fr-par-2` | Zone to target. If none is passed will use default zone from the config |


**Examples:**


Get a specific OS ID
```
scw baremetal os get
```




### List all available OS that can be install on a baremetal server

List all available OS that can be install on a baremetal server.

**Usage:**

```
scw baremetal os list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| offer-id |  | Filter OS by offer ID |
| zone | Default: `fr-par-1`<br />One of: `fr-par-2` | Zone to target. If none is passed will use default zone from the config |



## Server management commands

A server is a denomination of a type of instances provided by Scaleway


### Create a baremetal server

Create a new baremetal server. Once the server is created, you probably want to install an OS.

**Usage:**

```
scw baremetal server create [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| name | Required<br />Default: `<generated>` | Name of the server (≠hostname) |
| description |  | Description associated to the server, max 255 characters |
| type | Default: `GP-BM1-S`<br />One of: `GP-BM1-L`, `GP-BM1-M`, `GP-BM1-S`, `HC-BM1-L`, `HC-BM1-S`, `HM-BM1-XL`, `HM-BM1-M` | Server commercial type |
| tags.{index} |  | Tags to associate to the server |
| organization-id |  | Organization ID to use. If none is passed will use default organization ID from the config |
| zone | Default: `fr-par-1`<br />One of: `fr-par-2` | Zone to target. If none is passed will use default zone from the config |


**Examples:**


Create instance
```
scw baremetal server create
```

Create a GP-BM1-M instance, give it a name and add tags
```
scw baremetal server create name=foo tags.0=prod tags.1=blue type=GP-BM1-M
```




### Delete a baremetal server

Delete the server associated with the given ID.

**Usage:**

```
scw baremetal server delete <server-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| server-id | Required | ID of the server to delete |
| zone | Default: `fr-par-1`<br />One of: `fr-par-2` | Zone to target. If none is passed will use default zone from the config |


**Examples:**


Delete a baremetal server
```
scw baremetal server delete 11111111-1111-1111-1111-111111111111
```




### Get a specific baremetal server

Get the server associated with the given ID.

**Usage:**

```
scw baremetal server get <server-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| server-id | Required | ID of the server |
| zone | Default: `fr-par-1`<br />One of: `fr-par-2` | Zone to target. If none is passed will use default zone from the config |


**Examples:**


Get a given server
```
scw baremetal server get 11111111-1111-1111-1111-111111111111
```




### Install a baremetal server

Install an OS on the server associated with the given ID.

**Usage:**

```
scw baremetal server install <server-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| server-id | Required | Server ID to install |
| os-id | Required | ID of the OS to install on the server |
| hostname | Required | Hostname of the server |
| all-ssh-keys |  | Add all SSH keys on your baremetal instance (cannot be used with ssh-key-ids) |
| ssh-key-ids.{index} | Required | SSH key IDs authorized on the server (cannot be used with all-ssh-keys) |
| zone | Default: `fr-par-1`<br />One of: `fr-par-2` | Zone to target. If none is passed will use default zone from the config |


**Examples:**


Install an OS on a given server with a particular SSH key ID
```
scw baremetal server install 11111111-1111-1111-1111-111111111111 os-id=11111111-1111-1111-1111-111111111111 ssh-key-ids.0=11111111-1111-1111-1111-111111111111
```




### List baremetal servers

List baremetal servers.

**Usage:**

```
scw baremetal server list [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| order-by | One of: `created_at_asc`, `created_at_desc` | Order of the servers |
| tags.{index} |  | Filter servers by tags |
| status.{index} |  | Filter servers by status |
| name |  | Filter servers by name |
| organization-id |  | Filter servers by organization ID |
| zone | Default: `fr-par-1`<br />One of: `fr-par-2` | Zone to target. If none is passed will use default zone from the config |


**Examples:**


List all servers on your default zone
```
scw baremetal server list
```




### Reboot a baremetal server

Reboot the server associated with the given ID, use boot param to reboot in rescue.

**Usage:**

```
scw baremetal server reboot <server-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| server-id | Required | ID of the server to reboot |
| boot-type | Default: `normal`<br />One of: `unknown_boot_type`, `normal`, `rescue` | The type of boot |
| zone | Default: `fr-par-1`<br />One of: `fr-par-2` | Zone to target. If none is passed will use default zone from the config |


**Examples:**


Reboot a server using the same os
```
scw baremetal server reboot 11111111-1111-1111-1111-111111111111
```

Reboot a server in rescue mode
```
scw baremetal server reboot 11111111-1111-1111-1111-111111111111 boot-type=rescue
```




### Start a baremetal server

Start the server associated with the given ID.

**Usage:**

```
scw baremetal server start <server-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| server-id | Required | ID of the server to start |
| zone | Default: `fr-par-1`<br />One of: `fr-par-2` | Zone to target. If none is passed will use default zone from the config |


**Examples:**


Start a baremetal server
```
scw baremetal server start 11111111-1111-1111-1111-111111111111
```




### Stop a baremetal server

Stop the server associated with the given ID.

**Usage:**

```
scw baremetal server stop <server-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| server-id | Required | ID of the server to stop |
| zone | Default: `fr-par-1`<br />One of: `fr-par-2` | Zone to target. If none is passed will use default zone from the config |


**Examples:**


Stop a baremetal server
```
scw baremetal server stop 11111111-1111-1111-1111-111111111111
```




### Update a baremetal server

Update the server associated with the given ID.

**Usage:**

```
scw baremetal server update <server-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| server-id | Required | ID of the server to update |
| name |  | Name of the server (≠hostname), not updated if null |
| description |  | Description associated to the server, max 255 characters, not updated if null |
| tags.{index} |  | Tags associated to the server, not updated if null |
| zone | Default: `fr-par-1`<br />One of: `fr-par-2` | Zone to target. If none is passed will use default zone from the config |



### Wait for a server to reach a stable state (delivery and installation)

Wait for a server to reach a stable state. This is similar to using --wait flag on other action commands, but without requiring a new action on the server.

**Usage:**

```
scw baremetal server wait <server-id ...> [arg=value ...]
```


**Args:**

| Name |   | Description |
|------|---|-------------|
| server-id | Required | ID of the server affected by the action. |
| zone | Default: `fr-par-1` | Zone to target. If none is passed will use default zone from the config |


**Examples:**


Wait for a server to reach a stable state
```
scw baremetal server wait 11111111-1111-1111-1111-111111111111
```





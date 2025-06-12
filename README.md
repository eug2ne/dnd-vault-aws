# DnD Vault Project
Toy project creating a DnD vault for a DnD group with the following features using AWS

- Player authentication
	- Only authenticated players can join the vault
- File upload
	- Players can upload character sheets, character imgs, etc
	- Players can only access their own files
		- DM can access all files uploaded by players
- 1:1 DM chat
	- Players can chat with the DM 1:1
- Plot manager bot
	- DM can create and manage main+side plotline with events in each
		- DM can set state of each event as uninitiated/progressing/completed/terminated
	- Bot will send automatic updates to players when an event is finished
		- Players can look into their process easily

## Current Progress
### In progress
- Fetching data for dm/player

### Finished
- Player/DM auth
  - register+login+logout api
  - authorization middleware

### BUG
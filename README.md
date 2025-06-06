# DnD Vault Project
Creating a DnD vault for a DnD group with the following features using AWS

- Player/DM authentication
  - Only authenticated players/DM can join the vault
- File system
	- Players can upload character sheets, character imgs, etc
	- Players can only access their own files
		- DM can access all files uploaded by players
- 1:1 DM chat
	- Players can 1:1 chat with the DM
- Sidequest manager bot
	- DM can create new sidequests + set the state as uninitialized/progress/end
	- Bot will send automatic message to players when they finish a sidequest

## Current Progress
- Creating middleware to api
- Creating Player/DM auth
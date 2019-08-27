package minecraft

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// GameData is a loose wrapper around a part of the data found in the StartGame packet. It holds data sent
// specifically at the start of the game, such as the position of the player, the game mode, etc.
type GameData struct {
	// WorldName is the name of the world that the player spawns in. This name will be displayed at the top of
	// the player list when opening the in-game menu. It may contain colour codes and does not have to be an
	// actual world name, but instead, can be the server name.
	// If WorldName is left empty, the name of the Listener will be used to show above the player list
	// in-game.
	WorldName string
	// EntityUniqueID is the unique ID of the player. The unique ID is unique for the entire world and is
	// often used in packets. Most servers send an EntityUniqueID equal to the EntityRuntimeID.
	EntityUniqueID int64
	// EntityRuntimeID is the runtime ID of the player. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	EntityRuntimeID uint64
	// PlayerGameMode is the game mode the player currently has. It is a value from 0-4, with 0 being
	// survival mode, 1 being creative mode, 2 being adventure mode, 3 being survival spectator and 4 being
	// creative spectator.
	PlayerGameMode int32
	// PlayerPosition is the spawn position of the player in the world. In servers this is often the same as
	// the world's spawn position found below.
	PlayerPosition mgl32.Vec3
	// Pitch is the vertical rotation of the player. Facing straight forward yields a pitch of 0. Pitch is
	// measured in degrees.
	Pitch float32
	// Yaw is the horizontal rotation of the player. Yaw is also measured in degrees.
	Yaw float32
	// Dimension is the ID of the dimension that the player spawns in. It is a value from 0-2, with 0 being
	// the overworld, 1 being the nether and 2 being the end.
	Dimension int32
	// WorldSpawn is the block on which the world spawn of the world. This coordinate has no effect on the
	// place that the client spawns, but it does have an effect on the direction that a compass points.
	WorldSpawn protocol.BlockPos
	// GameRules defines game rules currently active with their respective values. The value of these game
	// rules may be either 'bool', 'int32' or 'float32'. Some game rules are server side only, and don't
	// necessarily need to be sent to the client.
	GameRules map[string]interface{}
	// Time is the total time that has elapsed since the start of the world.
	Time int64
	// Blocks is a list of all blocks and variants existing in the game. Failing to send any of the blocks
	// that are in the game, including any specific variants of that block, will crash mobile clients. It
	// seems Windows 10 games do not crash.
	Blocks []protocol.BlockEntry
	// Items is a list of all items existing in the game. Failing to send any of the default items that are in
	// the game will crash mobile clients.
	Items []protocol.ItemEntry
}
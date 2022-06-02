package systems

import (
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/entities"
	"github.com/tomknightdev/dwarven-fortresses/enums"
)

type Task struct {
	GameMap GameMap
}

func NewTask(gameMap GameMap) *Task {
	return &Task{
		GameMap: gameMap,
	}
}

func (t *Task) Update(w engine.World) {
	jobs := w.View(components.Task{}).Filter()
	var entitiesToRemove []engine.Entity

	var task *components.Task
	for _, job := range jobs {
		var pos *components.Position
		job.Get(&task, &pos)

		if task.Completed {
			switch task.InputModeEnum {
			case enums.InputModeChop:
				ent, ok := w.GetEntity(task.EntityID)
				if !ok {
					panic("entity not found")
				}

				var drop *components.Drops
				ent.Get(&drop)

				for i := 0; i < drop.DropCount; i++ {
					w.AddEntities(&entities.Resource{
						Position: *pos,
						Sprite:   components.NewSprite(assets.Images["log0"]),
						Resource: components.NewResource(),
					})
				}

				entitiesToRemove = append(entitiesToRemove, ent)
			case enums.InputModeBuild:
				w.AddEntities(&entities.Tile{
					Position: *pos,
					Sprite:   components.NewSprite(assets.Images["stairdown"]),
					TileType: components.NewTileType(enums.TileTypeStairDown),
				})
				w.AddEntities(&entities.Tile{
					Position: components.NewPosition(pos.X, pos.Y, pos.Z-1),
					Sprite:   components.NewSprite(assets.Images["stairup"]),
					TileType: components.NewTileType(enums.TileTypeStairUp),
				})
				index := t.GameMap.GetTileByTypeIndexFromPos(enums.TileTypeRock, components.NewPosition(pos.X, pos.Y, pos.Z-1))
				t.GameMap.UpdateTile(enums.TileTypeRock, index, enums.TileTypeRockFloor)
			}

			entitiesToRemove = append(entitiesToRemove, job)
		}
	}

	for _, job := range entitiesToRemove {
		w.RemoveEntity(job)
	}
}

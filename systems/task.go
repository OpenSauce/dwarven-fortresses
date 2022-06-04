package systems

import (
	"log"

	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/entities"
	"github.com/tomknightdev/dwarven-fortresses/enums"
	"github.com/tomknightdev/dwarven-fortresses/helpers"
)

type Task struct {
}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) Update(w engine.World) {
	jobs := w.View(components.Task{}).Filter()

	if len(jobs) == 0 {
		return
	}

	gms, found := w.View(components.GameMapSingleton{}).Get()
	if !found {
		panic("game map singleton not found")
	}

	var gmComp *components.GameMapSingleton
	gms.Get(&gmComp)

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
					w.AddEntities(&entities.Item{
						Position: *pos,
						Sprite:   components.NewSprite(assets.Images["log0"]),
						Item:     components.NewItem(),
					})
				}

				entitiesToRemove = append(entitiesToRemove, ent)
			case enums.InputModeBuild:
				w.AddEntities(&entities.Building{
					Position: *pos,
					Sprite:   components.NewSprite(assets.Images["stairdown"]),
					TileType: components.NewTileType(enums.TileTypeStairDown),
					Building: components.NewBuilding(),
				})
				gmComp.TilesByType[enums.TileTypeStairDown] = append(gmComp.TilesByType[enums.TileTypeStairDown], *pos)

				index, err := helpers.GetTileByTypeIndexFromPos(components.NewPosition(pos.X, pos.Y, pos.Z-1), gmComp.TilesByType[enums.TileTypeRock])
				if err != nil {
					log.Println(err)
					entitiesToRemove = append(entitiesToRemove, job)
					continue
				}

				gmComp.TilesToUpdateChan <- struct {
					FromTileType enums.TileTypeEnum
					ToTileType   enums.TileTypeEnum
					TileIndex    int
				}{

					enums.TileTypeRock,
					enums.TileTypeRockFloor,
					index,
				}

				w.AddEntities(&entities.Building{
					Position: components.NewPosition(pos.X, pos.Y, pos.Z-1),
					Sprite:   components.NewSprite(assets.Images["stairup"]),
					TileType: components.NewTileType(enums.TileTypeStairUp),
					Building: components.NewBuilding(),
				})
				gmComp.TilesByType[enums.TileTypeStairUp] = append(gmComp.TilesByType[enums.TileTypeStairUp], components.NewPosition(pos.X, pos.Y, pos.Z-1))

			case enums.InputModeMine:
				index, err := helpers.GetTileByTypeIndexFromPos(components.NewPosition(pos.X, pos.Y, pos.Z),
					gmComp.TilesByType[enums.TileTypeRock])
				if err != nil {
					log.Println(err)
					entitiesToRemove = append(entitiesToRemove, job)
					continue
				}

				gmComp.TilesToUpdateChan <- struct {
					FromTileType enums.TileTypeEnum
					ToTileType   enums.TileTypeEnum
					TileIndex    int
				}{
					enums.TileTypeRock,
					enums.TileTypeRockFloor,
					index,
				}
			}

			entitiesToRemove = append(entitiesToRemove, job)
		}
	}

	for _, job := range entitiesToRemove {
		w.RemoveEntity(job)
	}
}

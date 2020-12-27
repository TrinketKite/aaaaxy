package aaaaaa

import (
	"fmt"
	"io"

	"github.com/fardog/tmx"
)

// level is a parsed form of a loaded level.
type level struct {
	tiles map[Pos]*levelTile
}

// levelTile is a single tile in the level.
type levelTile struct {
	tile       Tile
	spawnables []*spawnable
	warpzone   *warpzone
}

// warpzone represents a warp tile. Whenever anything enters this tile, it gets
// moved to "to" and the direction transformed by "transform". For the game to
// work, every warpzone must be paired with an exact opposite elsewhere. This
// is ensured at load time.
type warpzone struct {
	toTile    Pos
	transform Orientation
}

type rawWarpzone struct {
	startTile, endTile Pos
	orientation        Orientation
}

type spawnable struct {
	// Entity ID. Used to decide what needs spawning. Unique within a level.
	id EntityID
	// Entity type. Used to spawn it on demand.
	entityType  string
	levelPos    Pos
	posInTile   Delta
	size        Delta
	orientation Orientation
}

func LoadLevel(r io.Reader) (*level, error) {
	t, err := tmx.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("invalid map: %v", err)
	}
	if t.Orientation != "orthogonal" {
		return nil, fmt.Errorf("unsupported map: got orientation %q, want orthogonal", t.Orientation)
	}
	if t.TileWidth != TileSize || t.TileHeight != TileSize {
		return nil, fmt.Errorf("unsupported map: got tile size %dx%d, want %dx%d", t.TileWidth, t.TileHeight, TileSize, TileSize)
	}
	if len(t.TileSets) != 1 {
		return nil, fmt.Errorf("unsupported map: got %d embedded tilesets, want 1", len(t.TileSets))
	}
	if len(t.Layers) != 1 {
		return nil, fmt.Errorf("unsupported map: got %d layers, want 1", len(t.Layers))
	}
	if len(t.ImageLayers) != 0 {
		return nil, fmt.Errorf("unsupported map: got %d image layers, want 0", len(t.ImageLayers))
	}
	tds, err := t.Layers[0].TileDefs(t.TileSets)
	if err != nil {
		return nil, fmt.Errorf("invalid map layer: %v", err)
	}
	level := level{}
	for i, td := range tds {
		pos := Pos{Y: i / t.Layers[0].Width, X: i % t.Layers[0].Width}
		orientation := Identity()
		if td.HorizontallyFlipped {
			orientation = FlipX().Concat(orientation)
		}
		if td.VerticallyFlipped {
			orientation = FlipY().Concat(orientation)
		}
		if td.DiagonallyFlipped {
			orientation = FlipD().Concat(orientation)
		}
		solid, err := td.Tile.Properties.Bool("solid")
		if err != nil {
			return nil, fmt.Errorf("invalid map: could not parse solid: %v", err)
		}
		level.tiles[pos] = &levelTile{
			tile: Tile{
				Solid:       solid,
				levelPos:    pos,
				image:       nil, // CachePic(td.Tile.Image),
				orientation: orientation,
			},
		}
	}
	warpzones := map[string][]rawWarpzone{}
	for _, og := range t.ObjectGroups {
		for _, o := range og.Objects {
			startTile := Pos{X: int(o.X) / TileSize, Y: int(o.Y) / TileSize}
			endTile := Pos{X: int(o.X+o.Width-1) / TileSize, Y: int(o.Y+o.Height-1) / TileSize}
			orientation := Identity()
			orientationProp := o.Properties.WithName("orientation")
			if orientationProp != nil {
				orientation, err = ParseOrientation(orientationProp.Value)
				if err != nil {
					return nil, fmt.Errorf("invalid orientation: %v", err)
				}
			}
			if o.Type == "warpzone" {
				// Warpzones must be paired by name.
				// Consider encoding their orientation by a tile name? Check what Tiled supports best.
				// Or maybe require a warp tile below the warpzone and lookup there?
				warpzones[o.Name] = append(warpzones[o.Name], rawWarpzone{
					startTile:   startTile,
					endTile:     endTile,
					orientation: orientation,
				})
			}
			delta := Delta{DX: int(o.X) % TileSize, DY: int(o.Y) % TileSize}
			// TODO: support orientations.
			ent := spawnable{
				id:          EntityID(o.ObjectID),
				entityType:  o.Type,
				levelPos:    startTile,
				posInTile:   delta,
				size:        Delta{DX: int(o.Width), DY: int(o.Height)},
				orientation: orientation,
			}
			for y := startTile.Y; y <= endTile.Y; y++ {
				for x := startTile.X; x <= endTile.X; x++ {
					pos := Pos{X: x, Y: y}
					level.tiles[pos].spawnables = append(level.tiles[pos].spawnables, &ent)
				}
			}
		}
	}
	for warpname, warppair := range warpzones {
		if len(warppair) != 2 {
			return nil, fmt.Errorf("unpaired warpzone %q: got %d, want 2", warpname, len(warppair))
		}
		for a := 0; a < 2; a++ {
			from := warppair[a]
			to := warppair[1-a]
			// Warp orientation: right = direction to walk the warp, down = orientation (for mirroring).
			transform := to.orientation.Concat(from.orientation.Inverse()).Concat(TurnAround())
			fromCenter2 := from.startTile.Add(from.endTile.Delta(Pos{}))
			toCenter2 := to.startTile.Add(to.endTile.Delta(Pos{}))
			for fromy := from.startTile.Y; fromy <= from.endTile.Y; fromy++ {
				for fromx := from.startTile.X; fromx <= from.endTile.X; fromx++ {
					fromPos := Pos{X: fromx, Y: fromy}
					fromPos2 := fromPos.Add(fromPos.Delta(Pos{}))
					toPos2 := toCenter2.Add(transform.Apply(fromPos2.Delta(fromCenter2)))
					toPos := toPos2.Scale(1, 2).Add(to.orientation.Apply(East()))
					level.tiles[fromPos].warpzone = &warpzone{
						toTile:    toPos,
						transform: transform,
					}
				}
			}
		}
	}
	return &level, nil
}

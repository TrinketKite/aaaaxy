module github.com/divVerent/aaaaxy

go 1.16

require (
	github.com/adrg/xdg v0.3.3
	github.com/fardog/tmx v0.0.0-20210504210836-02c45f261672
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/hajimehoshi/ebiten/v2 v2.2.0-alpha.12
	github.com/hajimehoshi/oto/v2 v2.0.0 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2
	golang.org/x/exp v0.0.0-20210831011441-b422fde0612a // indirect
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d
	golang.org/x/sys v0.0.0-20210831042530-f4d43177bf5e
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

// Use when playing around with extended/fixed tmx versions.
// replace github.com/fardog/tmx => github.com/divVerent/tmx v0.0.0-20210504110059-b8d75006ad02

// For debugging:
// replace github.com/hajimehoshi/ebiten/v2 => /home/rpolzer/src/ebiten
// replace github.com/go-gl/glfw/v3.3/glfw => /home/rpolzer/src/go-gl-glfw/v3.3/glfw

package toy

import (
	"github.com/go-gl/mathgl/mgl32"
	"time"
)

type shaderVars struct {
	startTime         time.Time
	lastRender        time.Time
	resolution        mgl32.Vec3
	frameNumber       int32
	channelTime       [4]float32
	channelResolution [4]mgl32.Vec3
	mouse             mgl32.Vec4
	date              mgl32.Vec4
	movieMode         bool
	movieTime         float32
}

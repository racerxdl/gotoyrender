package toy

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	defaultUniformValues = map[string]interface{}{
		"iResolution": &mgl32.Vec3{0, 0, 1},
		"iTime":       float32(0),
		"iTimeDelta":  float32(0),
		"iFrame":      int32(0),

		"iChannel0": int32(0),
		"iChannel1": int32(1),
		"iChannel2": int32(2),
		"iChannel3": int32(3),

		"iChannelResolution[0]": &mgl32.Vec3{0, 0, 1},
		"iChannelResolution[1]": &mgl32.Vec3{0, 0, 1},
		"iChannelResolution[2]": &mgl32.Vec3{0, 0, 1},
		"iChannelResolution[3]": &mgl32.Vec3{0, 0, 1},

		"iChannelTime[0]": float32(0),
		"iChannelTime[1]": float32(0),
		"iChannelTime[2]": float32(0),
		"iChannelTime[3]": float32(0),

		"iMouse": &mgl32.Vec4{0, 0, 0, 0},
		"iDate":  &mgl32.Vec4{0, 0, 0, 0},
	}

	defaultVertexFormat = glhf.AttrFormat{
		{Name: "position", Type: glhf.Vec2},
	}
)

/*

	startTime         time.Time
	lastRender        time.Time
	resolution        mgl32.Vec3
	frameNumber       int32
	channelTime       [4]float32
	channelResolution [4]mgl32.Vec3
	mouse             mgl32.Vec4
	date              mgl32.Vec4
*/

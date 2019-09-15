package toy

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/mathgl/mgl32"
)

var (
	defaultUniformValues = map[string]interface{}{
		"iTime":       float32(0),
		"iResolution": &mgl32.Vec3{0, 0, 1},
		"iChannel0":   int32(0),
		"iChannel1":   int32(1),
		"iChannel2":   int32(2),
		"iChannel3":   int32(3),
	}

	defaultVertexFormat = glhf.AttrFormat{
		{Name: "position", Type: glhf.Vec2},
	}
)

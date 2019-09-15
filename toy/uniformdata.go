package toy

import (
	"github.com/faiface/glhf"
	"github.com/go-gl/mathgl/mgl32"
)

type uniformData struct {
	glhf.Attr
	Id        int
	value     interface{}
	IsPointer bool
}

func (ud *uniformData) Value() interface{} {
	if !ud.IsPointer {
		return ud.value
	}
	switch ud.Type {
	case glhf.Vec2:
		return *ud.value.(*mgl32.Vec2)
	case glhf.Vec3:
		return *ud.value.(*mgl32.Vec3)
	case glhf.Vec4:
		return *ud.value.(*mgl32.Vec4)
	case glhf.Mat2:
		return *ud.value.(*mgl32.Mat2)
	case glhf.Mat23:
		return *ud.value.(*mgl32.Mat2x3)
	case glhf.Mat24:
		return *ud.value.(*mgl32.Mat2x4)
	case glhf.Mat3:
		return *ud.value.(*mgl32.Mat3)
	case glhf.Mat32:
		return *ud.value.(*mgl32.Mat3x2)
	case glhf.Mat34:
		return *ud.value.(*mgl32.Mat3x4)
	case glhf.Mat4:
		return *ud.value.(*mgl32.Mat4)
	case glhf.Mat42:
		return *ud.value.(*mgl32.Mat4x2)
	case glhf.Mat43:
		return *ud.value.(*mgl32.Mat4x3)
	case glhf.Int:
		return *ud.value.(*int32)
	case glhf.Float:
		return *ud.value.(*float32)
	default:
		panic("invalid attrtype")
	}
}

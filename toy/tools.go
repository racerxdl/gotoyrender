package toy

import (
	"fmt"
	"github.com/faiface/glhf"
	"github.com/go-gl/mathgl/mgl32"
	"reflect"
)

func getAttrType(v interface{}) (glhf.AttrType, bool, error) {
	switch v.(type) {
	case int32:
		return glhf.Int, false, nil
	case float32:
		return glhf.Float, false, nil
	case mgl32.Vec2:
		return glhf.Vec2, false, nil
	case mgl32.Vec3:
		return glhf.Vec3, false, nil
	case mgl32.Vec4:
		return glhf.Vec4, false, nil
	case mgl32.Mat2:
		return glhf.Mat2, false, nil
	case mgl32.Mat2x3:
		return glhf.Mat23, false, nil
	case mgl32.Mat2x4:
		return glhf.Mat24, false, nil
	case mgl32.Mat3:
		return glhf.Mat3, false, nil
	case mgl32.Mat3x2:
		return glhf.Mat32, false, nil
	case mgl32.Mat3x4:
		return glhf.Mat34, false, nil
	case mgl32.Mat4:
		return glhf.Mat4, false, nil
	case mgl32.Mat4x2:
		return glhf.Mat42, false, nil
	case mgl32.Mat4x3:
		return glhf.Mat43, false, nil
	case *mgl32.Vec2:
		return glhf.Vec2, true, nil
	case *mgl32.Vec3:
		return glhf.Vec3, true, nil
	case *mgl32.Vec4:
		return glhf.Vec4, true, nil
	case *mgl32.Mat2:
		return glhf.Mat2, true, nil
	case *mgl32.Mat2x3:
		return glhf.Mat23, true, nil
	case *mgl32.Mat2x4:
		return glhf.Mat24, true, nil
	case *mgl32.Mat3:
		return glhf.Mat3, true, nil
	case *mgl32.Mat3x2:
		return glhf.Mat32, true, nil
	case *mgl32.Mat3x4:
		return glhf.Mat34, true, nil
	case *mgl32.Mat4:
		return glhf.Mat4, true, nil
	case *mgl32.Mat4x2:
		return glhf.Mat42, true, nil
	case *mgl32.Mat4x3:
		return glhf.Mat43, true, nil
	case *int32:
		return glhf.Int, true, nil
	case *float32:
		return glhf.Float, true, nil
	default:
		return glhf.Int, false, fmt.Errorf("invalid type %s", reflect.TypeOf(v))
	}
}

package toy

import (
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/mathgl/mgl32"
	"image"
	"time"
)

const baseFragmentShader = `
#version 330 core

out vec4 fragColor;

uniform vec3 iResolution;
uniform float iTime;

uniform sampler2D iChannel0;
uniform sampler2D iChannel1;
uniform sampler2D iChannel2;
uniform sampler2D iChannel3;

vec4 texture(sampler2D s, vec2 c) { return texture2D(s,c); }
void mainImage( out vec4 fragColor, in vec2 fragCoord );

void main() {
    fragColor = vec4(0.0, 0.0, 0.0, 1.0);
	mainImage(fragColor, gl_FragCoord.xy);
    fragColor.w = 1.0;
}`

const defaultVertexShader = `
#version 330 core

in vec2 position;

void main() {
	gl_Position = vec4(position, 0.0, 1.0);
}
`

const defaultFragmentShader = `
void mainImage( out vec4 fragColor, in vec2 fragCoord ) {}
`

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

type Render struct {
	shader      *glhf.Shader
	screen      *glhf.VertexSlice
	uniforms    map[string]*uniformData
	uniformList glhf.AttrFormat
	startTime   time.Time
	resolution  mgl32.Vec3
	shaderDirty bool
	texChannels []*Texture

	fragmentShaderPiece string
}

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

func MakeToyRender(width, height int) *Render {
	tr := &Render{
		uniforms:    map[string]*uniformData{},
		uniformList: glhf.AttrFormat{},
		startTime:   time.Now(),
		resolution:  mgl32.Vec3{float32(width), float32(height), 1},
		shaderDirty: true,
		texChannels: make([]*Texture, 4),
	}

	tr.SetFragmentShaderPiece(defaultFragmentShader)

	for k, v := range defaultUniformValues {
		tr.SetUniformValue(k, v)
	}

	tr.updateShader()

	return tr
}

func (tr *Render) SetFragmentShaderPiece(main string) {
	tr.fragmentShaderPiece = baseFragmentShader + main
	tr.shaderDirty = true
}

func (tr *Render) SetUniformValue(name string, value interface{}) error {
	t, p, err := getAttrType(value)
	if err != nil {
		return err
	}

	added := tr.addUniform(name, t)

	if added {
		tr.shaderDirty = true
	}

	tr.uniforms[name].IsPointer = p
	tr.uniforms[name].value = value

	return nil
}

func (tr *Render) SetTextureData(n int, img *image.NRGBA) {
	b := img.Bounds()
	if tr.texChannels[n] == nil { // Create one
		mainthread.Call(func() {
			tr.texChannels[n] = NewTexture(b.Max.X, b.Max.Y, true, img.Pix)
		})
		return
	}

	t := tr.texChannels[n]

	if t.Width() == b.Max.X && t.Height() == b.Max.Y { // Existing at same size, so just replace contents
		mainthread.Call(func() {
			t.SetPixels(0, 0, b.Max.X, b.Max.Y, img.Pix)
		})
		return
	}

	// Exists but different size
	mainthread.Call(func() {
		tr.texChannels[n] = NewTexture(b.Max.X, b.Max.Y, true, img.Pix)
	})
}

func (tr *Render) Render() {
	tr.updateShader()

	mainthread.Call(func() {
		iTime := float32(time.Since(tr.startTime).Seconds())
		tr.SetUniformValue("iTime", iTime)
		tr.SetUniformValue("iResolution", &tr.resolution)

		tr.SetUniformValue("iChannel0", int32(0))
		tr.SetUniformValue("iChannel1", int32(1))
		tr.SetUniformValue("iChannel2", int32(2))
		tr.SetUniformValue("iChannel3", int32(3))

		// Clear the window.
		glhf.Clear(0, 0, 0, 1)

		// Here we Begin/End all necessary objects and finally draw the vertex
		// slice.
		tr.shader.Begin()
		tr.setUniforms()

		for _, v := range tr.texChannels {
			if v != nil {
				v.Begin()
			}
		}

		tr.screen.Begin()
		tr.screen.Draw()
		tr.screen.End()

		for _, v := range tr.texChannels {
			if v != nil {
				v.End()
			}
		}

		tr.shader.End()
	})
}

func (tr *Render) addUniform(name string, t glhf.AttrType) bool {
	if tr.uniforms[name] != nil {
		// Already exists
		return false
	}

	att := glhf.Attr{
		Name: name,
		Type: t,
	}

	tr.uniformList = append(tr.uniformList, att)

	tr.uniforms[name] = &uniformData{
		Attr:  att,
		Id:    len(tr.uniformList) - 1,
		value: nil,
	}

	return true
}

func (tr *Render) setUniforms() {
	for _, v := range tr.uniforms {
		tr.shader.SetUniformAttr(v.Id, v.Value())
	}
}

func (tr *Render) updateShader() {
	if !tr.shaderDirty {
		return
	}

	mainthread.Call(func() {
		var err error

		// Here we create a shader. The second arudment is the format of the uniform
		// attributes. Since our shader has no uniform attributes, the format is empty.
		tr.shader, err = glhf.NewShader(defaultVertexFormat, tr.uniformList, defaultVertexShader, tr.fragmentShaderPiece)

		if err != nil {
			panic(err)
		}

		tr.screen = glhf.MakeVertexSlice(tr.shader, 6, 6)

		tr.screen.Begin()

		tr.screen.SetVertexData([]float32{
			-1, -1,
			+1, -1,
			+1, +1,
			-1, -1,
			+1, +1,
			-1, +1,
		})

		tr.screen.End()
	})

	tr.shaderDirty = false
}

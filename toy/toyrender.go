package toy

import (
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/faiface/pixel"
	"github.com/go-gl/mathgl/mgl32"
	"image"
	"image/color"
	"time"
)

type Render struct {
	frame               *glhf.Frame
	shader              *glhf.Shader
	screen              *glhf.VertexSlice
	uniforms            map[string]*uniformData
	uniformList         glhf.AttrFormat
	startTime           time.Time
	resolution          mgl32.Vec3
	shaderDirty         bool
	texChannels         []*Texture
	fragmentShaderPiece string
	sprite              *pixel.Sprite
	bounds              pixel.Rect
	pixels              []uint8
	pixelsDirty         bool
}

func MakeToyRender(width, height int) *Render {
	tr := &Render{
		uniforms:    map[string]*uniformData{},
		uniformList: glhf.AttrFormat{},
		startTime:   time.Now(),
		resolution:  mgl32.Vec3{float32(width), float32(height), 1},
		shaderDirty: true,
		texChannels: make([]*Texture, 4),
		bounds:      pixel.R(0, 0, float64(width), float64(height)),
		pixelsDirty: true,
	}

	tr.sprite = pixel.NewSprite(tr, tr.Bounds())

	mainthread.Call(func() {
		tr.frame = glhf.NewFrame(width, height, true)
	})

	tr.SetFragmentShaderPiece(defaultFragmentShader)

	for k, v := range defaultUniformValues {
		tr.SetUniformValue(k, v)
	}

	tr.updateShader()

	return tr
}

// Bounds returns the rectangular bounds of the Canvas.
func (tr *Render) Bounds() pixel.Rect {
	return tr.bounds
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
		tr.frame.Begin()
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
		tr.frame.End()
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

func (tr *Render) ScreenShot() image.Image {
	tex := tr.frame.Texture()
	img := image.NewNRGBA(image.Rect(0, 0, tex.Width(), tex.Height()))

	pixels := tr.Pixels()
	copy(img.Pix, pixels)

	return img
}

// Draw draws the content of the Canvas onto another Target, transformed by the given Matrix, just
// like if it was a Sprite containing the whole Canvas.
func (tr *Render) Draw(t pixel.Target, matrix pixel.Matrix) {
	tr.sprite.Draw(t, matrix)
}

// DrawColorMask draws the content of the Canvas onto another Target, transformed by the given
// Matrix and multiplied by the given mask, just like if it was a Sprite containing the whole Canvas.
//
// If the color mask is nil, a fully opaque white mask will be used causing no effect.
func (tr *Render) DrawColorMask(t pixel.Target, matrix pixel.Matrix, mask color.Color) {
	tr.sprite.DrawColorMask(t, matrix, mask)
}

// Texture returns the underlying OpenGL Texture of this Canvas.
//
// Implements GLPicture interface.
func (tr *Render) Texture() *glhf.Texture {
	return tr.frame.Texture()
}

// Frame returns the underlying OpenGL Frame of this Canvas.
func (tr *Render) Frame() *glhf.Frame {
	return tr.frame
}

// Pixels returns an alpha-premultiplied RGBA sequence of the content of the Canvas.
func (tr *Render) Pixels() []uint8 {
	if tr.pixelsDirty {
		mainthread.Call(func() {
			tex := tr.frame.Texture()
			tex.Begin()
			tr.pixels = tex.Pixels(0, 0, tex.Width(), tex.Height())
			tex.End()
		})
		tr.pixelsDirty = false
	}

	return tr.pixels
}

func (tr *Render) Color(at pixel.Vec) pixel.RGBA {
	p := int((at.X * 4) + at.Y*tr.bounds.Max.X*4)
	if p > len(tr.pixels)-4 {
		return pixel.RGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 1,
		}
	}

	return pixel.RGBA{
		R: float64(tr.pixels[p]) / 256,
		G: float64(tr.pixels[p+1]) / 256,
		B: float64(tr.pixels[p+2]) / 256,
		A: float64(tr.pixels[p+3]) / 256,
	}
}

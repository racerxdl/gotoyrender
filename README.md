# Go ToyRender
ShaderToy like Go-Lang Render compatible with [Pixel](https://github.com/faiface/pixel) Sprites

It still WIP and doesn't present many of the ShaderToy features. But it does run some nice shaders.


### Usage

```go
package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/racerxdl/gotoyrender/toy"
	"golang.org/x/image/colornames"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Sample",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tr := toy.MakeToyRender(720, 480)
	tr.SetFragmentShaderPiece(`
        void mainImage( out vec4 fragColor, in vec2 fragCoord ) {
            vec2 position = 2. * (fragCoord.xy / iResolution.xy) - 1.;
            vec3 colour = vec3(0.0);
            float density = 0.15;
            float amplitude = 0.3;
            float frequency = 5.0;
            float scroll = 0.4;
            
            colour += vec3(0.1, 0.05, 0.05) * (1.0 / abs((position.y + (amplitude * sin(((position.x-0.0) + iTime * scroll) *frequency)))) * density);
            colour += vec3(0.05, 0.1, 0.05) * (1.0 / abs((position.y + (amplitude * sin(((position.x-0.3) + iTime * scroll) *frequency)))) * density);
            colour += vec3(0.05, 0.05, 0.1) * (1.0 / abs((position.y + (amplitude * sin(((position.x-0.6) + iTime * scroll) *frequency)))) * density);
            //
            fragColor = vec4( colour, 1.0 );
        }
    `)

	//tr.SetTextureData(0, img) Sets the iChannel0 texture to img

	for !win.Closed() {
		tr.Render()

		win.Clear(colornames.Wheat)

		tr.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

```

### API

*TODO*

### ShaderBasis (cmd/shaderbasis)

![ShaderBasis](https://user-images.githubusercontent.com/578310/64926847-df410980-d7d8-11e9-8a73-a4ce35599861.jpg)

### Star based on flight404 (cmd/star)

![Star](https://user-images.githubusercontent.com/578310/64926846-dcdeaf80-d7d8-11e9-800b-5edd4f2d3466.jpg)

### Neon Parallax by @stormoid (cmd/neonparalax)

![Neon Parallax](https://user-images.githubusercontent.com/578310/64926918-f59b9500-d7d9-11e9-8912-89f9c1712d03.jpg)

### Digital Brain by srtuss (cmd/digitalbrain)

![Digital Brain](https://user-images.githubusercontent.com/578310/64927045-0c42eb80-d7dc-11e9-8417-034edd9b0885.jpg)


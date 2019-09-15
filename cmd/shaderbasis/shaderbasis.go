package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/racerxdl/gotoyrender/toy"
	"golang.org/x/image/colornames"
)

const fragmentToy = `
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
`

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "ShaderBasis",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	tr := toy.MakeToyRender(720, 480)
	tr.SetFragmentShaderPiece(fragmentToy)

	//tr.Render()
	//
	//s := tr.ScreenShot()
	//
	//f, _ := os.Create("screenshot.jpg")
	//
	//jpeg.Encode(f, s, &jpeg.Options{
	//	Quality: 100,
	//})
	//
	//f.Close()

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

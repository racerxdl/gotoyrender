package main

import (
	"github.com/faiface/glhf"
	"github.com/faiface/mainthread"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/racerxdl/gotoyrender/toy"
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
	var win *glfw.Window

	defer func() {
		mainthread.Call(func() {
			glfw.Terminate()
		})
	}()

	mainthread.Call(func() {
		glfw.Init()

		glfw.WindowHint(glfw.ContextVersionMajor, 3)
		glfw.WindowHint(glfw.ContextVersionMinor, 3)
		glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
		glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
		glfw.WindowHint(glfw.Resizable, glfw.False)

		var err error

		win, err = glfw.CreateWindow(1280, 720, "Shader Basis", nil, nil)
		if err != nil {
			panic(err)
		}

		win.MakeContextCurrent()

		glhf.Init()
	})

	width, height := win.GetFramebufferSize()

	tr := toy.MakeToyRender(width, height)
	tr.SetFragmentShaderPiece(fragmentToy)

	shouldQuit := false
	for !shouldQuit {
		mainthread.Call(func() {
			if win.ShouldClose() {
				shouldQuit = true
			}

			// Clear the window.
			glhf.Clear(0, 0, 0, 1)
		})

		tr.Render()

		mainthread.Call(func() {
			win.SwapBuffers()
			glfw.PollEvents()
		})
	}
}

func main() {
	mainthread.Run(run)
}

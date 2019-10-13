package toy

const baseFragmentShader = `
#version 330 core

out vec4 fragColor;

uniform vec3 iResolution;           // viewport resolution (in pixels)
uniform float iTime;                // shader playback time (in seconds)
uniform float iTimeDelta;           // render time (in seconds)
uniform int iFrame;                 // shader playback frame

uniform sampler2D iChannel0;        // input channel. 00 = 2D/Cube
uniform sampler2D iChannel1;        // input channel. 01 = 2D/Cube
uniform sampler2D iChannel2;        // input channel. 02 = 2D/Cube
uniform sampler2D iChannel3;        // input channel. 03 = 2D/Cube

uniform vec3 iChannelResolution[4]; // channel resolution (in pixels)
uniform float iChannelTime[4];      // channel playback time (in seconds)

uniform vec4      iDate;            // (year, month, day, time in seconds)
uniform vec4      iMouse;           // mouse pixel coords. xy: current (if MLB down), zw: click

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

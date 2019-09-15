package toy

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

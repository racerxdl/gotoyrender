ShaderBasis Example
============

Same example from [https://github.com/DangerCenter/shaderbasis](https://github.com/DangerCenter/shaderbasis)

![ShaderBasis](https://user-images.githubusercontent.com/578310/64926847-df410980-d7d8-11e9-8a73-a4ce35599861.jpg)


Shader
======

```glsl
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
```
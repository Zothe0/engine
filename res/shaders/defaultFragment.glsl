#version 460 core
out vec4 fragColor;

in vec2 texCoord;

uniform sampler2D drawTexture;

void main()
{
	fragColor = texture(drawTexture, texCoord);
}
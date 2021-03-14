#version 410 core
out vec4 fragColor;

in vec3 oColorCoord;

void main()
{
	fragColor = vec4(oColorCoord, 1);
}
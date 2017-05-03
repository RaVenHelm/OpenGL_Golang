#version 330 core

layout(location = 0) in vec3 vert;

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

out vec3 color;

void main(){
    gl_Position = projection * camera * model * vec4(vert, 1.0);
    color = vert;
}

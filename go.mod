module github.com/unitoftime/flow-examples

go 1.18

//replace github.com/faiface/pixel => /home/jacob/go/src/github.com/unitoftime/pixel
replace github.com/unitoftime/ecs => /home/jacob/go/src/github.com/unitoftime/ecs

replace github.com/unitoftime/flow => /home/jacob/go/src/github.com/unitoftime/flow

//replace github.com/unitoftime/packer => /home/jacob/go/src/github.com/unitoftime/packer
//replace github.com/unitoftime/glitch => /home/jacob/go/src/github.com/unitoftime/glitch
//replace github.com/unitoftime/gl => /home/jacob/go/src/github.com/unitoftime/gl
//replace github.com/unitoftime/glfw => /home/jacob/go/src/github.com/unitoftime/glfw

require (
	github.com/faiface/pixel v0.10.0
	github.com/ungerik/go3d v0.0.0-20211026193542-07217314a07d
	github.com/unitoftime/ecs v0.0.0-20220105184333-0c45de2067df
	github.com/unitoftime/flow v0.0.0-20211123140059-6d204cb8a4fa
	github.com/unitoftime/glitch v0.0.0-20211229194034-9f3f6025ad6c
)

require (
	github.com/faiface/glhf v0.0.0-20211013000516-57b20770c369 // indirect
	github.com/faiface/mainthread v0.0.0-20171120011319-8b78f0a41ae3 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20211213063430-748e38ca8aec // indirect
	github.com/go-gl/mathgl v1.0.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20211111143520-d0d5ecc1a356 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/unitoftime/gl v0.0.0-20211214004148-0b2bf5e8079d // indirect
	github.com/unitoftime/glfw v0.0.0-20211214004214-579be7093d5c // indirect
	github.com/unitoftime/packer v0.0.0-20220105185326-f541e031de11 // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	honnef.co/go/js/dom v0.0.0-20210725211120-f030747120f2 // indirect
)

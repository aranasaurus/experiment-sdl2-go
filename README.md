experiment-sdl2-go
==================

This repo is to track me trying out [SDL2-go](https://github.com/veandco/sdl2-go).

To get sdl2-go to compile on my mac I had to add the following environment variables:

* CGO_CFLAGS="-I /usr/local/include -I /opt/X11/include"
* CGO_LDFLAGS="-F /Library/Frameworks"

I built SDL2 (and the extension libraries) from /usr/local/src the standard `configure`, `make`, `make install`
way, which installs the headers to /usr/local/include and the libraries to /usr/local/lib. I also have the SDL*.frameworks
installed in /Library/Frameworks. It seems to need both, I couldn't figure out the right mix of compiler/linker
flags to get it to build against solely the stuff in /usr/local/* or solely the stuff in /Library/Frameworks, so
__that's__ why the mix :P.

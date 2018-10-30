# Windows
## Downloads
1. Mingw-builds http://mingw-w64.org/doku.php/download/mingw-builds 
    - Version: latest (at time of writing 6.3.0)
    - Architecture: x86_x64
    - Threads: I selected win32 and it worked
    - Exception: I selected seh and it worked
    - Build revision: I selected 1 and it worked
    - Destination Folder: Select a Folder that your Windows User owns, 
  like something under %USERPROFILE%, otherwise you'll have plenty of UAC popups / run everything as admin
2. SDL2 http://libsdl.org/download-2.0.php
    - Current Runtime Binaries for win32-x64 
    - Current Development Libraries for mingw
3. SDL_image, mixer and ttf https://www.libsdl.org/projects/
    - Current Runtime Binaries for win32-x64 
    - Current Development Libraries for mingw
4. (Optional) Nice tool to edit environment variables https://www.rapidee.com/
5. (Optional) 7-zip http://www.7-zip.de/
6. (Optional) Super cool cmd alternative https://conemu.github.io/

## Building go-sdl2 
1. Extract the Runtime Binaries from 2. and 3. to the lib-Folder inside your mingw64-Folder from 1.
2. Extract the x86_64-w64-mingw32-Folder from the Development Libraries into your mingw64-Folder, everything from x86_64-w64-mingw32/bin needs to get into mingw64\bin, etc.
3. Add mingw64/bin and mingw64/lib to your path
4. Run conemu or cmd
    - Set the CGO_CFLAGS env variable, so gcc finds the SDL2 headers with set: ```set CGO_CFLAGS=-ID:\dev\mingw64\mingw64\include```
    - go get go-sdl2 -> ```go get -v github.com/veandco/go-sdl2/sdl```
    - ```go get -v github.com/veandco/go-sdl2/mix```
    - ```go get -v github.com/veandco/go-sdl2/ttf```
    - ```go get -v github.com/veandco/go-sdl2/img```
5. Run a example from go-sdl2 ...it should work

# Other OS
Look at the go-sdl2 install guide -> github.com/veandco/go-sdl2/

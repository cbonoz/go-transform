Go-Transform:
---
Transforms your input image into a spiraling away gif.

### How to use
* Specify your input in the INPUT_FILE constant (ex: static/input.jpg)
* go build
* ./go-transform
* See your output in out.gif - resizes to 256x256 (configurable).


### Gif Example:

<b>Input:</b><br/>
![](https://github.com/cbonoz/go-transform/blob/master/static/input.jpg)
<br/>

<b>Output:</b><br/>
![](https://github.com/cbonoz/go-transform/blob/master/out.gif)
<br/>


Packages:

unicode
net
image
bytes
encoding

Staple Pantry:
strings
math
errors
time
testing
flag
io

Rotate:
https://github.com/disintegration/imaging/blob/a5858022df0e1734a59f973fffe3f87b51c087ed/transform.go
### ACME PORTAL COMPANY - HTML3D

---

HTML3D is a command-line utility for converting 3D modeled geometries into HTML/CSS
representations

### USAGE
---
***path :*** required. path to triangularized model in .obj format
```
    html3d [ path ]
```
***output :*** filename for output files. defaults to name of .obj file
```
    html3d [ path ] -o [ output ]
```
***className :*** name of shared class for faces in generated HTML/CSS. defaults to 'face'
```
    html3d [ path ] -c [ className ]
```

### OUTPUTS
---
***output.html :*** contains 3D object faces
```
    <svg class="face" id="f-0"></svg>
```
***output.css :*** contains face-defining geometry data
```
    #f-0{
	    -webkit-clip-path: polygon(45.40% 53.71%,50.05% 53.71%,54.60% 46.29%);
	    clip-path: polygon(45.40% 53.71%,50.05% 53.71%,54.60% 46.29%);
    }
```

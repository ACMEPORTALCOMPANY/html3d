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
***class :*** name of shared class for faces in generated HTML/CSS. defaults to 'face'
```
    html3d [ path ] -class [ class ]
```
***fill :*** SVG face fill color. defaults to 'none'
```
    html3d [ path ] -fill [ fill ]
```
***output :*** filename for output files. defaults to name of .obj file
```
    html3d [ path ] -output [ output ]
```
***size :*** size of SVG canvas. defaults to 200
```
    html3d [ path ] -size [ size ]
```
***stroke :*** SVG stroke color. defaults to 'black'
```
    html3d [ path ] -stroke [ stroke ]
```

### OUTPUTS

---

***output.html :*** contains 3D object faces
```
    <svg viewBox="0 0 500 500">
	    <polygon class="face" id="f-0" points="37.76,382.12, 462.23,382.12, 293.30,117.87" fill="none" stroke="black" />
    </svg>
```
***output.css :*** contains transformation data
```
    .face {
        position: absolute;
    }

    #f-0 {
        transform: rotate3D(-0.00, 1.00, 0.00, 2.50rad);
    }
```

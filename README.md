### ACME PORTAL COMPANY - HTML3D

---

HTML3D is a command-line utility for converting 3D modeled geometries into HTML/CSS
representations

### HTML3d - ***arguments***
1. path ***[required]:*** path to triangularized model in .obj format
```
    html3d [ path ]
```
2. output: filename for output files. defaults to name of .obj file
```
    html3d [ path ] -o [ output ]
```
3. className: name of shared class for faces in generated HTML/CSS
```
    - ex. html3d [ path ] -c [ className ]
```

### HTML3d - ***outputs***
HTML3D produces 2 files
1. output.html: contains 3D object and faces
2. output.css: contains face-defining geometry data, 3D transformation data

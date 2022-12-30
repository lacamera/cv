# Curriculum Vitae
This resume is built from [YAML structured data](https://yaml.org/) and generated using Gos' [html/template](https://pkg.go.dev/html/template) syntax.  
To generate the `resume.html` simply run:
```sh
go run gen.go
```
To generate the PDF make sure you have [pandoc](https://pandoc.org/) and a LaTeX distribution installed (such as [TeX Live](https://tug.org/texlive/)) then go ahead and run:
```sh
pandoc -so resume.pdf resume.html
```
You may find it useful to recreate this setup for your own resumes.

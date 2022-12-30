# Curriculum Vitae
This resume is built from [YAML structured data](https://yaml.org/) and generated using Gos' [html/template](https://pkg.go.dev/html/template) syntax.  
To generate the `resume.html` simply run:
```sh
go run gen.go -c en.yaml resume.tmpl
```

To compile the HTML to PDF make sure you have [pandoc](https://pandoc.org/) and a LaTeX distribution installed (such as [TeX Live](https://tug.org/texlive/)) then go ahead and run:
```sh
pandoc -so resume.pdf resume.html
```
pandoc only knows how to convert the HTML, not the stylesheets however.
You can use your browser to generate the document (see Printing), dont forget to disable any extra settings.

You may find it useful to recreate this setup for your own resumes.

## Usage
```sh
gen.go [-o output] [-c yaml] /path/to/tmpl
```

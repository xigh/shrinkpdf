# shrinkpdf

A simple tool to extract and reduce the size of scanned pdf document.

# usage :

```bash
go build -v
shrinkpdf path_to_file.pdf
```

## returns :

```bash
$ ./shrinkpdf ../../assurance.pdf 
../../assurance.pdf: 9538388 bytes
page 1: image: 2463x3459
page 2: image: 2463x3459
page 3: image: 2463x3459
page 4: image: 2463x3459
page 5: image: 2463x3459
page 6: image: 2463x3459
page 7: image: 2463x3459
page 8: image: 2463x3459
page 9: image: 2463x3459
page 10: image: 2463x3459
page 11: image: 2463x3459
page 12: image: 2463x3459
saving "assurance.pdf"
assurance.pdf: 4492628 bytes
```

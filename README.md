# Is It Python?

This application scrapes the internet (stack overflow) looking for snippets of python (in <code> blocks), and then run it in an `eval`, and reports if the python snippet successfully ran

TODO:
* add security

notes:

scraper.go
    Scrape(sourl string) ([]links, []Snippets)

eval.go
    Eval()

runtime.go
    Scrapper
    Web

main.go

api.go
    grpc-web
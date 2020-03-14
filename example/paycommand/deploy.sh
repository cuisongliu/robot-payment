docker run --rm -v /Users/fanux/work/src/github.com/fanux/robot:/go/src/github.com/fanux/robot \
  -w /go/src/github.com/fanux/robot golang:1.12.7  go build -o bootstrap example/paycommand/faas.go
rm bootstrap
mv ../../bootstrap .
zip code.zip bootstrap
fun deploy -t template_pro.yml

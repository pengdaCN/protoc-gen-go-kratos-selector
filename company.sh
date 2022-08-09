#! /bin/bash
cd ..

mkdir -p company/protoc-gen-go-kratos-selector

for f in company/protoc-gen-go-kratos-selector/*; do
	rm -rf "$f"
done

for f in protoc-gen-go-kratos-selector/*; do
	cp -r "$f" company/protoc-gen-go-kratos-selector/
done

cd company/protoc-gen-go-kratos-selector

sed -i -e 's,github.com/pengdacn/protoc-gen-go-kratos-selector,gitlab.test.com/common/protoc-gen-go-kratos-selector,g' go.mod tmpl.go.tmpl selector/selector.proto README.md
find . -type f -name '*.go' \
  -exec sed -i -e 's,github.com/pengdacn/protoc-gen-go-kratos-selector,gitlab.test.com/common/protoc-gen-go-kratos-selector,g' {} \;
  
go mod tidy

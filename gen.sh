#!/bin/sh

for file in ./json/*; do
	f=$(basename "$file")
	name="${f%.*}"
	lower=$(echo "$name" | awk '{print tolower($0)}')
	gojson -input="$file" -o="gen_${lower}.go" -pkg="gozelle" -name="$name"
done

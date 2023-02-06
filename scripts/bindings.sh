#!/bin/bash -ex
#
# author:       Steve Huguenin-Elie <steve@brotherday.one>
# description:  Generate go bindings from Solidity ABI and JSON files
#
# online help:
#               https://www.baeldung.com/linux/bash-string-manipulation
#               https://medium.com/mkdir-awesome/case-transformation-in-bash-and-posix-with-examples-acdc1e0d0bc4
#

cd "$(dirname $0)"
cd ../contracts_abi

for filename in `find -type f`; do
        # preprocesing
        filename=${filename%%.}
        filename=${filename##*./}

        file=${filename##*/}
        suffix=${filename##*.}
        file=${file%.$suffix}
        path=${filename%/$file.$suffix}

        if [ "$suffix" == json ] || [ "$suffix" == abi ]; then
                rm ../client/${path}/${file,,}.go
                mkdir -p ../client/${path}
                touch ../client/${path}/${file,,}.go
                if [ "$suffix" == json ]; then
                        abigen --combined-json "${path:-.}/${file}.json" --pkg client --type ${file^} --out "../client/${path}/${file,,}.go"
                else
                        abigen --abi "${path:-.}/${file}.abi" --pkg client --type ${file^} --out "../client/"${path}"/${file,,}.go"
                fi
                echo "Added binding $(realpath ../client/'${path}'/${file,,}.go)."
        fi
done

unset filename
unset file
unset suffix

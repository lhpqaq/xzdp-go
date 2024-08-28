#!/bin/bash
if [ ! -f conf/test/conf.yaml ]; then
    cp conf/test/conf.example.yaml conf/test/conf.yaml
    echo "conf/test/conf.yaml has been created from conf/test/conf.example.yaml"
fi

go build -v xzdp
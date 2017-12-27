#! /bin/bash

tar cvf secrets.tar client-secret.json cms/app.yaml && \
travis encrypt-file secrets.tar --add

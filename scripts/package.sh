#!/bin/bash

npm install
npm run build
cd main
statik -src=../dist

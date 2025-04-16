#!/bin/bash

mkdir -p report

ghz \
  --proto ./proto/your.proto \
  --call yourpackage.YourService.Method \
  -d '{"param": "value"}' \
  -c 50 -n 5000 \
  --format html \
  --output ./report/$(date +%Y%m%d_%H%M)_report.html \
  --insecure \
  127.0.0.1:50051

echo "报告已生成：report/$(ls -t report | head -n 1)"

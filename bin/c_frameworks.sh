#!/bin/bash
echo "### C frameworks statistics rating"
echo ""
./ghstat -r civetweb/civetweb,jorisvink/kore,davidmoreno/onion,lpereira/lwan,emweb/wt -f stats/c_frameworks.csv
echo "[Detailed C frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/c_frameworks.csv)"
echo ""
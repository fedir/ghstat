#!/bin/bash
echo "## Groovy frameworks statistics rating"
echo ""
./ghstat -r gaelyk/gaelyk,kdabir/glide,javaConductor/gserv,grails/grails-core -f stats/groovy_frameworks.csv
echo "[Detailed Groovy frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/groovy_frameworks.csv)"
echo ""
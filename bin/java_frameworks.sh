#!/bin/bash
echo "### Java frameworks statistics rating"
echo ""
./ghstat -r grails/grails-core,playframework/playframework,vaadin/framework,lets-blade/blade,ninjaframework/ninja,bootique/bootique,jooby-project/jooby,pippo-java/pippo,ratpack/ratpack,primefaces/primefaces,netty/netty -f stats/java_frameworks.csv
echo "[Detailed Java frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/java_frameworks.csv)"
echo ""



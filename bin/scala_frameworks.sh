#!/bin/bash
echo "## Scala frameworks statistics rating"
echo ""
./ghstat -r analogweb/analogweb-scala,mesosphere/chaos,tumblr/colossus,twitter/finatra,lift/framework,dvarelap/peregrine,splink/pagelets,nafg/reactive,scalatra/scalatra,skinny-framework/skinny-framework,suzaku-io/suzaku,unfiltered/unfiltered,xitrum-framework/xitrum-new,outr/youi,ThoughtWorksInc/Binding.scala,UdashFramework/udash-core,widok/widok -f stats/scala_frameworks.csv
echo "[Detailed Scala frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/scala_frameworks.csv)"
echo ""

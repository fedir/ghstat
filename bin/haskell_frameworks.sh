#!/bin/bash
echo "### Haskell frameworks statistics rating"
echo ""
./ghstat -r yesodweb/yesod,snapframework/snap-core,agrafix/Spock,transient-haskell/axiom,myfreeweb/magicbane,positiondev/fn -f stats/haskell_frameworks.csv
echo "[Detailed Haskell frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/haskell_frameworks.csv)"
echo ""


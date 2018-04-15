#!/bin/bash
echo "## Crystal frameworks statistics rating"
echo ""
./ghstat -r amberframework/amber,kemalcr/kemal,jasonl99/lattice-core,luckyframework/lucky,vladfaust/prism,samueleaton/raze,spider-gazelle/spider-gazelle -f stats/crystal_frameworks.csv
echo "[Detailed Crystal frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/crystal_frameworks.csv)"
echo ""
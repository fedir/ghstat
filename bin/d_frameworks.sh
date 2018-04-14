#!/bin/bash
echo "### D frameworks statistics rating"
echo ""
./ghstat -r huntlabs/hunt,vibe-d/vibe.d,adamdruppe/arsd,DiamondMVC/Diamond -f stats/d_frameworks.csv
echo "[Detailed D frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/d_frameworks.csv)"
echo ""
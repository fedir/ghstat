#!/bin/bash
echo "### Elixir frameworks statistics rating"
echo ""
./ghstat -r kittoframework/kitto,phoenixframework/phoenix,slogsdon/placid,AntonFagerberg/rackla,AgilionApps/relax,sugar-framework/sugar,hexedpackets/trot -f stats/elixir_frameworks.csv
echo "[Detailed Elixir frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/elixir_frameworks.csv)"
echo ""
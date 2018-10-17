#!/bin/bash
echo "## Go microsevice toolkits statistics rating"
echo ""
./ghstat -r koding/kite,nytimes/gizmo,micro/go-micro,rsms/gotalk,gocircuit/circuit,go-kit/kit -f stats/go_microservice_toolkits.csv
echo "[Detailed Go microsevice toolkits statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/go_microservice_toolkits.csv)"
echo ""
#!/bin/bash
echo "## Python frameworks statistics rating"
echo ""
./ghstat -r bottlepy/bottle,plotly/dash,django/django,pallets/flask,Pylons/pyramid,channelcat/sanic,tornadoweb/tornado,web2py/web2py,TurboGears/tg2 -f stats/python_frameworks.csv
echo "[Detailed Python frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/python_frameworks.csv)"
echo ""
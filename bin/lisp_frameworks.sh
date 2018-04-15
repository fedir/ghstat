#!/bin/bash
echo "## Lisp frameworks statistics rating"
echo ""
./ghstat -r fukamachi/caveman,hargettp/hh-web,fukamachi/ningle,Shirakumo/radiance,eudoxia0/lucerne,joaotavora/snooze -f stats/lisp_frameworks.csv
echo "[Detailed Lisp frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/lisp_frameworks.csv)"
echo ""
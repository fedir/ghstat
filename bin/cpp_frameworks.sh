#!/bin/bash
echo "## C++ frameworks statistics rating"
echo ""
./ghstat -r ipkn/crow,cutelyst/cutelyst,oktal/pistache,jlaine/qdjango,treefrogframework/treefrog-framework -f stats/cpp_frameworks.csv
echo "[Detailed C++ frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/cpp_frameworks.csv)"
echo ""
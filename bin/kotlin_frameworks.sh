#!/bin/bash
echo "### Kotlin frameworks statistics rating"
echo ""
./ghstat -r ktorio/ktor,TinyMission/kara,http4k/http4k,jean79/yested,wasabifx/wasabi,kohesive/kovert,danneu/kog,hypercube1024/firefly -f stats/kotlin_frameworks.csv
echo "[Detailed Kotlin frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/kotlin_frameworks.csv)"
echo ""

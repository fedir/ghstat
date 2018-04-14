#!/bin/bash
echo "### Ruby frameworks statistics rating"
echo ""
./ghstat -r camping/camping,soveran/cuba,patriciomacadden/hobbit,hanami/hanami,ruby-hyperloop/hyper-react,padrino/padrino-framework,pakyow/pakyow,rack-app/rack-app,ramaze/ramaze,jeremyevans/roda,rails/rails,wardrop/Scorched,sinatra/sinatra,voltrb/volt -f stats/ruby_frameworks.csv
echo "[Detailed Ruby frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/ruby_frameworks.csv)"
echo ""
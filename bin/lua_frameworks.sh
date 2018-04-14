#!/bin/bash
echo "### Lua frameworks statistics rating"
echo ""
./ghstat -r leafo/lapis,sailorproject/sailor,keplerproject/orbit,luvit/luvit,Fizzadar/Luapress,kernelsauce/turbo,mongrel2/Tir,appwilldev/moochine -f stats/lua_frameworks.csv
echo "[Detailed Lua frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/lua_frameworks.csv)"
echo ""

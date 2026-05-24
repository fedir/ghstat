#!/bin/bash
echo "## Go frameworks statistics rating"
echo ""
./ghstat -r gin-gonic/gin,gofiber/fiber,labstack/echo,go-chi/chi,beego/beego,gohugoio/hugo,gobuffalo/buffalo,revel/revel,kataras/iris,go-macaron/macaron -f stats/go_frameworks.csv
echo "[Detailed Go frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/go_frameworks.csv)"
echo ""



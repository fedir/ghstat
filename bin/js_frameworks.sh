#!/bin/bash
echo "## JS frameworks statistics rating"
echo ""
./ghstat -r angular/angular,facebook/react,vuejs/vue,\
jquery/jquery,emberjs/ember.js,jashkenas/backbone,\
meteor/meteor,ractivejs/ractive,\
knockout/knockout,hyperapp/hyperapp,developit/preact,\
expressjs/express,hexojs/hexo,hyperapp/hyperapp,MithrilJS/mithril.js,\
totaljs/framework,enyojs/enyo,microjs/microjs.com,\
nuxt/nuxt.js,riot/riot,balderdashy/sails -f stats/js_frameworks.csv
echo "[Detailed JS frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/js_frameworks.csv)"
echo ""
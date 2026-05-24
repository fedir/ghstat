#!/bin/bash
echo "## JS frameworks statistics rating"
echo ""
./ghstat -r \
facebook/react,vuejs/vue,angular/angular,sveltejs/svelte,solidjs/solid,\
expressjs/express,fastify/fastify,nestjs/nest,koajs/koa,honojs/hono,\
vercel/next.js,nuxt/nuxt,remix-run/remix,sveltejs/kit,analogjs/analog \
-f stats/js_frameworks.csv
echo "[Detailed JS frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/js_frameworks.csv)"
echo ""

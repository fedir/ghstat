# ghstat

[![Build Status](https://travis-ci.org/fedir/ghstat.svg?branch=master)](https://travis-ci.org/fedir/ghstat)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/fedir/ghstat/badges/quality-score.png?b=master)](https://scrutinizer-ci.com/g/fedir/ghstat/?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/fedir/ghstat)](https://goreportcard.com/report/github.com/fedir/ghstat)
[![Maintainability](https://api.codeclimate.com/v1/badges/572b4413f5c5ebf49e36/maintainability)](https://codeclimate.com/github/fedir/go-github-statistics/maintainability)
[![codecov](https://codecov.io/gh/fedir/ghstat/branch/master/graph/badge.svg)](https://codecov.io/gh/fedir/ghstat)
[![GoDoc](https://godoc.org/github.com/fedir/ghstat?status.svg)](https://godoc.org/github.com/fedir/ghstat)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Statistical multi-criteria decision-making comparator for selected Github's projects.

Usage example with statistics of Go open source web frameworks, maintained on Github :

    go build
    mkdir tmp
    export GH_USR="your_gh_username" && export GH_PASS="your_gh_api_token"
    ./ghstat

Usage example to compare most famous JS frameworks

    ./ghstat -r angular/angular,facebook/react,vuejs/vue

Usage example to compare most famous PHP frameworks

    ./ghstat -r laravel/framework,symfony/symfony,yiisoft/yii2,bcit-ci/CodeIgniter

After that, `result.csv` file will be created (or updated, if it's already exists) with the statistics of selected repositories.

## Comparaison methodology

At the moment We choosed following metrics, here they are, in alphabetical order :

* Active forkers percentage - more is better
* Age in days - newest is better :)
* Closed issues, % - more is better
* Watchers - more is better
* Total commits - more is better
  * More precisely, it's total commits by existing contributors, commits of deleted accounts, will not be taken in account

## Ratings

### Go web frameworks rating

* The most popular project is `gohugoio/hugo`
* The newest project is `kataras/iris`
* The project with more commits is `gohugoio/hugo`
* The project with more tags is `labstack/echo`
* The project made by most notable top contributors is `astaxie/beego`
* The project with best errors resolving rate is `kataras/iris`
* The project with more commits by day is `gohugoio/hugo`
* The project with the most active community is `gobuffalo/buffalo`
* The best project (taking in account placements in all competitions) is `gohugoio/hugo`

[Detailed Go frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/go_frameworks.csv)

### Python web frameworks rating

* The most popular project is `pallets/flask`
* The newest project is `channelcat/sanic`
* The project with more commits is `django/django`
* The project with more tags is `django/django`
* The project made by most notable top contributors is `pallets/flask`
* The project with best errors resolving rate is `pallets/flask`
* The project with more commits by day is `django/django`
* The project with the most active community is `TurboGears/tg2`
* The best project (taking in account placements in all competitions) is `django/django`

[Detailed Python frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/python_frameworks.csv)

### Ruby web frameworks rating

* The most popular project is `rails/rails`
* The newest project is `ruby-hyperloop/hyper-react`
* The project with more commits is `rails/rails`
* The project with more tags is `rails/rails`
* The project made by most notable top contributors is `rails/rails`
* The project with best errors resolving rate is `padrino/padrino-framework`
* The project with more commits by day is `rails/rails`
* The project with the most active community is `Ramaze/ramaze`
* The best project (taking in account placements in all competitions) is `padrino/padrino-framework`

[Detailed Ruby frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/ruby_frameworks.csv)

### Clojure web frameworks rating

* The most popular project is `weavejester/compojure`
* The newest project is `fulcrologic/fulcro`
* The project with more commits is `fulcrologic/fulcro`
* The project with more tags is `juxt/yada`
* The project made by most notable top contributors is `fulcrologic/fulcro`
* The project with best errors resolving rate is `fulcrologic/fulcro`
* The project with more commits by day is `fulcrologic/fulcro`
* The project with the most active community is `fulcrologic/fulcro`
* The best project (taking in account placements in all competitions) is `fulcrologic/fulcro`

[Detailed Clojure frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/clojure_frameworks.csv)

### Erlang web frameworks rating

* The most popular project is `ninenines/cowboy`
* The newest project is `synrc/n2o`
* The project with more commits is `zotonic/zotonic`
* The project with more tags is `zotonic/zotonic`
* The project made by most notable top contributors is `mochi/mochiweb`
* The project with best errors resolving rate is `synrc/n2o`
* The project with more commits by day is `zotonic/zotonic`
* The project with the most active community is `kivra/giallo`
* The best project (taking in account placements in all competitions) is `synrc/n2o`

[Detailed Erlang frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/erlang_frameworks.csv)

### Haskell frameworks rating

* The most popular project is `yesodweb/yesod`
* The newest project is `myfreeweb/magicbane`
* The project with more commits is `yesodweb/yesod`
* The project with more tags is `yesodweb/yesod`
* The project made by most notable top contributors is `yesodweb/yesod`
* The project with best errors resolving rate is `snapframework/snap-core`
* The project with more commits by day is `yesodweb/yesod`
* The project with the most active community is `transient-haskell/axiom`
* The best project (taking in account placements in all competitions) is `yesodweb/yesod`

[Detailed Haskell frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/haskell_frameworks.csv)

### Lua frameworks rating

* The most popular project is `luvit/luvit`
* The newest project is `sailorproject/sailor`
* The project with more commits is `luvit/luvit`
* The project with more tags is `luvit/luvit`
* The project made by most notable top contributors is `luvit/luvit`
* The project with best errors resolving rate is `keplerproject/orbit`
* The project with more commits by day is `leafo/lapis`
* The project with the most active community is `kernelsauce/turbo`
* The best project (taking in account placements in all competitions) is `luvit/luvit`

[Detailed Lua frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/lua_frameworks.csv)

### Rust frameworks rating

* The most popular project is `iron/iron`
* The newest project is `actix/actix-web`
* The project with more commits is `actix/actix-web`
* The project with more tags is `SergioBenitez/Rocket`
* The project made by most notable top contributors is `iron/iron`
* The project with best errors resolving rate is `Ogeon/rustful`
* The project with more commits by day is `actix/actix-web`
* The project with the most active community is `tomaka/rouille`
* The best project (taking in account placements in all competitions) is `actix/actix-web`

[Detailed Rust frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/rust_frameworks.csv)

### C frameworks rating

* The most popular project is `lpereira/lwan`
* The newest project is `civetweb/civetweb`
* The project with more commits is `civetweb/civetweb`
* The project with more tags is `emweb/wt`
* The project made by most notable top contributors is `lpereira/lwan`
* The project with best errors resolving rate is `jorisvink/kore`
* The project with more commits by day is `civetweb/civetweb`
* The project with the most active community is `civetweb/civetweb`
* The best project (taking in account placements in all competitions) is `civetweb/civetweb`

[Detailed C frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/c_frameworks.csv)

### C++ frameworks rating

* The most popular project is `ipkn/crow`
* The newest project is `oktal/pistache`
* The project with more commits is `cutelyst/cutelyst`
* The project with more tags is `cutelyst/cutelyst`
* The project made by most notable top contributors is `treefrogframework/treefrog-framework`
* The project with best errors resolving rate is `treefrogframework/treefrog-framework`
* The project with more commits by day is `cutelyst/cutelyst`
* The project with the most active community is `cutelyst/cutelyst`
* The best project (taking in account placements in all competitions) is `treefrogframework/treefrog-framework`

[Detailed C++ frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/cpp_frameworks.csv)

### Java frameworks rating

* The most popular project is `netty/netty`
* The newest project is `bootique/bootique`
* The project with more commits is `vaadin/framework`
* The project with more tags is `vaadin/framework`
* The project made by most notable top contributors is `netty/netty`
* The project with best errors resolving rate is `grails/grails-core`
* The project with more commits by day is `vaadin/framework`
* The project with the most active community is `ratpack/ratpack`
* The best project (taking in account placements in all competitions) is `grails/grails-core`

[Detailed Java frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/java_frameworks.csv)

### JS frameworks rating

* The most popular project is `facebook/react`
* The newest project is `hyperapp/hyperapp`
* The project with more commits is `meteor/meteor`
* The project with more tags is `meteor/meteor`
* The project made by most notable top contributors is `facebook/react`
* The project with best errors resolving rate is `meteor/meteor`
* The project with more commits by day is `meteor/meteor`
* The project with the most active community is `riot/riot`
* The best project (taking in account placements in all competitions) is `meteor/meteor`

[Detailed JS frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/js_frameworks.csv)

### PHP frameworks rating

* The most popular project is `symfony/symfony`
* The newest project is `nova-framework/framework`
* The project with more commits is `symfony/symfony`
* The project with more tags is `nova-framework/framework`
* The project made by most notable top contributors is `laravel/framework`
* The project with best errors resolving rate is `laravel/framework`
* The project with more commits by day is `laravel/framework`
* The project with the most active community is `nova-framework/framework`
* The best project (taking in account placements in all competitions) is `laravel/framework`

[Detailed PHP frameworks statistics](https://github.com/fedir/ghstat/blob/master/stats/php_frameworks.csv)

### Cross-language frameworks rating

* The most popular project is `facebook/react`
* The newest project is `actix/actix-web`
* The project with more commits is `rails/rails`
* The project with more tags is `meteor/meteor`
* The project made by most notable top contributors is `facebook/react`
* The project with best errors resolving rate is `kataras/iris`
* The project with more commits by day is `rails/rails`
* The project with the most active community is `Ramaze/ramaze`
* The best project (taking in account placements in all competitions) is `meteor/meteor`

[Detailed cross-language frameworks rating](https://github.com/fedir/ghstat/blob/master/stats/all_frameworks.csv)

### PHP CMS statistics rating

* The most popular project is `WordPress/WordPress`
* The newest project is `neos/neos-development-collection`
* The project with more commits is `drupal/drupal`
* The project with more tags is `TYPO3/TYPO3.CMS`
* The project made by most notable top contributors is `sulu/sulu-standard`
* The project with best errors resolving rate is `bolt/bolt`
* The project with more commits by day is `concrete5/concrete5`
* The project with the most active community is `spip/SPIP`
* The best project (taking in account placements in all competitions) is `bolt/bolt`

[Detailed PHP CMS statistics](https://github.com/fedir/ghstat/blob/master/stats/php_cms.csv)

### Java CMS statistics rating

* The most popular project is `liferay/liferay-portal`
* The newest project is `gentics/mesh`
* The project with more commits is `liferay/liferay-portal`
* The project with more tags is `nuxeo/nuxeo`
* The project made by most notable top contributors is `liferay/liferay-portal`
* The project with best errors resolving rate is `Softmotions/ncms`
* The project with more commits by day is `liferay/liferay-portal`
* The project with the most active community is `Softmotions/ncms`
* The best project (taking in account placements in all competitions) is `gentics/mesh`

[Detailed Java CMS statistics](https://github.com/fedir/ghstat/blob/master/stats/java_cms.csv)
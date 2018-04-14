#!/bin/bash
echo "### PHP frameworks statistics rating"
echo ""
./ghstat -r laravel/framework,symfony/symfony,yiisoft/yii2,bcit-ci/CodeIgniter,nova-framework/framework,cakephp/cakephp,pradosoft/prado,phalcon/cphalcon,nette/nette,PHPixie/Project,slimphp/Slim,leocavalcante/siler -f stats/php_frameworks.csv
echo "[Detailed PHP frameworks frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/php_frameworks.csv)"
echo ""
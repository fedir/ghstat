echo "### Cross-language CMS rating"
echo ""
./ghstat -r \
dotCMS/core,alkacon/opencms-core,gentics/mesh,Softmotions/ncms,liferay/liferay-portal,\
bogeblad/infoglue,nuxeo/nuxeo,lutece-platform/lutece-core,alkacon/opencms-core,exoplatform/ecms,\
Victoire/victoire,backbee/backbee-php,bolt/bolt,concrete5/concrete5,contao/core,\
forkcms/forkcms,getgrav/grav,joomla/joomla-cms,octobercms/october,pagekit/pagekit,redkite-labs/RedKiteCms,roadiz/roadiz,sulu/sulu-standard,\
spip/SPIP,neos/neos-development-collection,WordPress/WordPress,modxcms/revolution,novius-os/novius-os,\
LavaLite/cms,picocms/Pico,daylightstudio/FUEL-CMS,thelia/thelia,typicms/base,AsgardCms/Platform,odirleiborgert/borgert-cms,redaxscript/redaxscript,getkirby/starterkit,processwire/processwire,\
symfony-cmf/symfony-cmf,zikula/core,TYPO3/TYPO3.CMS,drupal/drupal,\
keystonejs/keystone,Dynalon/mdwiki,directus/directus,strapi/strapi,netlify/netlify-cms,apostrophecms/apostrophe\
 -f stats/all_cms.csv
echo "[Detailed cross-language CMS statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/all_cms.csv)"
echo ""
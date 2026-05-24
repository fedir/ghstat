echo "## Cross-language CMS rating"
echo ""
./ghstat -r \
WordPress/WordPress,drupal/drupal,joomla/joomla-cms,getgrav/grav,craftcms/cms,statamic/cms,octobercms/october,\
TYPO3/TYPO3.CMS,concrete5/concrete5,neos/neos-development-collection,processwire/processwire,\
contao/core,modxcms/revolution,getkirby/starterkit,picocms/Pico,forkcms/forkcms,zikula/core,sulu/sulu-standard,\
keystonejs/keystone,directus/directus,strapi/strapi,netlify/netlify-cms,apostrophecms/apostrophe,\
dotCMS/core,alkacon/opencms-core,gentics/mesh,\
nuxeo/nuxeo,lutece-platform/lutece-core,exoplatform/ecms \
-f stats/all_cms.csv
echo "[Detailed cross-language CMS statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/all_cms.csv)"
echo ""

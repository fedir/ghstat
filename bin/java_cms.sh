#!/bin/bash
echo "## Java CMS statistics rating"
echo ""
./ghstat -r dotCMS/core,alkacon/opencms-core,gentics/mesh,Softmotions/ncms,liferay/liferay-portal,\
bogeblad/infoglue,nuxeo/nuxeo,lutece-platform/lutece-core,alkacon/opencms-core,exoplatform/ecms -f stats/java_cms.csv
echo "[Detailed Java CMS statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/java_cms.csv)"
echo ""
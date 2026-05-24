#!/bin/bash
echo "## Java CMS statistics rating"
echo ""
./ghstat -r dotCMS/core,alkacon/opencms-core,gentics/mesh,nuxeo/nuxeo,lutece-platform/lutece-core,exoplatform/ecms,bogeblad/infoglue -f stats/java_cms.csv
echo "[Detailed Java CMS statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/java_cms.csv)"

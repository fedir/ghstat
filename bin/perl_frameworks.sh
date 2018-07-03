#!/bin/bash
echo "## Perl frameworks statistics rating"
echo ""
./ghstat -r kraih/mojo,perl-catalyst/catalyst-runtime,tokuhirom/Amon,PerlDancer/Dancer,PerlDancer/Dancer2 -f stats/perl_frameworks.csv
echo "[Detailed Perl frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/perl_frameworks.csv)"
echo ""

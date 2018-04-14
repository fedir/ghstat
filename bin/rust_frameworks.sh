#!/bin/bash
echo "### Rust frameworks statistics rating"
echo ""
./ghstat -r SergioBenitez/rocket,iron/iron,actix/actix-web,gotham-rs/gotham,nickel-org/nickel.rs,Ogeon/rustful,rustless/rustless,tomaka/rouille,sappworks/sapper,mehcode/shio-rs -f stats/rust_frameworks.csv
echo "[Detailed Rust frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/rust_frameworks.csv)"
echo ""
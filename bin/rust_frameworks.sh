#!/bin/bash
echo "## Rust frameworks statistics rating"
echo ""
./ghstat -r actix/actix-web,tokio-rs/axum,rwf2/Rocket,seanmonstar/warp,leptos-rs/leptos,iron/iron,nickel-org/nickel.rs,tomaka/rouille,gotham-rs/gotham -f stats/rust_frameworks.csv
echo "[Detailed Rust frameworks statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/rust_frameworks.csv)"
echo ""
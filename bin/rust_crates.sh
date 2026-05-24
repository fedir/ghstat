#!/bin/bash
echo "## Rust crates statistics rating"
echo ""
./ghstat -r \
tokio-rs/tokio,\
serde-rs/serde,\
clap-rs/clap,\
tokio-rs/tracing,\
hyperium/hyper,\
seanmonstar/reqwest,\
dtolnay/anyhow,\
dtolnay/thiserror,\
dtolnay/syn,\
rust-lang/regex,\
rayon-rs/rayon,\
crossbeam-rs/crossbeam,\
actix/actix-web,\
tokio-rs/axum,\
launchbadge/sqlx,\
diesel-rs/diesel,\
SeaQL/sea-orm,\
rusqlite/rusqlite,\
rustls/rustls,\
rustwasm/wasm-bindgen,\
tauri-apps/tauri,\
bevyengine/bevy,\
ratatui-org/ratatui,\
pest-parser/pest,\
async-rs/async-std \
-f stats/rust_crates.csv
echo "[Detailed Rust crates statistics with ratings](https://github.com/fedir/ghstat/blob/master/stats/rust_crates.csv)"
echo ""

.PHONY: build test vet lint clean cache-clear clone-clear rate-limit run-go run-go-microservices run-all run-cncf run-rust-crates run-devops help clean-data-cms clean-data-databases clean-data-langs clean-data-go clean-data-rust clean-data-js clean-data-python clean-data-ruby clean-data-java clean-data-cncf clean-data-all

BINARY := ghstat
STATS_DIR := stats
CACHE_DIR := tmp
DATA_DIR := test_data

PKGS := $(shell go list ./... 2>/dev/null | grep -v '/tmp/' | grep -v '/test_data/')

## build: compile the binary
build:
	go build -o $(BINARY) .

## test: run tests with race detector and coverage
test:
	@echo "" > coverage.txt
	@for d in $(PKGS); do \
		go test -race -coverprofile=profile.out -covermode=atomic $$d; \
		if [ -f profile.out ]; then cat profile.out >> coverage.txt && rm profile.out; fi \
	done

## vet: run go vet
vet:
	go vet $(PKGS)

## clean: remove binary and API cache (preserves local clones in tmp/projects/)
clean:
	rm -f $(BINARY)
	rm -rf $(CACHE_DIR)

## cache-clear: clear HTTP response cache (preserves local clones)
cache-clear: build
	./$(BINARY) -cc -t $(CACHE_DIR)

## clone-clear: remove locally cloned repositories (tmp/projects/ can be several GB)
clone-clear:
	rm -rf $(CACHE_DIR)/projects

## rate-limit: show current GitHub API rate limit status
rate-limit: build
	./$(BINARY) -l

## run-go: fetch and rank Go frameworks
run-go: build
	./$(BINARY) -f $(STATS_DIR)/go_frameworks.csv -t $(CACHE_DIR)

## run-devops: fetch and rank DevOps tools
run-devops: build
	bash bin/devops_tools.sh

## run-cncf: fetch and rank top 50 CNCF cloud native projects
run-cncf: build
	bash bin/cncf_projects.sh

## run-rust-crates: fetch and rank top 25 Rust crates
run-rust-crates: build
	bash bin/rust_crates.sh

## run-go-microservices: fetch and rank Go microservice toolkits
run-go-microservices: build
	./$(BINARY) -r koding/kite,nytimes/gizmo,micro/go-micro,rsms/gotalk,gocircuit/circuit,go-kit/kit \
		-f $(STATS_DIR)/go_microservice_toolkits.csv -t $(CACHE_DIR)

## run-all: run all framework/CMS comparisons and regenerate ratings.md
run-all: build
	bash bin/build_all.sh > ratings.md

## clean-data-cms: remove locally cloned CMS project repos
clean-data-cms:
	rm -rf $(DATA_DIR)/projects/WordPress_WordPress $(DATA_DIR)/projects/drupal_drupal \
	       $(DATA_DIR)/projects/joomla_joomla-cms $(DATA_DIR)/projects/getgrav_grav \
	       $(DATA_DIR)/projects/craftcms_cms $(DATA_DIR)/projects/statamic_cms \
	       $(DATA_DIR)/projects/octobercms_october $(DATA_DIR)/projects/TYPO3_TYPO3.CMS \
	       $(DATA_DIR)/projects/concrete5_concrete5 $(DATA_DIR)/projects/neos_neos-development-collection \
	       $(DATA_DIR)/projects/processwire_processwire $(DATA_DIR)/projects/contao_core \
	       $(DATA_DIR)/projects/modxcms_revolution $(DATA_DIR)/projects/getkirby_starterkit \
	       $(DATA_DIR)/projects/picocms_Pico $(DATA_DIR)/projects/forkcms_forkcms \
	       $(DATA_DIR)/projects/zikula_core $(DATA_DIR)/projects/sulu_sulu-standard \
	       $(DATA_DIR)/projects/keystonejs_keystone $(DATA_DIR)/projects/directus_directus \
	       $(DATA_DIR)/projects/strapi_strapi $(DATA_DIR)/projects/apostrophecms_apostrophe \
	       $(DATA_DIR)/projects/dotCMS_core $(DATA_DIR)/projects/alkacon_opencms-core \
	       $(DATA_DIR)/projects/gentics_mesh $(DATA_DIR)/projects/nuxeo_nuxeo \
	       $(DATA_DIR)/projects/lutece-platform_lutece-core $(DATA_DIR)/projects/exoplatform_ecms \
	       $(DATA_DIR)/projects/bogeblad_infoglue $(DATA_DIR)/projects/netlify_netlify-cms \
	       $(DATA_DIR)/projects/hexojs_hexo

## clean-data-databases: remove locally cloned database repos
clean-data-databases:
	rm -rf $(DATA_DIR)/projects/mysql_mysql-server $(DATA_DIR)/projects/postgres_postgres \
	       $(DATA_DIR)/projects/mongodb_mongo $(DATA_DIR)/projects/MariaDB_server \
	       $(DATA_DIR)/projects/redis_redis $(DATA_DIR)/projects/dgraph-io_dgraph \
	       $(DATA_DIR)/projects/apache_hbase $(DATA_DIR)/projects/apache_cassandra \
	       $(DATA_DIR)/projects/basho_riak $(DATA_DIR)/projects/pouchdb_pouchdb \
	       $(DATA_DIR)/projects/apache_couchdb

## clean-data-langs: remove locally cloned language repos
clean-data-langs:
	rm -rf $(DATA_DIR)/projects/rust-lang_rust $(DATA_DIR)/projects/golang_go \
	       $(DATA_DIR)/projects/python_cpython $(DATA_DIR)/projects/php_php-src \
	       $(DATA_DIR)/projects/ruby_ruby $(DATA_DIR)/projects/lua_lua \
	       $(DATA_DIR)/projects/crystal-lang_crystal $(DATA_DIR)/projects/elixir-lang_elixir \
	       $(DATA_DIR)/projects/erlang_otp $(DATA_DIR)/projects/haskell_ghc

## clean-data-go: remove locally cloned Go framework repos
clean-data-go:
	rm -rf $(DATA_DIR)/projects/gin-gonic_gin $(DATA_DIR)/projects/gofiber_fiber \
	       $(DATA_DIR)/projects/labstack_echo $(DATA_DIR)/projects/go-chi_chi \
	       $(DATA_DIR)/projects/beego_beego $(DATA_DIR)/projects/gohugoio_hugo \
	       $(DATA_DIR)/projects/gobuffalo_buffalo $(DATA_DIR)/projects/revel_revel \
	       $(DATA_DIR)/projects/kataras_iris $(DATA_DIR)/projects/go-macaron_macaron \
	       $(DATA_DIR)/projects/koding_kite $(DATA_DIR)/projects/nytimes_gizmo \
	       $(DATA_DIR)/projects/micro_go-micro $(DATA_DIR)/projects/rsms_gotalk \
	       $(DATA_DIR)/projects/gocircuit_circuit $(DATA_DIR)/projects/go-kit_kit

## clean-data-rust: remove locally cloned Rust repos
clean-data-rust:
	rm -rf $(DATA_DIR)/projects/actix_actix-web $(DATA_DIR)/projects/tokio-rs_axum \
	       $(DATA_DIR)/projects/rwf2_Rocket $(DATA_DIR)/projects/seanmonstar_warp \
	       $(DATA_DIR)/projects/leptos-rs_leptos $(DATA_DIR)/projects/iron_iron \
	       $(DATA_DIR)/projects/nickel-org_nickel.rs $(DATA_DIR)/projects/tomaka_rouille \
	       $(DATA_DIR)/projects/gotham-rs_gotham $(DATA_DIR)/projects/tokio-rs_tokio \
	       $(DATA_DIR)/projects/serde-rs_serde $(DATA_DIR)/projects/clap-rs_clap \
	       $(DATA_DIR)/projects/tokio-rs_tracing $(DATA_DIR)/projects/hyperium_hyper \
	       $(DATA_DIR)/projects/seanmonstar_reqwest $(DATA_DIR)/projects/dtolnay_anyhow \
	       $(DATA_DIR)/projects/dtolnay_thiserror $(DATA_DIR)/projects/dtolnay_syn \
	       $(DATA_DIR)/projects/rust-lang_regex $(DATA_DIR)/projects/rayon-rs_rayon \
	       $(DATA_DIR)/projects/crossbeam-rs_crossbeam $(DATA_DIR)/projects/launchbadge_sqlx \
	       $(DATA_DIR)/projects/diesel-rs_diesel $(DATA_DIR)/projects/SeaQL_sea-orm \
	       $(DATA_DIR)/projects/rusqlite_rusqlite $(DATA_DIR)/projects/rustls_rustls \
	       $(DATA_DIR)/projects/rustwasm_wasm-bindgen $(DATA_DIR)/projects/tauri-apps_tauri \
	       $(DATA_DIR)/projects/bevyengine_bevy $(DATA_DIR)/projects/ratatui_ratatui \
	       $(DATA_DIR)/projects/pest-parser_pest $(DATA_DIR)/projects/async-rs_async-std

## clean-data-js: remove locally cloned JS framework repos
clean-data-js:
	rm -rf $(DATA_DIR)/projects/facebook_react $(DATA_DIR)/projects/vuejs_vue \
	       $(DATA_DIR)/projects/angular_angular $(DATA_DIR)/projects/sveltejs_svelte \
	       $(DATA_DIR)/projects/solidjs_solid $(DATA_DIR)/projects/expressjs_express \
	       $(DATA_DIR)/projects/fastify_fastify $(DATA_DIR)/projects/nestjs_nest \
	       $(DATA_DIR)/projects/koajs_koa $(DATA_DIR)/projects/honojs_hono \
	       $(DATA_DIR)/projects/vercel_next.js $(DATA_DIR)/projects/nuxt_nuxt \
	       $(DATA_DIR)/projects/remix-run_remix $(DATA_DIR)/projects/sveltejs_kit \
	       $(DATA_DIR)/projects/analogjs_analog

## clean-data-python: remove locally cloned Python framework repos
clean-data-python:
	rm -rf $(DATA_DIR)/projects/django_django $(DATA_DIR)/projects/pallets_flask \
	       $(DATA_DIR)/projects/encode_django-rest-framework $(DATA_DIR)/projects/tornadoweb_tornado \
	       $(DATA_DIR)/projects/channelcat_sanic $(DATA_DIR)/projects/bottlepy_bottle \
	       $(DATA_DIR)/projects/Pylons_pyramid $(DATA_DIR)/projects/plotly_dash \
	       $(DATA_DIR)/projects/TurboGears_tg2 $(DATA_DIR)/projects/web2py_web2py

## clean-data-ruby: remove locally cloned Ruby framework repos
clean-data-ruby:
	rm -rf $(DATA_DIR)/projects/rails_rails $(DATA_DIR)/projects/sinatra_sinatra \
	       $(DATA_DIR)/projects/hanami_hanami $(DATA_DIR)/projects/padrino_padrino-framework \
	       $(DATA_DIR)/projects/jeremyevans_roda $(DATA_DIR)/projects/camping_camping \
	       $(DATA_DIR)/projects/rack-app_rack-app $(DATA_DIR)/projects/pakyow_pakyow \
	       $(DATA_DIR)/projects/ramaze_ramaze $(DATA_DIR)/projects/voltrb_volt

## clean-data-java: remove locally cloned Java/JVM framework repos
clean-data-java:
	rm -rf $(DATA_DIR)/projects/spring-projects_spring-framework $(DATA_DIR)/projects/eclipse_vert.x \
	       $(DATA_DIR)/projects/playframework_playframework $(DATA_DIR)/projects/ninjaframework_ninja \
	       $(DATA_DIR)/projects/jooby-project_jooby $(DATA_DIR)/projects/ratpack_ratpack \
	       $(DATA_DIR)/projects/netty_netty $(DATA_DIR)/projects/ktorio_ktor \
	       $(DATA_DIR)/projects/http4k_http4k $(DATA_DIR)/projects/grails_grails-core \
	       $(DATA_DIR)/projects/scalatra_scalatra $(DATA_DIR)/projects/lift_framework

## clean-data-cncf: remove locally cloned CNCF project repos
clean-data-cncf:
	rm -rf $(DATA_DIR)/projects/kubernetes_kubernetes $(DATA_DIR)/projects/argoproj_argo-cd \
	       $(DATA_DIR)/projects/fluxcd_flux2 $(DATA_DIR)/projects/helm_helm \
	       $(DATA_DIR)/projects/prometheus_prometheus $(DATA_DIR)/projects/grafana_grafana \
	       $(DATA_DIR)/projects/envoyproxy_envoy $(DATA_DIR)/projects/istio_istio \
	       $(DATA_DIR)/projects/containerd_containerd $(DATA_DIR)/projects/open-telemetry_opentelemetry-collector \
	       $(DATA_DIR)/projects/jaegertracing_jaeger $(DATA_DIR)/projects/fluent_fluentd \
	       $(DATA_DIR)/projects/coredns_coredns $(DATA_DIR)/projects/etcd-io_etcd \
	       $(DATA_DIR)/projects/thanos-io_thanos $(DATA_DIR)/projects/cortexproject_cortex \
	       $(DATA_DIR)/projects/cri-o_cri-o $(DATA_DIR)/projects/goharbor_harbor \
	       $(DATA_DIR)/projects/linkerd_linkerd2 $(DATA_DIR)/projects/open-policy-agent_opa \
	       $(DATA_DIR)/projects/rook_rook $(DATA_DIR)/projects/cert-manager_cert-manager \
	       $(DATA_DIR)/projects/crossplane_crossplane $(DATA_DIR)/projects/cilium_cilium \
	       $(DATA_DIR)/projects/falcosecurity_falco $(DATA_DIR)/projects/kedacore_keda \
	       $(DATA_DIR)/projects/dapr_dapr $(DATA_DIR)/projects/knative_serving \
	       $(DATA_DIR)/projects/kubeedge_kubeedge $(DATA_DIR)/projects/kyverno_kyverno \
	       $(DATA_DIR)/projects/spiffe_spire $(DATA_DIR)/projects/vitessio_vitess \
	       $(DATA_DIR)/projects/tikv_tikv $(DATA_DIR)/projects/backstage_backstage \
	       $(DATA_DIR)/projects/nats-io_nats-server $(DATA_DIR)/projects/grafana_loki \
	       $(DATA_DIR)/projects/longhorn_longhorn $(DATA_DIR)/projects/strimzi_strimzi-kafka-operator \
	       $(DATA_DIR)/projects/chaos-mesh_chaos-mesh $(DATA_DIR)/projects/kubevirt_kubevirt \
	       $(DATA_DIR)/projects/operator-framework_operator-sdk $(DATA_DIR)/projects/openfeature_spec \
	       $(DATA_DIR)/projects/kubeflow_kubeflow $(DATA_DIR)/projects/containers_podman \
	       $(DATA_DIR)/projects/tektoncd_pipeline $(DATA_DIR)/projects/buildpacks_pack \
	       $(DATA_DIR)/projects/emissary-ingress_emissary $(DATA_DIR)/projects/artifact-hub_hub \
	       $(DATA_DIR)/projects/openkruise_kruise $(DATA_DIR)/projects/wasmCloud_wasmCloud

## clean-data-all: remove all locally cloned repos from test_data and tmp
clean-data-all:
	rm -rf test_data/projects $(CACHE_DIR)/projects

## help: show this help
help:
	@grep -E '^## ' Makefile | sed 's/^## //' | column -t -s ':'

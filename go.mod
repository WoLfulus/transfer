module github.com/wolfulus/transfer

go 1.13

replace (
	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.2
	github.com/containerd/containerd v1.3.0-0.20190507210959-7c1e88399ec0 => github.com/containerd/containerd v1.3.0-beta.2.0.20190823190603-4a2f61c4f2b4
	github.com/docker/distribution v2.7.1+incompatible => github.com/docker/distribution v2.7.1-0.20190205005809-0d3efadf0154+incompatible
	github.com/docker/distribution/digest => github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/docker/docker => github.com/docker/engine v1.4.2-0.20191011211953-adfac697dc5b
	github.com/jaguilar/vt100 => github.com/tonistiigi/vt100 v0.0.0-20190402012908-ad4c4a574305
	github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
	golang.org/x/crypto v0.0.0-20190129210102-0709b304e793 => golang.org/x/crypto v0.0.0-20180904163835-0709b304e793
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/Microsoft/hcsshim v0.8.6 // indirect
	github.com/Shopify/logrus-bugsnag v0.0.0-20171204204709-577dee27f20d // indirect
	github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412 // indirect
	github.com/bitly/go-hostpool v0.0.0-20171023180738-a3a6125de932 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/bugsnag/bugsnag-go v1.5.3 // indirect
	github.com/bugsnag/panicwrap v1.2.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cloudflare/cfssl v0.0.0-20190911221928-1a911ca1b1d6 // indirect
	github.com/containerd/cgroups v0.0.0-20191011165608-5fbad35c2a7e // indirect
	github.com/containerd/containerd v1.3.0 // indirect
	github.com/containerd/fifo v0.0.0-20190816180239-bda0ff6ed73c // indirect
	github.com/containerd/ttrpc v0.0.0-20191024154809-5046735d25c3 // indirect
	github.com/containerd/typeurl v0.0.0-20190911142611-5eb25027c9fd // indirect
	github.com/docker/cli v0.0.0-20191024152931-69b73f7519a4
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v1.14.0-0.20190319215453-e7b5f7dbe98c
	github.com/docker/docker-credential-helpers v0.6.3 // indirect
	github.com/docker/go v1.5.1-1 // indirect
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/libtrust v0.0.0-20160708172513-aabc10ec26b7 // indirect
	github.com/eknkc/basex v1.0.0
	github.com/foomo/htpasswd v0.0.0-20180422071726-cb63c4ac0e50
	github.com/gin-gonic/gin v1.7.0
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/gogo/googleapis v1.3.0 // indirect
	github.com/google/certificate-transparency-go v1.0.21 // indirect
	github.com/gorilla/mux v1.7.3 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/jinzhu/gorm v1.9.11 // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/mattn/go-shellwords v1.0.6 // indirect
	github.com/miekg/pkcs11 v1.0.3 // indirect
	github.com/moby/buildkit v0.6.2 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/runtime-spec v1.0.1 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/theupdateframework/notary v0.6.1 // indirect
	github.com/tonistiigi/fsutil v0.0.0-20191018213012-0f039a052ca1 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	github.com/xlab/handysort v0.0.0-20150421192137-fb3537ed64a1 // indirect
	github.com/zmap/zcrypto v0.0.0-20191023172918-f1f7c0342eee // indirect
	github.com/zmap/zlint v1.0.2 // indirect
	go.etcd.io/bbolt v1.3.3 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	google.golang.org/grpc v1.24.0 // indirect
	gopkg.in/dancannon/gorethink.v3 v3.0.5 // indirect
	gopkg.in/fatih/pool.v2 v2.0.0 // indirect
	gopkg.in/gorethink/gorethink.v3 v3.0.5 // indirect
	vbom.ml/util v0.0.0-20180919145318-efcd4e0f9787 // indirect
)

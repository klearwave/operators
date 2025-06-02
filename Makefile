VERSION ?= unstable
OS ?= darwin
ARCH ?= amd64

#
# developer tools
#
VENDIR_VERSION ?= v0.30.0
vendir:
        @curl -L https://github.com/vmware-tanzu/carvel-vendir/releases/download/$(VENDIR_VERSION)/vendir-$(OS)-$(ARCH) -o /usr/local/bin/vendir && chmod +x /usr/local/bin/vendir

YOT_VERSION ?= v0.6.5
yot:
        @curl -L https://github.com/vmware-tanzu-labs/yaml-overlay-tool/releases/download/$(YOT_VERSION)/yot_$(YOT_VERSION)_Darwin_x86_64.tar.gz -o /usr/local/bin/yot.tar.gz && \
                tar zxvf /usr/local/bin/yot.tar.gz -C /usr/local/bin && \
                chmod +x /usr/local/bin/yot

#
# code generation
#
CAPABILITY ?=
COMPONENT ?=

# create component scaffold
scaffold:
	@scripts/scaffold.sh

# download from upstream
download:
	@scripts/download.sh

download-all:
	for capability in `find capabilities -type d -maxdepth 1 -mindepth 1 -exec basename {} \;`; do \
		for component in `find capabilities/$$capability/.source -type d -maxdepth 1 -mindepth 1 -exec basename {} \;`; do \
			CAPABILITY=$$capability COMPONENT=$$component scripts/download.sh; \
		done; \
	done

# overlay changes from upstream
overlay:
	@scripts/overlay.sh

overlay-all:
	for capability in `find capabilities -type d -maxdepth 1 -mindepth 1 -exec basename {} \;`; do \
		for component in `find capabilities/$$capability/.source -type d -maxdepth 1 -mindepth 1 -exec basename {} \;`; do \
			CAPABILITY=$$capability COMPONENT=$$component scripts/overlay.sh; \
		done; \
	done

operator-init:
	@VERSION=$(VERSION) scripts/operator-init.sh

operator-create:
	@VERSION=$(VERSION) scripts/operator-create.sh

# preserve manually modified assets
preserve:
	@scripts/preserve-assets.sh

restore:
	@scripts/restore-assets.sh

#
# testing helpers
#
manifests-apply:
	@kubectl create ns klearwave-$(CAPABILITY)-system --dry-run=client -o yaml | kubectl apply -f -
	@kubectl apply -f capabilities/$(CAPABILITY)/.source/$(COMPONENT)/manifests/

manifests-delete:
	@kubectl delete -f capabilities/$(CAPABILITY)/.source/$(COMPONENT)/manifests/

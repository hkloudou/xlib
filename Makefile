.PHONY: default
default:
	-cd $(shell git rev-parse --show-toplevel) && git autotag -commit 'auto commit'
	@echo current version:`git describe`
pub:
	-cd $(shell git rev-parse --show-toplevel) && git autotag -commit 'auto commit' -t -i -f
	@echo current version:`git describe`
all: bitvector.a bfilter.a

bitvector.a:
	(cd src/container/bitvector; $(GOBIN)/gomake install)

bfilter.a:
	(cd src/container/bfilter; $(GOBIN)/gomake install)

compiled:
	@echo "[Compiling] everything ..."
	go build

mal.bin: compiled
	@echo "[Copying] executable as mal.bin ..."
	cp lisg $@

SHELL := bash
mal: compiled
	@echo "[Running] the tests ..."
	cat <(echo -e '#!/bin/sh\nexec "$$0" "$$@"') mal.bin > $@
	chmod +x mal

clean:
	go clean
	rm -f mal.bin mal

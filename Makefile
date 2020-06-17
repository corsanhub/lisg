step0_repl:
	go build
	cp lisg $@

step1_read_print:
	go build
	cp lisg $@

step2_eval:
	go build
	cp lisg $@

mal.bin: step0_repl step1_read_print
	@echo "[Copying] executable as mal.bin ..."
	cp lisg $@

SHELL := bash
mal: compiled
	@echo "[Running] the tests ..."
	cat <(echo -e '#!/bin/sh\nexec "$$0" "$$@"') mal.bin > $@
	chmod +x mal

clean:
	go clean
	#rm -f mal.bin mal
	rm -f mal.bin mal
	rm -f lisg step0_repl step1_read_print step2_eval

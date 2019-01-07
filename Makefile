GOPATH	= $(CURDIR)
BINDIR	= $(CURDIR)/bin

PROGRAMS = check_vault_seal

build:
	env GOPATH=$(GOPATH) go install $(PROGRAMS)

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/bin

strip: build
	strip --strip-all $(BINDIR)/check_vault_seal

install: strip destdirs install-bin

install-bin:
	install -m 0755 $(BINDIR)/check_vault_seal $(DESTDIR)/usr/bin

clean:
	/bin/rm -f bin/check_vault_seal

uninstall:
	/bin/rm -f $(DESTDIR)/usr/bin

all: build strip install


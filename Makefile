include $(GOROOT)/src/Make.$(GOARCH)

TARG=google/language/translate
GOFILES=\
	google/language/translate.go\

include $(GOROOT)/src/Make.pkg

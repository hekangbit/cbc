all:
	$(MAKE) -C parser
	go build .

clean:
	$(MAKE) -C parser clean

all:
	$(MAKE) -C parser
	go build .

clean:
	rm -rf cbc
	$(MAKE) -C parser clean

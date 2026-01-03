all:
	$(MAKE) -C grammar
	go build .

clean:
	rm -rf cbc
	rm -rf parser

PROJECT_NAME := cbc
SRC_DIR := src
GRAMMAR_DIR := grammar
PARSER_DIR := parser
TARGET_BIN := cbc
MAIN_SRCS := $(wildcard $(SRC_DIR)/*.go)
PARSER_SRCS := \
	$(PARSER_DIR)/Cb.interp \
	$(PARSER_DIR)/Cb.tokens \
	$(PARSER_DIR)/CbLexer.interp \
	$(PARSER_DIR)/CbLexer.tokens \
	$(PARSER_DIR)/cb_base_listener.go \
	$(PARSER_DIR)/cb_base_visitor.go \
	$(PARSER_DIR)/cb_lexer.go \
	$(PARSER_DIR)/cb_parser.go \
	$(PARSER_DIR)/cb_visitor.go

.PHONY: all clean

all: ${TARGET_BIN}

${PARSER_SRCS}: $(GRAMMAR_DIR)/Makefile $(GRAMMAR_DIR)/*.g4
	@echo "generate parser..."
	$(MAKE) -C $(GRAMMAR_DIR)
	@echo "generate parser done"

${TARGET_BIN}: ${PARSER_SRCS} ${MAIN_SRCS}
	@echo "build $(TARGET_BIN)..."
	go build -o $(TARGET_BIN) ./$(SRC_DIR)
	@echo "build done"

clean:
	rm -f $(TARGET_BIN)
	rm -rf $(PARSER_DIR)
GRAMMAR_FILE := Cb.g4
OUTPUT_DIR := ../parser
PACKAGE_NAME := parser
ANTLR_JAR := /usr/local/lib/antlr-4.13.2-complete.jar
ANTLR_CMD = java -Xmx500M -cp "${ANTLR_JAR}:$CLASSPATH" org.antlr.v4.Tool

all: $(GRAMMAR_FILE)
		${ANTLR_CMD} -Dlanguage=Go -no-listener -visitor -package ${PACKAGE_NAME} -o ${OUTPUT_DIR} $<

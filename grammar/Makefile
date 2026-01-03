GrammarFile := Cb.g4

all: $(GrammarFile)
		java -Xmx500M -cp "/usr/local/lib/antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go -no-visitor -package parser -o ../parser $<

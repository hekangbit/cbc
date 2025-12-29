
GrammarFile := Cb.g4
ParserFiles := *.go *.interp *.tokens

$(ParserFiles): $(GrammarFile)
		java -Xmx500M -cp "/usr/local/lib/antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool -Dlanguage=Go -no-visitor -package parser $<

clean:
		rm -f $(ParserFiles)
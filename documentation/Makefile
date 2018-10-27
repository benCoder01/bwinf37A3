Dokumentation.pdf: template/eisvogel.tex Dokumentation.md clean
	pandoc Dokumentation.md --template template/eisvogel.tex -o Dokumentation.pdf
clean: 
	$(RM) Dokumentation.pdf

open: 
	google-chrome Dokumentation.pdf
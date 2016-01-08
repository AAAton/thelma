#Thelma
This is a project in the NLP field by Anton Södergren and Dalibor Lovric. The code takes a novel (or general piece of text in Swedish), by Selma Lagerlöf and finds and extracts the characters within that novel. 

It uses the POS tagger called Stagger, developed and maintained by Robert Östling. 

##Installing
In order to make the file, you need to have Go installed. Once downloaded navigate to the folder and write `go build`. This produces a binary file called `thelma`.

You also need to download a model file for the part of speech tagger. We have chosen to use the model file for Swedish, by the Stockholm Umeå Corpus, supplied on [Staggers website](http://www.ling.su.se/english/nlp/tools/stagger/stagger-the-stockholm-tagger-1.98986).

##Running the code
This should not be very difficult:
`./thelma -f [path/to/textfile/containing/the/novel]`
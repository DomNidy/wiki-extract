# Wiki-Extract

Simple CLI program written in Golang. Extracts the text content from a list of wikipedia urls. Useful for datascience/ML purposes.

By default, the config file & output directories are located in the home directory of the current user.

## What does it do?

After providing URL(s) to parse text from, wiki-extract will request the pages and output 3 different text files.

### These files are:

- The raw HTML of the page `Article_name_raw.txt`
- Parsed "contentful" text (text which is not metadata/html stuff) `Article_name_parsed.txt`
- List of related wikipedia urls the current wikipedia article links to `Article_name_related.txt`

## How to use it

_Will document this soon :)_

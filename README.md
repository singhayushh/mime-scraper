<div align="center">
  <h1>Mime Scraper</h1>
  <h3>Scrapes mimetype metadata and exports as json</h3>
</div>

## Contents

  -   [Contents](#contents)
  -   [Description](#description)
      -   [What's the problem we are trying to solve?](#whats-the-problem-we-are-trying-to-solve)
      -   [How can ABC help?](#how-can-abc-help)
      -   [The idea](#the-idea)
  -   [Project structure](#project-structure)
  -   [Getting started](#getting-started)
      -   [Prerequisites](#prerequisites)
          -   [Softwares needed](#softwares-needed)
          -   [Knowledge needed](#knowledge-needed)
      -   [Installing](#installing)
  -   [Future Scope](#future-scope)
  -   [Built with](#built-with)
  -   [Authors](#authors)

## Description

This project is about a web scraping script that scrapes data from [Complete MIME Types List](https://www.freeformatter.com/mime-types-list.html), and exports them to a json file.

### What's the problem we are trying to solve?

For one of my internal projects, where users could upload any file and store in a custom cloud db, it was required for the files to be identified via their mimetypes along with some basic info about the file types along with a suitable icon for each file type.

### How can Mime Scraper help?

Mime Scraper creates a json of all existing mimetypes supported by HTTP, with the following 5 fields:
- name: name of the file type
- mimetype: mimetype of the file
- extension: extension of the file associated
- icon: this field is to be populated with the relevant icon theme you'd be using, so it has been kept empty. (I've used tabler icons, you may request the final file for taber icons integration)
- info: url for more information about the mimetype, mostly points to some wikipedia article.

### The idea

The idea was to write a super fast script that can scrape and parse data quickly in an exportable format. I've used Go, for its high speed, and the colly framework - the fastest available atm. The data fetched by scraping is parsed to a custom struct type and pushed to a list of the same data type. Once scraping is completed, the list is written to `mimetypes.json` file using the `os` package.

## Project structure

I've used the standard golang project structing. The code resides in the following directory on my Unix system:
`$GOPATH/src/github.com/singhayushh/mime-scraper`

## Getting started

### Prerequisites

#### Softwares needed

- Go

#### Knowledge needed

- HTML and XPath
- Mimetypes

### Installing

- create a directory under `$GOPATH/src/github.com` and name it singhayushh
- clone the repo inside this directory
- cd to the cloned directory and run `go get` to install modules from `go.mod` file
- run `go run main.go` or create a executable build.

## Built with

- Go
- Colly

## Authors

<a href="https://github.com/singhayushh/mime-scraper/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=singhayushh/mime-scraper" />
</a>

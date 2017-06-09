[![Build Status](https://travis-ci.org/xwb1989/sqlparser.svg?branch=master)](https://travis-ci.org/xwb1989/sqlparser)

##Notice

The backbone of this repo is extracted from [youtube/vitess](https://github.com/youtube/vitess).

Inside youtube/vitess there is a very nicely written sql parser. However as it's not a self-contained application, I created this one. 
It applies the same LICENSE as youtube/vitess.

##What's More In this REPO

* Rewrite SQL Queries by self-defined rewriter
* Able to parse Create Table statements way much better

##Usage

    import (
        "github.com/xwb1989/sqlparser"
    )

Then use
    
    sqlparser.Parse(sql)

See `parse_test.go` for more `Parse` usage.

See `rewriter_test.go` for `Rewriter` usage.

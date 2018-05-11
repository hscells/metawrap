package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/hscells/metawrap"
	"github.com/alexflint/go-arg"
	"bytes"
)

var mm metawrap.MetaMap

type args struct {
	Path string `arg:"help:path to MetaMap binary"`
}

func (args) Version() string {
	return "metawrap 01.May.2018"
}

func (args) Description() string {
	return `wrapper for MetaMap`
}

func handleMap(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	s := bytes.NewBuffer(b).String()
	log.Println(s)
	candidates, err := mm.Map(s)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, candidates)
	return
}

func handleCandidates(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	s := bytes.NewBuffer(b).String()
	candidates, err := mm.PreferredCandidates(s)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, candidates)
	return
}

func main() {
	// Parse the args into the struct
	var args args
	arg.MustParse(&args)

	mm = metawrap.NewMetaMapClient(args.Path)

	router := gin.Default()

	// Main query interface.
	router.GET("/mm/map", handleMap)
	router.GET("/mm/candidates", handleCandidates)

	log.Println("let's go!")
	log.Fatal(http.ListenAndServe(":4646", router))
}

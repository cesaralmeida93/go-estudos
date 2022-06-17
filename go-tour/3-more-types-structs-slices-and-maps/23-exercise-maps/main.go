package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

       var stringArray []string = strings.Fields(s)
       wordMap := make(map[string]int)
       for _, word := range stringArray {
              count, ok := wordMap[word]
              if ok {
                     wordMap[word] = count + 1
              } else {
                     wordMap[word] = 1
              }
       }
       return wordMap
}

func main() {
       wc.Test(WordCount)
}
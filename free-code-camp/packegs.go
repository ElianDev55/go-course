package main

import (
	"endcoding/csv"
	"os"
	"strconv"
)

type Alum struct { // struct or interface in typeScript
  Name string
  Score []float64
}

func readCvg(name string) ([]Alum, error) {
  file, err := os.Open(name)

  if err!= nil {
    return nil, err
  }
  
  reader := csv.NewReader(file)
  reader.Comma = ';'

  records, err := reader.ReadAll()
  if err!= nil {
    return nil, err
  }

  var alumnos []Alum

for _, records := range records {
    name := records[0]
    var scores []float64
    for _, socoreStr := range records[1:] {
        socore, err := strconv.ParseFloat(socoreStr, 64)
        if err != nil {
            return nil, err
        }
        scores = append(scores, socore)
    }
    alumnos = append(alumnos, Alum{Name: name, Score: scores})
}



}

func main() {

}

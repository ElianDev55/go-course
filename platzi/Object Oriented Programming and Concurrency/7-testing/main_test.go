package main

import "testing"

func TestSum(t *testing.T){
	
	/*
	total := Sum(5,5)

	if total != 10 {
		t.Error("Expected 10, got ", total)
	}
		*/

		tables := []struct {
			a int
			b int 
			n int 
		}{
			{5, 5, 10},
      {10, 15, 25},
      {-5, 5, 0},
		}

		for _, table := range tables {
      total := Sum(table.a, table.b)
      if total != table.n {
        t.Errorf("Expected %d, got %d", table.n, total)
      }
    }
}

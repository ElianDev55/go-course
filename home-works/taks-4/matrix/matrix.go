package matrix

import "fmt"



type Matrix struct {
	w  int64
	h  int64
	m  [][]int64
	q  bool
}

func New(arrays ...[]int64) (*Matrix,error) {
	m := Matrix{
		w: int64(len(arrays)),
	}

	for key, value := range arrays {
		if int64(len(value)) != m.w {
				return nil,fmt.Errorf("row %d length is not equal to row %d",key+1,m.w)
		}
	}
	m.h = int64(len(arrays[0]))
	
	for _ , v := range arrays{
		m.m = append(m.m,v)
	}

	m.q = m.h == m.w

	return &m,nil
	
}

func (m *Matrix) Print() {
	for _,v := range m.m{
		fmt.Println("Array:",v)
	}
	fmt.Println()
	fmt.Println("Rows:",m.h,"Columns:",m.w)
	fmt.Println("Matrix is square:",m.q)
}


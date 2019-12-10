package main
import(
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestGetInput(t*testing.T){
	tests := []struct {
		startIndex       int
		program []int
		expected []int

		
	}{
		{0, []int{1,2,3,4,5}, []int{1,2,3,4}},
		{4, []int{1,2,3,4,5,6,7,8, 9}, []int{5,6,7,8}  },
	
	}
	for _, test := range tests {
		result := getInput(test.startIndex, test.program)
	
	assert.Equal(t, test.expected, result)
	}
}

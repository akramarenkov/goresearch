package args

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type fields struct {
	first  int
	second int
}

func (flds fields) Modify() (int, int) {
	flds.first++

	flds.modifySecond()

	return flds.first, flds.second
}

func (flds *fields) modifySecond() {
	flds.second += 2
}

func TestModifyReceiver(t *testing.T) {
	receiver := fields{}

	first, second := receiver.Modify()
	require.Equal(t, fields{}, receiver)
	require.Equal(t, 1, first)
	require.Equal(t, 2, second)
}

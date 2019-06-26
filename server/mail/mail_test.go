package mail

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSend(t *testing.T) {
	err := Send("imuge@qq.com", "证书", "imuge")
	require.Nil(t, err)
}

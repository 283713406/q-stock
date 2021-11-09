package cmds

import (
	"testing"

	"github.com/283713406/q-stock/models"
	"github.com/stretchr/testify/require"
)

func TestExportPic(t *testing.T) {
	e := Exportor{
		Stocks: []models.ExportorData{
			{
				Name: "中文名称",
				Code: "1234code",
			}, {
				Name: "中文名称1",
				Code: "code12345",
			},
		},
	}

	PicChuckSize = 1
	_, err := e.ExportPic(_ctx, "/tmp/test.png")
	require.Nil(t, err)
}

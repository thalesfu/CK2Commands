package livs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珀拉瓦尔PooravereBarony struct {
	feud.BaseBarony
}

var BPooravere珀拉瓦尔 feud.Barony = &珀拉瓦尔PooravereBarony{}

func init() {
	f := BPooravere珀拉瓦尔.(*珀拉瓦尔PooravereBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pooravere",
		TitleName: "珀拉瓦尔",
		TitleCode: "b_pooravere",
	}
}

package zeila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 路哈亚LughayaBarony struct {
	feud.BaseBarony
}

var BLughaya路哈亚 feud.Barony = &路哈亚LughayaBarony{}

func init() {
	f := BLughaya路哈亚.(*路哈亚LughayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lughaya",
		TitleName: "路哈亚",
		TitleCode: "b_lughaya",
	}
}

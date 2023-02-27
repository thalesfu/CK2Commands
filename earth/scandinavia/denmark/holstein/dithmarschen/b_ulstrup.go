package dithmarschen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔斯楚普UlstrupBarony struct {
	feud.BaseBarony
}

var BUlstrup乌尔斯楚普 feud.Barony = &乌尔斯楚普UlstrupBarony{}

func init() {
    f := BUlstrup乌尔斯楚普.(*乌尔斯楚普UlstrupBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulstrup",
		TitleName: "乌尔斯楚普",
		TitleCode: "b_ulstrup",
	}
}

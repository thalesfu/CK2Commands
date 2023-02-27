package medapata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑斯婆呼Sas_bahuBarony struct {
	feud.BaseBarony
}

var BSas_bahu娑斯婆呼 feud.Barony = &娑斯婆呼Sas_bahuBarony{}

func init() {
    f := BSas_bahu娑斯婆呼.(*娑斯婆呼Sas_bahuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sas_bahu",
		TitleName: "娑斯婆呼",
		TitleCode: "b_sas_bahu",
	}
}

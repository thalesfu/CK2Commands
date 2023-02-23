package penthievre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢代阿克LoudeacBarony struct {
	feud.BaseBarony
}

var BLoudeac卢代阿克 feud.Barony = &卢代阿克LoudeacBarony{}

func init() {
	f := BLoudeac卢代阿克.(*卢代阿克LoudeacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "loudeac",
		TitleName: "卢代阿克",
		TitleCode: "b_loudeac",
	}
}

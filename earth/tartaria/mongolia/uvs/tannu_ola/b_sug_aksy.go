package tannu_ola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏格阿克瑟Sug_aksyBarony struct {
	feud.BaseBarony
}

var BSug_aksy苏格阿克瑟 feud.Barony = &苏格阿克瑟Sug_aksyBarony{}

func init() {
    f := BSug_aksy苏格阿克瑟.(*苏格阿克瑟Sug_aksyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sug_aksy",
		TitleName: "苏格阿克瑟",
		TitleCode: "b_sug_aksy",
	}
}

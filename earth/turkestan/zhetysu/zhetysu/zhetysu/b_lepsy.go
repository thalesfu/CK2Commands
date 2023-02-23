package zhetysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列普瑟LepsyBarony struct {
	feud.BaseBarony
}

var BLepsy列普瑟 feud.Barony = &列普瑟LepsyBarony{}

func init() {
	f := BLepsy列普瑟.(*列普瑟LepsyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lepsy",
		TitleName: "列普瑟",
		TitleCode: "b_lepsy",
	}
}

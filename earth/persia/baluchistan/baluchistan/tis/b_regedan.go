package tis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷德丹RegedanBarony struct {
	feud.BaseBarony
}

var BRegedan雷德丹 feud.Barony = &雷德丹RegedanBarony{}

func init() {
	f := BRegedan雷德丹.(*雷德丹RegedanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "regedan",
		TitleName: "雷德丹",
		TitleCode: "b_regedan",
	}
}

package agder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗莱克菲尤尔FlekkefjordBarony struct {
	feud.BaseBarony
}

var BFlekkefjord弗莱克菲尤尔 feud.Barony = &弗莱克菲尤尔FlekkefjordBarony{}

func init() {
	f := BFlekkefjord弗莱克菲尤尔.(*弗莱克菲尤尔FlekkefjordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "flekkefjord",
		TitleName: "弗莱克菲尤尔",
		TitleCode: "b_flekkefjord",
	}
}

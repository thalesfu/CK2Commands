package asirgarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑那跋陀SanawadBarony struct {
	feud.BaseBarony
}

var BSanawad娑那跋陀 feud.Barony = &娑那跋陀SanawadBarony{}

func init() {
    f := BSanawad娑那跋陀.(*娑那跋陀SanawadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanawad",
		TitleName: "娑那跋陀",
		TitleCode: "b_sanawad",
	}
}

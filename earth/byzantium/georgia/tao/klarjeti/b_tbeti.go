package klarjeti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特贝蒂TbetiBarony struct {
	feud.BaseBarony
}

var BTbeti特贝蒂 feud.Barony = &特贝蒂TbetiBarony{}

func init() {
    f := BTbeti特贝蒂.(*特贝蒂TbetiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tbeti",
		TitleName: "特贝蒂",
		TitleCode: "b_tbeti",
	}
}

package kairwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰卢拉JalulaBarony struct {
	feud.BaseBarony
}

var BJalula杰卢拉 feud.Barony = &杰卢拉JalulaBarony{}

func init() {
    f := BJalula杰卢拉.(*杰卢拉JalulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jalula",
		TitleName: "杰卢拉",
		TitleCode: "b_jalula",
	}
}

package derbent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拒马JumaBarony struct {
	feud.BaseBarony
}

var BJuma拒马 feud.Barony = &拒马JumaBarony{}

func init() {
    f := BJuma拒马.(*拒马JumaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "juma",
		TitleName: "拒马",
		TitleCode: "b_juma",
	}
}

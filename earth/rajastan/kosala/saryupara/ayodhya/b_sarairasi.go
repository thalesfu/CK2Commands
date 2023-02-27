package ayodhya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨罗依罗斯SarairasiBarony struct {
	feud.BaseBarony
}

var BSarairasi萨罗依罗斯 feud.Barony = &萨罗依罗斯SarairasiBarony{}

func init() {
    f := BSarairasi萨罗依罗斯.(*萨罗依罗斯SarairasiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarairasi",
		TitleName: "萨罗依罗斯",
		TitleCode: "b_sarairasi",
	}
}

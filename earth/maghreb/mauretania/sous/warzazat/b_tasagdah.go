package warzazat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰塞格达TasagdahBarony struct {
	feud.BaseBarony
}

var BTasagdah泰塞格达 feud.Barony = &泰塞格达TasagdahBarony{}

func init() {
    f := BTasagdah泰塞格达.(*泰塞格达TasagdahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tasagdah",
		TitleName: "泰塞格达",
		TitleCode: "b_tasagdah",
	}
}

package belgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基希讷乌ChisinauBarony struct {
	feud.BaseBarony
}

var BChisinau基希讷乌 feud.Barony = &基希讷乌ChisinauBarony{}

func init() {
    f := BChisinau基希讷乌.(*基希讷乌ChisinauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chisinau",
		TitleName: "基希讷乌",
		TitleCode: "b_chisinau",
	}
}

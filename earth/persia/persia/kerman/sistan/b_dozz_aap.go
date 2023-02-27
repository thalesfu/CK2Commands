package sistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达兹阿拉Dozz_aapBarony struct {
	feud.BaseBarony
}

var BDozz_aap达兹阿拉 feud.Barony = &达兹阿拉Dozz_aapBarony{}

func init() {
    f := BDozz_aap达兹阿拉.(*达兹阿拉Dozz_aapBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dozz_aap",
		TitleName: "达兹阿拉",
		TitleCode: "b_dozz_aap",
	}
}

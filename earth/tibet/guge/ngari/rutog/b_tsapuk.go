package rutog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎普TsapukBarony struct {
	feud.BaseBarony
}

var BTsapuk扎普 feud.Barony = &扎普TsapukBarony{}

func init() {
    f := BTsapuk扎普.(*扎普TsapukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsapuk",
		TitleName: "扎普",
		TitleCode: "b_tsapuk",
	}
}

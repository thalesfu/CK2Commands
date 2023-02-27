package khokh_serkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 察斯特TsastBarony struct {
	feud.BaseBarony
}

var BTsast察斯特 feud.Barony = &察斯特TsastBarony{}

func init() {
    f := BTsast察斯特.(*察斯特TsastBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsast",
		TitleName: "察斯特",
		TitleCode: "b_tsast",
	}
}

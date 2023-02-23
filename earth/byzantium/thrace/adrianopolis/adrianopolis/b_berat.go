package adrianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 培拉特BeratBarony struct {
	feud.BaseBarony
}

var BBerat培拉特 feud.Barony = &培拉特BeratBarony{}

func init() {
	f := BBerat培拉特.(*培拉特BeratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "berat",
		TitleName: "培拉特",
		TitleCode: "b_berat",
	}
}

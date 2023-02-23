package tadjrift

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌伊费特OuifatBarony struct {
	feud.BaseBarony
}

var BOuifat乌伊费特 feud.Barony = &乌伊费特OuifatBarony{}

func init() {
	f := BOuifat乌伊费特.(*乌伊费特OuifatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouifat",
		TitleName: "乌伊费特",
		TitleCode: "b_ouifat",
	}
}

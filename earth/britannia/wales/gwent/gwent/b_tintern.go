package gwent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 延特恩TinternBarony struct {
	feud.BaseBarony
}

var BTintern延特恩 feud.Barony = &延特恩TinternBarony{}

func init() {
	f := BTintern延特恩.(*延特恩TinternBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tintern",
		TitleName: "延特恩",
		TitleCode: "b_tintern",
	}
}

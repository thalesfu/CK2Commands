package burkhan_buudai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朝格特TsogtBarony struct {
	feud.BaseBarony
}

var BTsogt朝格特 feud.Barony = &朝格特TsogtBarony{}

func init() {
    f := BTsogt朝格特.(*朝格特TsogtBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tsogt",
		TitleName: "朝格特",
		TitleCode: "b_tsogt",
	}
}

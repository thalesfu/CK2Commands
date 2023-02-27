package derby

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切斯特菲尔德ChesterfieldBarony struct {
	feud.BaseBarony
}

var BChesterfield切斯特菲尔德 feud.Barony = &切斯特菲尔德ChesterfieldBarony{}

func init() {
    f := BChesterfield切斯特菲尔德.(*切斯特菲尔德ChesterfieldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chesterfield",
		TitleName: "切斯特菲尔德",
		TitleCode: "b_chesterfield",
	}
}

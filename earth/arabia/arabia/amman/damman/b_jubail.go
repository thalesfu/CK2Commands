package damman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朱拜勒JubailBarony struct {
	feud.BaseBarony
}

var BJubail朱拜勒 feud.Barony = &朱拜勒JubailBarony{}

func init() {
	f := BJubail朱拜勒.(*朱拜勒JubailBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jubail",
		TitleName: "朱拜勒",
		TitleCode: "b_jubail",
	}
}

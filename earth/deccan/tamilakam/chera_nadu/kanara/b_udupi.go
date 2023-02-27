package kanara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤杜毗UdupiBarony struct {
	feud.BaseBarony
}

var BUdupi尤杜毗 feud.Barony = &尤杜毗UdupiBarony{}

func init() {
    f := BUdupi尤杜毗.(*尤杜毗UdupiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udupi",
		TitleName: "尤杜毗",
		TitleCode: "b_udupi",
	}
}

package yaik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 列比亚日耶LebyazheBarony struct {
	feud.BaseBarony
}

var BLebyazhe列比亚日耶 feud.Barony = &列比亚日耶LebyazheBarony{}

func init() {
    f := BLebyazhe列比亚日耶.(*列比亚日耶LebyazheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lebyazhe",
		TitleName: "列比亚日耶",
		TitleCode: "b_lebyazhe",
	}
}

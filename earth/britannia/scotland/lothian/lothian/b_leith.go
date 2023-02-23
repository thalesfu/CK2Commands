package lothian

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利斯LeithBarony struct {
	feud.BaseBarony
}

var BLeith利斯 feud.Barony = &利斯LeithBarony{}

func init() {
	f := BLeith利斯.(*利斯LeithBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leith",
		TitleName: "利斯",
		TitleCode: "b_leith",
	}
}

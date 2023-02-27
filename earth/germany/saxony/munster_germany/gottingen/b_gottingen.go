package gottingen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哥廷根GottingenBarony struct {
	feud.BaseBarony
}

var BGottingen哥廷根 feud.Barony = &哥廷根GottingenBarony{}

func init() {
    f := BGottingen哥廷根.(*哥廷根GottingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gottingen",
		TitleName: "哥廷根",
		TitleCode: "b_gottingen",
	}
}

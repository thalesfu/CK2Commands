package deltuva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德尔图瓦DeltuvaBarony struct {
	feud.BaseBarony
}

var BDeltuva德尔图瓦 feud.Barony = &德尔图瓦DeltuvaBarony{}

func init() {
    f := BDeltuva德尔图瓦.(*德尔图瓦DeltuvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "deltuva",
		TitleName: "德尔图瓦",
		TitleCode: "b_deltuva",
	}
}

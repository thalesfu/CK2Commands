package wurttemberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格明德GmundBarony struct {
	feud.BaseBarony
}

var BGmund格明德 feud.Barony = &格明德GmundBarony{}

func init() {
    f := BGmund格明德.(*格明德GmundBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gmund",
		TitleName: "格明德",
		TitleCode: "b_gmund",
	}
}

package kartli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德马尼西DmanisiBarony struct {
	feud.BaseBarony
}

var BDmanisi德马尼西 feud.Barony = &德马尼西DmanisiBarony{}

func init() {
	f := BDmanisi德马尼西.(*德马尼西DmanisiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dmanisi",
		TitleName: "德马尼西",
		TitleCode: "b_dmanisi",
	}
}

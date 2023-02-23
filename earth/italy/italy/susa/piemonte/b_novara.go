package piemonte

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺瓦拉NovaraBarony struct {
	feud.BaseBarony
}

var BNovara诺瓦拉 feud.Barony = &诺瓦拉NovaraBarony{}

func init() {
	f := BNovara诺瓦拉.(*诺瓦拉NovaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novara",
		TitleName: "诺瓦拉",
		TitleCode: "b_novara",
	}
}

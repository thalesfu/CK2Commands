package sibi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 娑薜荔陀SaprendaBarony struct {
	feud.BaseBarony
}

var BSaprenda娑薜荔陀 feud.Barony = &娑薜荔陀SaprendaBarony{}

func init() {
	f := BSaprenda娑薜荔陀.(*娑薜荔陀SaprendaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saprenda",
		TitleName: "娑薜荔陀",
		TitleCode: "b_saprenda",
	}
}

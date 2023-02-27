package makran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈乌尔Al_haurBarony struct {
	feud.BaseBarony
}

var BAl_haur哈乌尔 feud.Barony = &哈乌尔Al_haurBarony{}

func init() {
    f := BAl_haur哈乌尔.(*哈乌尔Al_haurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "al_haur",
		TitleName: "哈乌尔",
		TitleCode: "b_al_haur",
	}
}

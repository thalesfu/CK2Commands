package ascalon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡提尔HuidreBarony struct {
	feud.BaseBarony
}

var BHuidre胡提尔 feud.Barony = &胡提尔HuidreBarony{}

func init() {
	f := BHuidre胡提尔.(*胡提尔HuidreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huidre",
		TitleName: "胡提尔",
		TitleCode: "b_huidre",
	}
}

package bidar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆多摩罗BhatambraBarony struct {
	feud.BaseBarony
}

var BBhatambra婆多摩罗 feud.Barony = &婆多摩罗BhatambraBarony{}

func init() {
    f := BBhatambra婆多摩罗.(*婆多摩罗BhatambraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhatambra",
		TitleName: "婆多摩罗",
		TitleCode: "b_bhatambra",
	}
}

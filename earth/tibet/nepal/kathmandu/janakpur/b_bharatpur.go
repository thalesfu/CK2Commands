package janakpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆罗多补罗BharatpurBarony struct {
	feud.BaseBarony
}

var BBharatpur婆罗多补罗 feud.Barony = &婆罗多补罗BharatpurBarony{}

func init() {
	f := BBharatpur婆罗多补罗.(*婆罗多补罗BharatpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bharatpur",
		TitleName: "婆罗多补罗",
		TitleCode: "b_bharatpur",
	}
}

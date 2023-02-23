package york

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米德尔斯堡MiddlesbroughBarony struct {
	feud.BaseBarony
}

var BMiddlesbrough米德尔斯堡 feud.Barony = &米德尔斯堡MiddlesbroughBarony{}

func init() {
	f := BMiddlesbrough米德尔斯堡.(*米德尔斯堡MiddlesbroughBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "middlesbrough",
		TitleName: "米德尔斯堡",
		TitleCode: "b_middlesbrough",
	}
}

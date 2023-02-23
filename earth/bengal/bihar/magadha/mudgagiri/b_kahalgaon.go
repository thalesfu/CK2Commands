package mudgagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦诃罗伽罗摩KahalgaonBarony struct {
	feud.BaseBarony
}

var BKahalgaon迦诃罗伽罗摩 feud.Barony = &迦诃罗伽罗摩KahalgaonBarony{}

func init() {
	f := BKahalgaon迦诃罗伽罗摩.(*迦诃罗伽罗摩KahalgaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kahalgaon",
		TitleName: "迦诃罗伽罗摩",
		TitleCode: "b_kahalgaon",
	}
}

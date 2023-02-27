package kartli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特莫格维TmoghviBarony struct {
	feud.BaseBarony
}

var BTmoghvi特莫格维 feud.Barony = &特莫格维TmoghviBarony{}

func init() {
    f := BTmoghvi特莫格维.(*特莫格维TmoghviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tmoghvi",
		TitleName: "特莫格维",
		TitleCode: "b_tmoghvi",
	}
}

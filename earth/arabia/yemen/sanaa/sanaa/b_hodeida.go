package sanaa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 荷台达HodeidaBarony struct {
	feud.BaseBarony
}

var BHodeida荷台达 feud.Barony = &荷台达HodeidaBarony{}

func init() {
    f := BHodeida荷台达.(*荷台达HodeidaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hodeida",
		TitleName: "荷台达",
		TitleCode: "b_hodeida",
	}
}

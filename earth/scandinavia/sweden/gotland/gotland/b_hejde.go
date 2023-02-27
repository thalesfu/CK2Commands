package gotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑德HejdeBarony struct {
	feud.BaseBarony
}

var BHejde黑德 feud.Barony = &黑德HejdeBarony{}

func init() {
    f := BHejde黑德.(*黑德HejdeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hejde",
		TitleName: "黑德",
		TitleCode: "b_hejde",
	}
}

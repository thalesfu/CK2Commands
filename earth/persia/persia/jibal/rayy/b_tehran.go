package rayy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德黑兰TehranBarony struct {
	feud.BaseBarony
}

var BTehran德黑兰 feud.Barony = &德黑兰TehranBarony{}

func init() {
    f := BTehran德黑兰.(*德黑兰TehranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tehran",
		TitleName: "德黑兰",
		TitleCode: "b_tehran",
	}
}

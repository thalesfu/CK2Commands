package lingtsang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 色须SerxuBarony struct {
	feud.BaseBarony
}

var BSerxu色须 feud.Barony = &色须SerxuBarony{}

func init() {
	f := BSerxu色须.(*色须SerxuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "serxu",
		TitleName: "色须",
		TitleCode: "b_serxu",
	}
}

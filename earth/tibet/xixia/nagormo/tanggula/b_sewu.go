package tanggula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 色务SewuBarony struct {
	feud.BaseBarony
}

var BSewu色务 feud.Barony = &色务SewuBarony{}

func init() {
	f := BSewu色务.(*色务SewuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sewu",
		TitleName: "色务",
		TitleCode: "b_sewu",
	}
}

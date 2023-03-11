package trans-portage

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼扬多马NyandomaBarony struct {
	feud.BaseBarony
}

var BNyandoma尼扬多马 feud.Barony = &尼扬多马NyandomaBarony{}

func init() {
    f := BNyandoma尼扬多马.(*尼扬多马NyandomaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nyandoma",
		TitleName: "尼扬多马",
		TitleCode: "b_nyandoma",
	}
}

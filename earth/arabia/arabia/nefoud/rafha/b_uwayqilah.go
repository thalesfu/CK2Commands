package rafha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧韦吉莱UwayqilahBarony struct {
	feud.BaseBarony
}

var BUwayqilah欧韦吉莱 feud.Barony = &欧韦吉莱UwayqilahBarony{}

func init() {
    f := BUwayqilah欧韦吉莱.(*欧韦吉莱UwayqilahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uwayqilah",
		TitleName: "欧韦吉莱",
		TitleCode: "b_uwayqilah",
	}
}

package mainz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 魏尔堡WeilburgBarony struct {
	feud.BaseBarony
}

var BWeilburg魏尔堡 feud.Barony = &魏尔堡WeilburgBarony{}

func init() {
	f := BWeilburg魏尔堡.(*魏尔堡WeilburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "weilburg",
		TitleName: "魏尔堡",
		TitleCode: "b_weilburg",
	}
}

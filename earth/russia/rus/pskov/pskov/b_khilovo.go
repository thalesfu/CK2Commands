package pskov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希洛沃KhilovoBarony struct {
	feud.BaseBarony
}

var BKhilovo希洛沃 feud.Barony = &希洛沃KhilovoBarony{}

func init() {
    f := BKhilovo希洛沃.(*希洛沃KhilovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khilovo",
		TitleName: "希洛沃",
		TitleCode: "b_khilovo",
	}
}

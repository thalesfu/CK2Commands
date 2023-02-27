package veliky_ustug

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢扎LuzaBarony struct {
	feud.BaseBarony
}

var BLuza卢扎 feud.Barony = &卢扎LuzaBarony{}

func init() {
    f := BLuza卢扎.(*卢扎LuzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luza",
		TitleName: "卢扎",
		TitleCode: "b_luza",
	}
}

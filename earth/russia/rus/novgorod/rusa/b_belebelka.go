package rusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别列别尔卡BelebelkaBarony struct {
	feud.BaseBarony
}

var BBelebelka别列别尔卡 feud.Barony = &别列别尔卡BelebelkaBarony{}

func init() {
    f := BBelebelka别列别尔卡.(*别列别尔卡BelebelkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belebelka",
		TitleName: "别列别尔卡",
		TitleCode: "b_belebelka",
	}
}

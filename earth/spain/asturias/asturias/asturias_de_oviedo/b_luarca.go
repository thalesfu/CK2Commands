package asturias_de_oviedo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢阿尔卡LuarcaBarony struct {
	feud.BaseBarony
}

var BLuarca卢阿尔卡 feud.Barony = &卢阿尔卡LuarcaBarony{}

func init() {
    f := BLuarca卢阿尔卡.(*卢阿尔卡LuarcaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "luarca",
		TitleName: "卢阿尔卡",
		TitleCode: "b_luarca",
	}
}

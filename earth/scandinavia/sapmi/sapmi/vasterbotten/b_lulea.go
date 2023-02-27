package vasterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吕勒奥LuleaBarony struct {
	feud.BaseBarony
}

var BLulea吕勒奥 feud.Barony = &吕勒奥LuleaBarony{}

func init() {
    f := BLulea吕勒奥.(*吕勒奥LuleaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lulea",
		TitleName: "吕勒奥",
		TitleCode: "b_lulea",
	}
}

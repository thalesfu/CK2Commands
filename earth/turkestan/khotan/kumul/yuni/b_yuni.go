package yuni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扜泥YuniBarony struct {
	feud.BaseBarony
}

var BYuni扜泥 feud.Barony = &扜泥YuniBarony{}

func init() {
    f := BYuni扜泥.(*扜泥YuniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yuni",
		TitleName: "扜泥",
		TitleCode: "b_yuni",
	}
}

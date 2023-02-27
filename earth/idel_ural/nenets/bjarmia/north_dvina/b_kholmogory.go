package north_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尔莫戈雷KholmogoryBarony struct {
	feud.BaseBarony
}

var BKholmogory霍尔莫戈雷 feud.Barony = &霍尔莫戈雷KholmogoryBarony{}

func init() {
    f := BKholmogory霍尔莫戈雷.(*霍尔莫戈雷KholmogoryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kholmogory",
		TitleName: "霍尔莫戈雷",
		TitleCode: "b_kholmogory",
	}
}

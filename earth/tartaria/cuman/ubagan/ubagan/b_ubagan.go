package ubagan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌巴甘UbaganBarony struct {
	feud.BaseBarony
}

var BUbagan乌巴甘 feud.Barony = &乌巴甘UbaganBarony{}

func init() {
    f := BUbagan乌巴甘.(*乌巴甘UbaganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ubagan",
		TitleName: "乌巴甘",
		TitleCode: "b_ubagan",
	}
}

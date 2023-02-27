package farafra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 白色沙漠White_desertBarony struct {
	feud.BaseBarony
}

var BWhite_desert白色沙漠 feud.Barony = &白色沙漠White_desertBarony{}

func init() {
    f := BWhite_desert白色沙漠.(*白色沙漠White_desertBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "white_desert",
		TitleName: "白色沙漠",
		TitleCode: "b_white_desert",
	}
}

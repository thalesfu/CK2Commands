package zhongba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 隆格尔LunggarBarony struct {
	feud.BaseBarony
}

var BLunggar隆格尔 feud.Barony = &隆格尔LunggarBarony{}

func init() {
    f := BLunggar隆格尔.(*隆格尔LunggarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lunggar",
		TitleName: "隆格尔",
		TitleCode: "b_lunggar",
	}
}

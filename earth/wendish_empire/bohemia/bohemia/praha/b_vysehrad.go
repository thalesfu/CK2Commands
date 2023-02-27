package praha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 高堡VysehradBarony struct {
	feud.BaseBarony
}

var BVysehrad高堡 feud.Barony = &高堡VysehradBarony{}

func init() {
    f := BVysehrad高堡.(*高堡VysehradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vysehrad",
		TitleName: "高堡",
		TitleCode: "b_vysehrad",
	}
}

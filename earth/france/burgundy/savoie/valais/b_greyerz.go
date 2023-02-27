package valais

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格雷耶茨GreyerzBarony struct {
	feud.BaseBarony
}

var BGreyerz格雷耶茨 feud.Barony = &格雷耶茨GreyerzBarony{}

func init() {
    f := BGreyerz格雷耶茨.(*格雷耶茨GreyerzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "greyerz",
		TitleName: "格雷耶茨",
		TitleCode: "b_greyerz",
	}
}

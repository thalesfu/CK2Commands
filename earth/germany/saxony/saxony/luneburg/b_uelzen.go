package luneburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于尔岑UelzenBarony struct {
	feud.BaseBarony
}

var BUelzen于尔岑 feud.Barony = &于尔岑UelzenBarony{}

func init() {
	f := BUelzen于尔岑.(*于尔岑UelzenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uelzen",
		TitleName: "于尔岑",
		TitleCode: "b_uelzen",
	}
}

package chilia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基利亚ChiliaBarony struct {
	feud.BaseBarony
}

var BChilia基利亚 feud.Barony = &基利亚ChiliaBarony{}

func init() {
    f := BChilia基利亚.(*基利亚ChiliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chilia",
		TitleName: "基利亚",
		TitleCode: "b_chilia",
	}
}

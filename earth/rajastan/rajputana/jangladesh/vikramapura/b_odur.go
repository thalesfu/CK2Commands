package vikramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌杜尔OdurBarony struct {
	feud.BaseBarony
}

var BOdur乌杜尔 feud.Barony = &乌杜尔OdurBarony{}

func init() {
	f := BOdur乌杜尔.(*乌杜尔OdurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "odur",
		TitleName: "乌杜尔",
		TitleCode: "b_odur",
	}
}

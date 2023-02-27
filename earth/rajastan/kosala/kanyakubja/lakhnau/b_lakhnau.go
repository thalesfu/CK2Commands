package lakhnau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 落伽奈LakhnauBarony struct {
	feud.BaseBarony
}

var BLakhnau落伽奈 feud.Barony = &落伽奈LakhnauBarony{}

func init() {
    f := BLakhnau落伽奈.(*落伽奈LakhnauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lakhnau",
		TitleName: "落伽奈",
		TitleCode: "b_lakhnau",
	}
}

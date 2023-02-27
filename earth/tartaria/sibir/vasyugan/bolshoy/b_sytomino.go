package bolshoy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑟托米诺SytominoBarony struct {
	feud.BaseBarony
}

var BSytomino瑟托米诺 feud.Barony = &瑟托米诺SytominoBarony{}

func init() {
    f := BSytomino瑟托米诺.(*瑟托米诺SytominoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sytomino",
		TitleName: "瑟托米诺",
		TitleCode: "b_sytomino",
	}
}

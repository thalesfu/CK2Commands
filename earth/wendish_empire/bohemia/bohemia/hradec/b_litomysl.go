package hradec

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利托米什尔LitomyslBarony struct {
	feud.BaseBarony
}

var BLitomysl利托米什尔 feud.Barony = &利托米什尔LitomyslBarony{}

func init() {
    f := BLitomysl利托米什尔.(*利托米什尔LitomyslBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "litomysl",
		TitleName: "利托米什尔",
		TitleCode: "b_litomysl",
	}
}

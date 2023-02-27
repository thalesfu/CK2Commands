package kuban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔皮利KirpiliBarony struct {
	feud.BaseBarony
}

var BKirpili基尔皮利 feud.Barony = &基尔皮利KirpiliBarony{}

func init() {
    f := BKirpili基尔皮利.(*基尔皮利KirpiliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirpili",
		TitleName: "基尔皮利",
		TitleCode: "b_kirpili",
	}
}

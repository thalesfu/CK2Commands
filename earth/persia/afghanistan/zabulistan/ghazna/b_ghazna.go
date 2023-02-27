package ghazna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽色尼GhaznaBarony struct {
	feud.BaseBarony
}

var BGhazna伽色尼 feud.Barony = &伽色尼GhaznaBarony{}

func init() {
    f := BGhazna伽色尼.(*伽色尼GhaznaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghazna",
		TitleName: "伽色尼",
		TitleCode: "b_ghazna",
	}
}

package senoussi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾卢JaluBarony struct {
	feud.BaseBarony
}

var BJalu贾卢 feud.Barony = &贾卢JaluBarony{}

func init() {
    f := BJalu贾卢.(*贾卢JaluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jalu",
		TitleName: "贾卢",
		TitleCode: "b_jalu",
	}
}

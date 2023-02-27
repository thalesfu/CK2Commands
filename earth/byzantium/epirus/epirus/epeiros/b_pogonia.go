package epeiros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波戈尼亚PogoniaBarony struct {
	feud.BaseBarony
}

var BPogonia波戈尼亚 feud.Barony = &波戈尼亚PogoniaBarony{}

func init() {
    f := BPogonia波戈尼亚.(*波戈尼亚PogoniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pogonia",
		TitleName: "波戈尼亚",
		TitleCode: "b_pogonia",
	}
}

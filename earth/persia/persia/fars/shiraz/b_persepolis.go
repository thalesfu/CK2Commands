package shiraz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波斯波利斯PersepolisBarony struct {
	feud.BaseBarony
}

var BPersepolis波斯波利斯 feud.Barony = &波斯波利斯PersepolisBarony{}

func init() {
    f := BPersepolis波斯波利斯.(*波斯波利斯PersepolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "persepolis",
		TitleName: "波斯波利斯",
		TitleCode: "b_persepolis",
	}
}

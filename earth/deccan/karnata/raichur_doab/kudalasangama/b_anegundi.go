package kudalasangama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿内古尼提AnegundiBarony struct {
	feud.BaseBarony
}

var BAnegundi阿内古尼提 feud.Barony = &阿内古尼提AnegundiBarony{}

func init() {
    f := BAnegundi阿内古尼提.(*阿内古尼提AnegundiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anegundi",
		TitleName: "阿内古尼提",
		TitleCode: "b_anegundi",
	}
}

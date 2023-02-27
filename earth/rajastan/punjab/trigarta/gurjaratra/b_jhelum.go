package gurjaratra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰赫勒姆JhelumBarony struct {
	feud.BaseBarony
}

var BJhelum杰赫勒姆 feud.Barony = &杰赫勒姆JhelumBarony{}

func init() {
    f := BJhelum杰赫勒姆.(*杰赫勒姆JhelumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jhelum",
		TitleName: "杰赫勒姆",
		TitleCode: "b_jhelum",
	}
}

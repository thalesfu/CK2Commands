package samtho

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雪环ZholhunBarony struct {
	feud.BaseBarony
}

var BZholhun雪环 feud.Barony = &雪环ZholhunBarony{}

func init() {
    f := BZholhun雪环.(*雪环ZholhunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zholhun",
		TitleName: "雪环",
		TitleCode: "b_zholhun",
	}
}

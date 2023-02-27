package tarvagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 昆古尔图格KungurtugBarony struct {
	feud.BaseBarony
}

var BKungurtug昆古尔图格 feud.Barony = &昆古尔图格KungurtugBarony{}

func init() {
    f := BKungurtug昆古尔图格.(*昆古尔图格KungurtugBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kungurtug",
		TitleName: "昆古尔图格",
		TitleCode: "b_kungurtug",
	}
}

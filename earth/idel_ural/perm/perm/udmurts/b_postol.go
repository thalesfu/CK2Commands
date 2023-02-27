package udmurts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波斯托尔PostolBarony struct {
	feud.BaseBarony
}

var BPostol波斯托尔 feud.Barony = &波斯托尔PostolBarony{}

func init() {
    f := BPostol波斯托尔.(*波斯托尔PostolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "postol",
		TitleName: "波斯托尔",
		TitleCode: "b_postol",
	}
}

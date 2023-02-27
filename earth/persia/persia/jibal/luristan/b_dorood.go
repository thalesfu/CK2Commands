package luristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 道鲁德DoroodBarony struct {
	feud.BaseBarony
}

var BDorood道鲁德 feud.Barony = &道鲁德DoroodBarony{}

func init() {
    f := BDorood道鲁德.(*道鲁德DoroodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dorood",
		TitleName: "道鲁德",
		TitleCode: "b_dorood",
	}
}

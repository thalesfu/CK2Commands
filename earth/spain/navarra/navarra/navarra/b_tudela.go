package navarra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图德拉TudelaBarony struct {
	feud.BaseBarony
}

var BTudela图德拉 feud.Barony = &图德拉TudelaBarony{}

func init() {
    f := BTudela图德拉.(*图德拉TudelaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tudela",
		TitleName: "图德拉",
		TitleCode: "b_tudela",
	}
}

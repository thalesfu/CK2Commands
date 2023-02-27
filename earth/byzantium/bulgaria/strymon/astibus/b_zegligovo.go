package astibus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 热格利戈沃ZegligovoBarony struct {
	feud.BaseBarony
}

var BZegligovo热格利戈沃 feud.Barony = &热格利戈沃ZegligovoBarony{}

func init() {
    f := BZegligovo热格利戈沃.(*热格利戈沃ZegligovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zegligovo",
		TitleName: "热格利戈沃",
		TitleCode: "b_zegligovo",
	}
}

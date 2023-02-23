package sennar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 辛贾SinjahBarony struct {
	feud.BaseBarony
}

var BSinjah辛贾 feud.Barony = &辛贾SinjahBarony{}

func init() {
	f := BSinjah辛贾.(*辛贾SinjahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sinjah",
		TitleName: "辛贾",
		TitleCode: "b_sinjah",
	}
}

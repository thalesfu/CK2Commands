package havelberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶里肖JerichowBarony struct {
	feud.BaseBarony
}

var BJerichow耶里肖 feud.Barony = &耶里肖JerichowBarony{}

func init() {
    f := BJerichow耶里肖.(*耶里肖JerichowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jerichow",
		TitleName: "耶里肖",
		TitleCode: "b_jerichow",
	}
}

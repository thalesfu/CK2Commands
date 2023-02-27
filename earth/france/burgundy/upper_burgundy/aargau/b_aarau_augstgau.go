package aargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿劳Aarau_augstgauBarony struct {
	feud.BaseBarony
}

var BAarau_augstgau阿劳 feud.Barony = &阿劳Aarau_augstgauBarony{}

func init() {
    f := BAarau_augstgau阿劳.(*阿劳Aarau_augstgauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aarau_augstgau",
		TitleName: "阿劳",
		TitleCode: "b_aarau_augstgau",
	}
}

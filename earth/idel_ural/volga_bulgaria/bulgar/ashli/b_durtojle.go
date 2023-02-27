package ashli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 久尔秋利DurtojleBarony struct {
	feud.BaseBarony
}

var BDurtojle久尔秋利 feud.Barony = &久尔秋利DurtojleBarony{}

func init() {
    f := BDurtojle久尔秋利.(*久尔秋利DurtojleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "durtojle",
		TitleName: "久尔秋利",
		TitleCode: "b_durtojle",
	}
}

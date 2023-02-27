package dendi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉杜JiduBarony struct {
	feud.BaseBarony
}

var BJidu吉杜 feud.Barony = &吉杜JiduBarony{}

func init() {
    f := BJidu吉杜.(*吉杜JiduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jidu",
		TitleName: "吉杜",
		TitleCode: "b_jidu",
	}
}

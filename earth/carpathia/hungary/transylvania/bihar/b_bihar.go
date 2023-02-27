package bihar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比豪尔BiharBarony struct {
	feud.BaseBarony
}

var BBihar比豪尔 feud.Barony = &比豪尔BiharBarony{}

func init() {
    f := BBihar比豪尔.(*比豪尔BiharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bihar",
		TitleName: "比豪尔",
		TitleCode: "b_bihar",
	}
}

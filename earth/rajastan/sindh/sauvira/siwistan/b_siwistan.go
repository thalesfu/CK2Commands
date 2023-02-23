package siwistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室毗湿檀那SiwistanBarony struct {
	feud.BaseBarony
}

var BSiwistan室毗湿檀那 feud.Barony = &室毗湿檀那SiwistanBarony{}

func init() {
	f := BSiwistan室毗湿檀那.(*室毗湿檀那SiwistanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siwistan",
		TitleName: "室毗湿檀那",
		TitleCode: "b_siwistan",
	}
}

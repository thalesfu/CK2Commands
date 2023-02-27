package gowrie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邓凯尔德DunkeldBarony struct {
	feud.BaseBarony
}

var BDunkeld邓凯尔德 feud.Barony = &邓凯尔德DunkeldBarony{}

func init() {
    f := BDunkeld邓凯尔德.(*邓凯尔德DunkeldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dunkeld",
		TitleName: "邓凯尔德",
		TitleCode: "b_dunkeld",
	}
}

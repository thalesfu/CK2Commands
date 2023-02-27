package friesach

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔特霍芬AlthofenBarony struct {
	feud.BaseBarony
}

var BAlthofen阿尔特霍芬 feud.Barony = &阿尔特霍芬AlthofenBarony{}

func init() {
    f := BAlthofen阿尔特霍芬.(*阿尔特霍芬AlthofenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "althofen",
		TitleName: "阿尔特霍芬",
		TitleCode: "b_althofen",
	}
}

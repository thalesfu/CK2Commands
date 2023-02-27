package salerno

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格罗波利AgropoliBarony struct {
	feud.BaseBarony
}

var BAgropoli阿格罗波利 feud.Barony = &阿格罗波利AgropoliBarony{}

func init() {
    f := BAgropoli阿格罗波利.(*阿格罗波利AgropoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agropoli",
		TitleName: "阿格罗波利",
		TitleCode: "b_agropoli",
	}
}

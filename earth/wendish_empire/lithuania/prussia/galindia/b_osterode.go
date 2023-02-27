package galindia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯特罗德OsterodeBarony struct {
	feud.BaseBarony
}

var BOsterode奥斯特罗德 feud.Barony = &奥斯特罗德OsterodeBarony{}

func init() {
    f := BOsterode奥斯特罗德.(*奥斯特罗德OsterodeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osterode",
		TitleName: "奥斯特罗德",
		TitleCode: "b_osterode",
	}
}

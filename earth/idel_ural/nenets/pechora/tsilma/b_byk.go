package tsilma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝克BykBarony struct {
	feud.BaseBarony
}

var BByk贝克 feud.Barony = &贝克BykBarony{}

func init() {
    f := BByk贝克.(*贝克BykBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "byk",
		TitleName: "贝克",
		TitleCode: "b_byk",
	}
}

package markakol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝科沃BykovoBarony struct {
	feud.BaseBarony
}

var BBykovo贝科沃 feud.Barony = &贝科沃BykovoBarony{}

func init() {
    f := BBykovo贝科沃.(*贝科沃BykovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bykovo",
		TitleName: "贝科沃",
		TitleCode: "b_bykovo",
	}
}

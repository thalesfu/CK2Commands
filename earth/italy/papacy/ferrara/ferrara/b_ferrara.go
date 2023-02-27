package ferrara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 费拉拉FerraraBarony struct {
	feud.BaseBarony
}

var BFerrara费拉拉 feud.Barony = &费拉拉FerraraBarony{}

func init() {
    f := BFerrara费拉拉.(*费拉拉FerraraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ferrara",
		TitleName: "费拉拉",
		TitleCode: "b_ferrara",
	}
}

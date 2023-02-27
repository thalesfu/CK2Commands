package altmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔贝克WalbeckBarony struct {
	feud.BaseBarony
}

var BWalbeck瓦尔贝克 feud.Barony = &瓦尔贝克WalbeckBarony{}

func init() {
    f := BWalbeck瓦尔贝克.(*瓦尔贝克WalbeckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "walbeck",
		TitleName: "瓦尔贝克",
		TitleCode: "b_walbeck",
	}
}

package uppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌普萨拉UppsalaBarony struct {
	feud.BaseBarony
}

var BUppsala乌普萨拉 feud.Barony = &乌普萨拉UppsalaBarony{}

func init() {
	f := BUppsala乌普萨拉.(*乌普萨拉UppsalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uppsala",
		TitleName: "乌普萨拉",
		TitleCode: "b_uppsala",
	}
}

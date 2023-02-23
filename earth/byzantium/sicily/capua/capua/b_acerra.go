package capua

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿切拉AcerraBarony struct {
	feud.BaseBarony
}

var BAcerra阿切拉 feud.Barony = &阿切拉AcerraBarony{}

func init() {
	f := BAcerra阿切拉.(*阿切拉AcerraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "acerra",
		TitleName: "阿切拉",
		TitleCode: "b_acerra",
	}
}

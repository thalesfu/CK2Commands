package zarma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉TeraBarony struct {
	feud.BaseBarony
}

var BTera特拉 feud.Barony = &特拉TeraBarony{}

func init() {
    f := BTera特拉.(*特拉TeraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tera",
		TitleName: "特拉",
		TitleCode: "b_tera",
	}
}

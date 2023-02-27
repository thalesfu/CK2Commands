package assab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉海托RehaytoBarony struct {
	feud.BaseBarony
}

var BRehayto拉海托 feud.Barony = &拉海托RehaytoBarony{}

func init() {
    f := BRehayto拉海托.(*拉海托RehaytoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rehayto",
		TitleName: "拉海托",
		TitleCode: "b_rehayto",
	}
}

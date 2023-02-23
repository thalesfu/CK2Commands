package metz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧丹勒罗芒AudunleromanBarony struct {
	feud.BaseBarony
}

var BAudunleroman欧丹勒罗芒 feud.Barony = &欧丹勒罗芒AudunleromanBarony{}

func init() {
	f := BAudunleroman欧丹勒罗芒.(*欧丹勒罗芒AudunleromanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "audunleroman",
		TitleName: "欧丹勒罗芒",
		TitleCode: "b_audunleroman",
	}
}

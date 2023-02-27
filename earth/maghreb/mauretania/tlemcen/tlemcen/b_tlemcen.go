package tlemcen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特莱姆森TlemcenBarony struct {
	feud.BaseBarony
}

var BTlemcen特莱姆森 feud.Barony = &特莱姆森TlemcenBarony{}

func init() {
    f := BTlemcen特莱姆森.(*特莱姆森TlemcenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tlemcen",
		TitleName: "特莱姆森",
		TitleCode: "b_tlemcen",
	}
}

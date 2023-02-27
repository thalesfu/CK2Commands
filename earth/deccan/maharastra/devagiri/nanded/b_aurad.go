package nanded

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥罗陀AuradBarony struct {
	feud.BaseBarony
}

var BAurad奥罗陀 feud.Barony = &奥罗陀AuradBarony{}

func init() {
    f := BAurad奥罗陀.(*奥罗陀AuradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aurad",
		TitleName: "奥罗陀",
		TitleCode: "b_aurad",
	}
}

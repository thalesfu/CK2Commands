package cholamandalam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌莱尤尔UraiyurBarony struct {
	feud.BaseBarony
}

var BUraiyur乌莱尤尔 feud.Barony = &乌莱尤尔UraiyurBarony{}

func init() {
    f := BUraiyur乌莱尤尔.(*乌莱尤尔UraiyurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uraiyur",
		TitleName: "乌莱尤尔",
		TitleCode: "b_uraiyur",
	}
}

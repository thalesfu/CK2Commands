package cakrakuta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇迦陀罗城JagdalpurBarony struct {
	feud.BaseBarony
}

var BJagdalpur阇迦陀罗城 feud.Barony = &阇迦陀罗城JagdalpurBarony{}

func init() {
	f := BJagdalpur阇迦陀罗城.(*阇迦陀罗城JagdalpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jagdalpur",
		TitleName: "阇迦陀罗城",
		TitleCode: "b_jagdalpur",
	}
}

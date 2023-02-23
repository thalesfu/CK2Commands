package reggio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特罗佩阿TropeaBarony struct {
	feud.BaseBarony
}

var BTropea特罗佩阿 feud.Barony = &特罗佩阿TropeaBarony{}

func init() {
	f := BTropea特罗佩阿.(*特罗佩阿TropeaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tropea",
		TitleName: "特罗佩阿",
		TitleCode: "b_tropea",
	}
}

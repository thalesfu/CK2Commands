package qamdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 类乌齐RiwocheBarony struct {
	feud.BaseBarony
}

var BRiwoche类乌齐 feud.Barony = &类乌齐RiwocheBarony{}

func init() {
	f := BRiwoche类乌齐.(*类乌齐RiwocheBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "riwoche",
		TitleName: "类乌齐",
		TitleCode: "b_riwoche",
	}
}

package bam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿比沙坎AbeshkanBarony struct {
	feud.BaseBarony
}

var BAbeshkan阿比沙坎 feud.Barony = &阿比沙坎AbeshkanBarony{}

func init() {
	f := BAbeshkan阿比沙坎.(*阿比沙坎AbeshkanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abeshkan",
		TitleName: "阿比沙坎",
		TitleCode: "b_abeshkan",
	}
}

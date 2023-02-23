package almaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 则克台镇ZeketaizhenBarony struct {
	feud.BaseBarony
}

var BZeketaizhen则克台镇 feud.Barony = &则克台镇ZeketaizhenBarony{}

func init() {
	f := BZeketaizhen则克台镇.(*则克台镇ZeketaizhenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zeketaizhen",
		TitleName: "则克台镇",
		TitleCode: "b_zeketaizhen",
	}
}

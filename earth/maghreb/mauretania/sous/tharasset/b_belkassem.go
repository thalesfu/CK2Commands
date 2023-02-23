package tharasset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝尔卡塞姆BelkassemBarony struct {
	feud.BaseBarony
}

var BBelkassem贝尔卡塞姆 feud.Barony = &贝尔卡塞姆BelkassemBarony{}

func init() {
	f := BBelkassem贝尔卡塞姆.(*贝尔卡塞姆BelkassemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belkassem",
		TitleName: "贝尔卡塞姆",
		TitleCode: "b_belkassem",
	}
}

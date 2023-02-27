package galatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯波尼亚AsponiaBarony struct {
	feud.BaseBarony
}

var BAsponia阿斯波尼亚 feud.Barony = &阿斯波尼亚AsponiaBarony{}

func init() {
    f := BAsponia阿斯波尼亚.(*阿斯波尼亚AsponiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asponia",
		TitleName: "阿斯波尼亚",
		TitleCode: "b_asponia",
	}
}

package quattara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿卜杜勒纳比AbdannabiBarony struct {
	feud.BaseBarony
}

var BAbdannabi阿卜杜勒纳比 feud.Barony = &阿卜杜勒纳比AbdannabiBarony{}

func init() {
	f := BAbdannabi阿卜杜勒纳比.(*阿卜杜勒纳比AbdannabiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abdannabi",
		TitleName: "阿卜杜勒纳比",
		TitleCode: "b_abdannabi",
	}
}

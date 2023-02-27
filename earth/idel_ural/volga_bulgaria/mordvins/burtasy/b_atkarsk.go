package burtasy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿特卡尔斯克AtkarskBarony struct {
	feud.BaseBarony
}

var BAtkarsk阿特卡尔斯克 feud.Barony = &阿特卡尔斯克AtkarskBarony{}

func init() {
    f := BAtkarsk阿特卡尔斯克.(*阿特卡尔斯克AtkarskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atkarsk",
		TitleName: "阿特卡尔斯克",
		TitleCode: "b_atkarsk",
	}
}

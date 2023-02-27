package kuma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔斯兰别克ArslanbekBarony struct {
	feud.BaseBarony
}

var BArslanbek阿尔斯兰别克 feud.Barony = &阿尔斯兰别克ArslanbekBarony{}

func init() {
    f := BArslanbek阿尔斯兰别克.(*阿尔斯兰别克ArslanbekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arslanbek",
		TitleName: "阿尔斯兰别克",
		TitleCode: "b_arslanbek",
	}
}

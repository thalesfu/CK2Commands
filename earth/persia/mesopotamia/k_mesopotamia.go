package mesopotamia

import (
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/jazira"
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/mosul"
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/mudar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MesopotamiaKingdom interface {
	feud.Kingdom
	DJazira贾兹拉() jazira.JaziraDuke
	DMosul摩苏尔() mosul.MosulDuke
	DMudar穆达尔() mudar.MudarDuke
}

type 贾兹拉MesopotamiaKingdom struct {
	feud.BaseKingdom
	贾兹拉Jazira jazira.JaziraDuke
	摩苏尔Mosul  mosul.MosulDuke
	穆达尔Mudar  mudar.MudarDuke
}

func (k *贾兹拉MesopotamiaKingdom) DJazira贾兹拉() jazira.JaziraDuke {
	return k.贾兹拉Jazira
}

func (k *贾兹拉MesopotamiaKingdom) DMosul摩苏尔() mosul.MosulDuke {
	return k.摩苏尔Mosul
}

func (k *贾兹拉MesopotamiaKingdom) DMudar穆达尔() mudar.MudarDuke {
	return k.穆达尔Mudar
}

var KMesopotamia贾兹拉 MesopotamiaKingdom = &贾兹拉MesopotamiaKingdom{}

func init() {
	f := KMesopotamia贾兹拉.(*贾兹拉MesopotamiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "mesopotamia",
		TitleName: "贾兹拉",
		TitleCode: "k_mesopotamia",
		Dukes:     map[string]feud.Duke{},
	}

	f.贾兹拉Jazira = jazira.DJazira贾兹拉
	f.贾兹拉Jazira.SetParent(f)

	f.摩苏尔Mosul = mosul.DMosul摩苏尔
	f.摩苏尔Mosul.SetParent(f)

	f.穆达尔Mudar = mudar.DMudar穆达尔
	f.穆达尔Mudar.SetParent(f)

}

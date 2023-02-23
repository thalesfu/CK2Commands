package malwa

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/anupa"
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/dadhipadra"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MalwaKingdom interface {
	feud.Kingdom
	DAnupa阿怒波() anupa.AnupaDuke
	DDadhipadra阿般提() dadhipadra.DadhipadraDuke
}

type 摩腊婆MalwaKingdom struct {
	feud.BaseKingdom
	阿怒波Anupa      anupa.AnupaDuke
	阿般提Dadhipadra dadhipadra.DadhipadraDuke
}

func (k *摩腊婆MalwaKingdom) DAnupa阿怒波() anupa.AnupaDuke {
	return k.阿怒波Anupa
}

func (k *摩腊婆MalwaKingdom) DDadhipadra阿般提() dadhipadra.DadhipadraDuke {
	return k.阿般提Dadhipadra
}

var KMalwa摩腊婆 MalwaKingdom = &摩腊婆MalwaKingdom{}

func init() {
	f := KMalwa摩腊婆.(*摩腊婆MalwaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "malwa",
		TitleName: "摩腊婆",
		TitleCode: "k_malwa",
		Dukes:     map[string]feud.Duke{},
	}

	f.阿怒波Anupa = anupa.DAnupa阿怒波
	f.阿怒波Anupa.SetParent(f)

	f.阿般提Dadhipadra = dadhipadra.DDadhipadra阿般提
	f.阿般提Dadhipadra.SetParent(f)

}

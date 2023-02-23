package britannia

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BritanniaEmpire interface {
	feud.Empire
	KEngland英格兰() england.EnglandKingdom
	KIreland爱尔兰() ireland.IrelandKingdom
	KScotland苏格兰() scotland.ScotlandKingdom
	KWales威尔士() wales.WalesKingdom
}

type 不列颠尼亚BritanniaEmpire struct {
	feud.BaseEmpire
	英格兰England  england.EnglandKingdom
	爱尔兰Ireland  ireland.IrelandKingdom
	苏格兰Scotland scotland.ScotlandKingdom
	威尔士Wales    wales.WalesKingdom
}

func (e *不列颠尼亚BritanniaEmpire) KEngland英格兰() england.EnglandKingdom {
	return e.英格兰England
}

func (e *不列颠尼亚BritanniaEmpire) KIreland爱尔兰() ireland.IrelandKingdom {
	return e.爱尔兰Ireland
}

func (e *不列颠尼亚BritanniaEmpire) KScotland苏格兰() scotland.ScotlandKingdom {
	return e.苏格兰Scotland
}

func (e *不列颠尼亚BritanniaEmpire) KWales威尔士() wales.WalesKingdom {
	return e.威尔士Wales
}

var EBritannia不列颠尼亚 BritanniaEmpire = &不列颠尼亚BritanniaEmpire{}

func init() {
	f := EBritannia不列颠尼亚.(*不列颠尼亚BritanniaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "britannia",
		TitleName: "不列颠尼亚",
		TitleCode: "e_britannia",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.英格兰England = england.KEngland英格兰
	f.英格兰England.SetParent(f)
	f.爱尔兰Ireland = ireland.KIreland爱尔兰
	f.爱尔兰Ireland.SetParent(f)
	f.苏格兰Scotland = scotland.KScotland苏格兰
	f.苏格兰Scotland.SetParent(f)
	f.威尔士Wales = wales.KWales威尔士
	f.威尔士Wales.SetParent(f)
}

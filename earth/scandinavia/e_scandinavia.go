package scandinavia

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/denmark"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/estonia"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/finland"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sweden"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ScandinaviaEmpire interface {
    feud.Empire
    KDenmark丹麦() 	denmark.DenmarkKingdom
    KEstonia爱沙尼亚() 	estonia.EstoniaKingdom
    KFinland芬兰() 	finland.FinlandKingdom
    KNorway挪威() 	norway.NorwayKingdom
    KSapmi拉普兰() 	sapmi.SapmiKingdom
    KSweden瑞典() 	sweden.SwedenKingdom
}

type 斯堪的纳维亚ScandinaviaEmpire struct {
	feud.BaseEmpire
	丹麦Denmark 	denmark.DenmarkKingdom
	爱沙尼亚Estonia 	estonia.EstoniaKingdom
	芬兰Finland 	finland.FinlandKingdom
	挪威Norway 	norway.NorwayKingdom
	拉普兰Sapmi 	sapmi.SapmiKingdom
	瑞典Sweden 	sweden.SwedenKingdom
}

func (e *斯堪的纳维亚ScandinaviaEmpire) KDenmark丹麦() denmark.DenmarkKingdom {
	return e.丹麦Denmark
}
    
func (e *斯堪的纳维亚ScandinaviaEmpire) KEstonia爱沙尼亚() estonia.EstoniaKingdom {
	return e.爱沙尼亚Estonia
}
    
func (e *斯堪的纳维亚ScandinaviaEmpire) KFinland芬兰() finland.FinlandKingdom {
	return e.芬兰Finland
}
    
func (e *斯堪的纳维亚ScandinaviaEmpire) KNorway挪威() norway.NorwayKingdom {
	return e.挪威Norway
}
    
func (e *斯堪的纳维亚ScandinaviaEmpire) KSapmi拉普兰() sapmi.SapmiKingdom {
	return e.拉普兰Sapmi
}
    
func (e *斯堪的纳维亚ScandinaviaEmpire) KSweden瑞典() sweden.SwedenKingdom {
	return e.瑞典Sweden
}
    
var EScandinavia斯堪的纳维亚 ScandinaviaEmpire = &斯堪的纳维亚ScandinaviaEmpire{}

func init() {
	f := EScandinavia斯堪的纳维亚.(*斯堪的纳维亚ScandinaviaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "scandinavia",
		TitleName: "斯堪的纳维亚",
		TitleCode: "e_scandinavia",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.丹麦Denmark = denmark.KDenmark丹麦
	f.丹麦Denmark.SetParent(f)
	f.爱沙尼亚Estonia = estonia.KEstonia爱沙尼亚
	f.爱沙尼亚Estonia.SetParent(f)
	f.芬兰Finland = finland.KFinland芬兰
	f.芬兰Finland.SetParent(f)
	f.挪威Norway = norway.KNorway挪威
	f.挪威Norway.SetParent(f)
	f.拉普兰Sapmi = sapmi.KSapmi拉普兰
	f.拉普兰Sapmi.SetParent(f)
	f.瑞典Sweden = sweden.KSweden瑞典
	f.瑞典Sweden.SetParent(f)
}

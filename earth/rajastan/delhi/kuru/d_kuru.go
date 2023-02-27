package kuru

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/kuru/delhi"
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/kuru/hastinapura"
	"github.com/thalesfu/CK2Commands/earth/rajastan/delhi/kuru/sthanisvara"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KuruDuke interface {
    feud.Duke
    CDelhi德里() 	delhi.DelhiCounty
    CHastinapura象城() 	hastinapura.HastinapuraCounty
    CSthanisvara萨他泥湿伐罗() 	sthanisvara.SthanisvaraCounty
}

type 俱卢KuruDuke struct {
	feud.BaseDuke
	德里Delhi 	delhi.DelhiCounty
	象城Hastinapura 	hastinapura.HastinapuraCounty
	萨他泥湿伐罗Sthanisvara 	sthanisvara.SthanisvaraCounty
}

func (d *俱卢KuruDuke) CDelhi德里() delhi.DelhiCounty {
	return d.德里Delhi
}
    
func (d *俱卢KuruDuke) CHastinapura象城() hastinapura.HastinapuraCounty {
	return d.象城Hastinapura
}
    
func (d *俱卢KuruDuke) CSthanisvara萨他泥湿伐罗() sthanisvara.SthanisvaraCounty {
	return d.萨他泥湿伐罗Sthanisvara
}
    
var DKuru俱卢 KuruDuke = &俱卢KuruDuke{}

func init() {
	f := DKuru俱卢.(*俱卢KuruDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kuru",
		TitleName: "俱卢",
		TitleCode: "d_kuru",
		Counties:  map[string]feud.County{},
	}

	f.德里Delhi = delhi.CDelhi德里
	f.德里Delhi.SetParent(f)
	
	f.象城Hastinapura = hastinapura.CHastinapura象城
	f.象城Hastinapura.SetParent(f)
	
	f.萨他泥湿伐罗Sthanisvara = sthanisvara.CSthanisvara萨他泥湿伐罗
	f.萨他泥湿伐罗Sthanisvara.SetParent(f)
	
}

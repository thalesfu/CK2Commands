package fes

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/fes/fes"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/fes/infa"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/fes/meknes"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/fes/muluja"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/fes/taza"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FesDuke interface {
    feud.Duke
    CFes非斯() 	fes.FesCounty
    CInfa安法() 	infa.InfaCounty
    CMeknes梅克内斯() 	meknes.MeknesCounty
    CMuluja穆卢耶() 	muluja.MulujaCounty
    CTaza塔扎() 	taza.TazaCounty
}

type 非斯FesDuke struct {
	feud.BaseDuke
	非斯Fes 	fes.FesCounty
	安法Infa 	infa.InfaCounty
	梅克内斯Meknes 	meknes.MeknesCounty
	穆卢耶Muluja 	muluja.MulujaCounty
	塔扎Taza 	taza.TazaCounty
}

func (d *非斯FesDuke) CFes非斯() fes.FesCounty {
	return d.非斯Fes
}
    
func (d *非斯FesDuke) CInfa安法() infa.InfaCounty {
	return d.安法Infa
}
    
func (d *非斯FesDuke) CMeknes梅克内斯() meknes.MeknesCounty {
	return d.梅克内斯Meknes
}
    
func (d *非斯FesDuke) CMuluja穆卢耶() muluja.MulujaCounty {
	return d.穆卢耶Muluja
}
    
func (d *非斯FesDuke) CTaza塔扎() taza.TazaCounty {
	return d.塔扎Taza
}
    
var DFes非斯 FesDuke = &非斯FesDuke{}

func init() {
	f := DFes非斯.(*非斯FesDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "fes",
		TitleName: "非斯",
		TitleCode: "d_fes",
		Counties:  map[string]feud.County{},
	}

	f.非斯Fes = fes.CFes非斯
	f.非斯Fes.SetParent(f)
	
	f.安法Infa = infa.CInfa安法
	f.安法Infa.SetParent(f)
	
	f.梅克内斯Meknes = meknes.CMeknes梅克内斯
	f.梅克内斯Meknes.SetParent(f)
	
	f.穆卢耶Muluja = muluja.CMuluja穆卢耶
	f.穆卢耶Muluja.SetParent(f)
	
	f.塔扎Taza = taza.CTaza塔扎
	f.塔扎Taza.SetParent(f)
	
}

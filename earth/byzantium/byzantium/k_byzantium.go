package byzantium

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/achaia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/aegean_islands"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/athens"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/krete"
	"github.com/thalesfu/CK2Commands/earth/byzantium/byzantium/thessalonika"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ByzantiumKingdom interface {
    feud.Kingdom
    DAchaia亚该亚() 	achaia.AchaiaDuke
    DAegean_islands爱琴群岛() 	aegean_islands.Aegean_islandsDuke
    DAthens希腊() 	athens.AthensDuke
    DKrete克里特() 	krete.KreteDuke
    DThessalonika帖撒罗尼迦() 	thessalonika.ThessalonikaDuke
}

type 希腊ByzantiumKingdom struct {
	feud.BaseKingdom
	亚该亚Achaia 	achaia.AchaiaDuke
	爱琴群岛Aegean_islands 	aegean_islands.Aegean_islandsDuke
	希腊Athens 	athens.AthensDuke
	克里特Krete 	krete.KreteDuke
	帖撒罗尼迦Thessalonika 	thessalonika.ThessalonikaDuke
}

func (k *希腊ByzantiumKingdom) DAchaia亚该亚() achaia.AchaiaDuke {
	return k.亚该亚Achaia
}
    
func (k *希腊ByzantiumKingdom) DAegean_islands爱琴群岛() aegean_islands.Aegean_islandsDuke {
	return k.爱琴群岛Aegean_islands
}
    
func (k *希腊ByzantiumKingdom) DAthens希腊() athens.AthensDuke {
	return k.希腊Athens
}
    
func (k *希腊ByzantiumKingdom) DKrete克里特() krete.KreteDuke {
	return k.克里特Krete
}
    
func (k *希腊ByzantiumKingdom) DThessalonika帖撒罗尼迦() thessalonika.ThessalonikaDuke {
	return k.帖撒罗尼迦Thessalonika
}
    
var KByzantium希腊 ByzantiumKingdom = &希腊ByzantiumKingdom{}

func init() {
	f := KByzantium希腊.(*希腊ByzantiumKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "byzantium",
		TitleName: "希腊",
		TitleCode: "k_byzantium",
		Dukes:  map[string]feud.Duke{},
	}

	f.亚该亚Achaia = achaia.DAchaia亚该亚
	f.亚该亚Achaia.SetParent(f)
	
	f.爱琴群岛Aegean_islands = aegean_islands.DAegean_islands爱琴群岛
	f.爱琴群岛Aegean_islands.SetParent(f)
	
	f.希腊Athens = athens.DAthens希腊
	f.希腊Athens.SetParent(f)
	
	f.克里特Krete = krete.DKrete克里特
	f.克里特Krete.SetParent(f)
	
	f.帖撒罗尼迦Thessalonika = thessalonika.DThessalonika帖撒罗尼迦
	f.帖撒罗尼迦Thessalonika.SetParent(f)
	
}

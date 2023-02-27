package maris

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/maris/galich_mersky"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/maris/hlynov"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/maris/merya"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MarisDuke interface {
    feud.Duke
    CGalich_mersky加利奇梅尔斯基() 	galich_mersky.Galich_merskyCounty
    CHlynov赫雷诺夫() 	hlynov.HlynovCounty
    CMerya马里() 	merya.MeryaCounty
}

type 马里MarisDuke struct {
	feud.BaseDuke
	加利奇梅尔斯基Galich_mersky 	galich_mersky.Galich_merskyCounty
	赫雷诺夫Hlynov 	hlynov.HlynovCounty
	马里Merya 	merya.MeryaCounty
}

func (d *马里MarisDuke) CGalich_mersky加利奇梅尔斯基() galich_mersky.Galich_merskyCounty {
	return d.加利奇梅尔斯基Galich_mersky
}
    
func (d *马里MarisDuke) CHlynov赫雷诺夫() hlynov.HlynovCounty {
	return d.赫雷诺夫Hlynov
}
    
func (d *马里MarisDuke) CMerya马里() merya.MeryaCounty {
	return d.马里Merya
}
    
var DMaris马里 MarisDuke = &马里MarisDuke{}

func init() {
	f := DMaris马里.(*马里MarisDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "maris",
		TitleName: "马里",
		TitleCode: "d_maris",
		Counties:  map[string]feud.County{},
	}

	f.加利奇梅尔斯基Galich_mersky = galich_mersky.CGalich_mersky加利奇梅尔斯基
	f.加利奇梅尔斯基Galich_mersky.SetParent(f)
	
	f.赫雷诺夫Hlynov = hlynov.CHlynov赫雷诺夫
	f.赫雷诺夫Hlynov.SetParent(f)
	
	f.马里Merya = merya.CMerya马里
	f.马里Merya.SetParent(f)
	
}

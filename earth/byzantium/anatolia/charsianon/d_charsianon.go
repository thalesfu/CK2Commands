package charsianon

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/charsianon/charsianon"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/charsianon/galatia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/anatolia/charsianon/kaisereia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CharsianonDuke interface {
    feud.Duke
    CCharsianon哈耳西亚农() 	charsianon.CharsianonCounty
    CGalatia加拉太() 	galatia.GalatiaCounty
    CKaisereia凯撒利亚() 	kaisereia.KaisereiaCounty
}

type 哈耳西亚农CharsianonDuke struct {
	feud.BaseDuke
	哈耳西亚农Charsianon 	charsianon.CharsianonCounty
	加拉太Galatia 	galatia.GalatiaCounty
	凯撒利亚Kaisereia 	kaisereia.KaisereiaCounty
}

func (d *哈耳西亚农CharsianonDuke) CCharsianon哈耳西亚农() charsianon.CharsianonCounty {
	return d.哈耳西亚农Charsianon
}
    
func (d *哈耳西亚农CharsianonDuke) CGalatia加拉太() galatia.GalatiaCounty {
	return d.加拉太Galatia
}
    
func (d *哈耳西亚农CharsianonDuke) CKaisereia凯撒利亚() kaisereia.KaisereiaCounty {
	return d.凯撒利亚Kaisereia
}
    
var DCharsianon哈耳西亚农 CharsianonDuke = &哈耳西亚农CharsianonDuke{}

func init() {
	f := DCharsianon哈耳西亚农.(*哈耳西亚农CharsianonDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "charsianon",
		TitleName: "哈耳西亚农",
		TitleCode: "d_charsianon",
		Counties:  map[string]feud.County{},
	}

	f.哈耳西亚农Charsianon = charsianon.CCharsianon哈耳西亚农
	f.哈耳西亚农Charsianon.SetParent(f)
	
	f.加拉太Galatia = galatia.CGalatia加拉太
	f.加拉太Galatia.SetParent(f)
	
	f.凯撒利亚Kaisereia = kaisereia.CKaisereia凯撒利亚
	f.凯撒利亚Kaisereia.SetParent(f)
	
}

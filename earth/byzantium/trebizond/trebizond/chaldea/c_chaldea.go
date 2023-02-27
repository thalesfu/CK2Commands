package chaldea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChaldeaCounty interface {
    feud.County
    BCamachus卡马库斯() 	feud.Barony
    BCotyora索图拉() 	feud.Barony
    BHeracleopolis赫拉克莱奥波利斯() 	feud.Barony
    BIbora伊博拉() 	feud.Barony
    BKerasous凯拉苏斯() 	feud.Barony
    BPodandos波丹多斯() 	feud.Barony
    BTilgarimo提格里莫() 	feud.Barony
}

type 迦勒底ChaldeaCounty struct {
	feud.BaseCounty
	卡马库斯Camachus 	feud.Barony
	索图拉Cotyora 	feud.Barony
	赫拉克莱奥波利斯Heracleopolis 	feud.Barony
	伊博拉Ibora 	feud.Barony
	凯拉苏斯Kerasous 	feud.Barony
	波丹多斯Podandos 	feud.Barony
	提格里莫Tilgarimo 	feud.Barony
}

func (c *迦勒底ChaldeaCounty) BCamachus卡马库斯() feud.Barony {
	return c.卡马库斯Camachus
}
    
func (c *迦勒底ChaldeaCounty) BCotyora索图拉() feud.Barony {
	return c.索图拉Cotyora
}
    
func (c *迦勒底ChaldeaCounty) BHeracleopolis赫拉克莱奥波利斯() feud.Barony {
	return c.赫拉克莱奥波利斯Heracleopolis
}
    
func (c *迦勒底ChaldeaCounty) BIbora伊博拉() feud.Barony {
	return c.伊博拉Ibora
}
    
func (c *迦勒底ChaldeaCounty) BKerasous凯拉苏斯() feud.Barony {
	return c.凯拉苏斯Kerasous
}
    
func (c *迦勒底ChaldeaCounty) BPodandos波丹多斯() feud.Barony {
	return c.波丹多斯Podandos
}
    
func (c *迦勒底ChaldeaCounty) BTilgarimo提格里莫() feud.Barony {
	return c.提格里莫Tilgarimo
}
    
var CChaldea迦勒底 ChaldeaCounty = &迦勒底ChaldeaCounty{}

func init() {
	f := CChaldea迦勒底.(*迦勒底ChaldeaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "705",
		Title:     "chaldea",
		TitleName: "迦勒底",
		TitleCode: "c_chaldea",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡马库斯Camachus = BCamachus卡马库斯
	f.卡马库斯Camachus.SetParent(f)
	
	f.索图拉Cotyora = BCotyora索图拉
	f.索图拉Cotyora.SetParent(f)
	
	f.赫拉克莱奥波利斯Heracleopolis = BHeracleopolis赫拉克莱奥波利斯
	f.赫拉克莱奥波利斯Heracleopolis.SetParent(f)
	
	f.伊博拉Ibora = BIbora伊博拉
	f.伊博拉Ibora.SetParent(f)
	
	f.凯拉苏斯Kerasous = BKerasous凯拉苏斯
	f.凯拉苏斯Kerasous.SetParent(f)
	
	f.波丹多斯Podandos = BPodandos波丹多斯
	f.波丹多斯Podandos.SetParent(f)
	
	f.提格里莫Tilgarimo = BTilgarimo提格里莫
	f.提格里莫Tilgarimo.SetParent(f)
	
}

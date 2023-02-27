package euphrates

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EuphratesCounty interface {
    feud.County
    BAmirli阿米力() 	feud.Barony
    BBalad拜莱德() 	feud.Barony
    BBayji拜吉() 	feud.Barony
    BDujail杜贾尔() 	feud.Barony
    BFaris法里斯() 	feud.Barony
    BIshaqi伊斯哈奇() 	feud.Barony
    BSamarra萨迈拉() 	feud.Barony
    BTagrit提克里特() 	feud.Barony
}

type 萨迈拉EuphratesCounty struct {
	feud.BaseCounty
	阿米力Amirli 	feud.Barony
	拜莱德Balad 	feud.Barony
	拜吉Bayji 	feud.Barony
	杜贾尔Dujail 	feud.Barony
	法里斯Faris 	feud.Barony
	伊斯哈奇Ishaqi 	feud.Barony
	萨迈拉Samarra 	feud.Barony
	提克里特Tagrit 	feud.Barony
}

func (c *萨迈拉EuphratesCounty) BAmirli阿米力() feud.Barony {
	return c.阿米力Amirli
}
    
func (c *萨迈拉EuphratesCounty) BBalad拜莱德() feud.Barony {
	return c.拜莱德Balad
}
    
func (c *萨迈拉EuphratesCounty) BBayji拜吉() feud.Barony {
	return c.拜吉Bayji
}
    
func (c *萨迈拉EuphratesCounty) BDujail杜贾尔() feud.Barony {
	return c.杜贾尔Dujail
}
    
func (c *萨迈拉EuphratesCounty) BFaris法里斯() feud.Barony {
	return c.法里斯Faris
}
    
func (c *萨迈拉EuphratesCounty) BIshaqi伊斯哈奇() feud.Barony {
	return c.伊斯哈奇Ishaqi
}
    
func (c *萨迈拉EuphratesCounty) BSamarra萨迈拉() feud.Barony {
	return c.萨迈拉Samarra
}
    
func (c *萨迈拉EuphratesCounty) BTagrit提克里特() feud.Barony {
	return c.提克里特Tagrit
}
    
var CEuphrates萨迈拉 EuphratesCounty = &萨迈拉EuphratesCounty{}

func init() {
	f := CEuphrates萨迈拉.(*萨迈拉EuphratesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "696",
		Title:     "euphrates",
		TitleName: "萨迈拉",
		TitleCode: "c_euphrates",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿米力Amirli = BAmirli阿米力
	f.阿米力Amirli.SetParent(f)
	
	f.拜莱德Balad = BBalad拜莱德
	f.拜莱德Balad.SetParent(f)
	
	f.拜吉Bayji = BBayji拜吉
	f.拜吉Bayji.SetParent(f)
	
	f.杜贾尔Dujail = BDujail杜贾尔
	f.杜贾尔Dujail.SetParent(f)
	
	f.法里斯Faris = BFaris法里斯
	f.法里斯Faris.SetParent(f)
	
	f.伊斯哈奇Ishaqi = BIshaqi伊斯哈奇
	f.伊斯哈奇Ishaqi.SetParent(f)
	
	f.萨迈拉Samarra = BSamarra萨迈拉
	f.萨迈拉Samarra.SetParent(f)
	
	f.提克里特Tagrit = BTagrit提克里特
	f.提克里特Tagrit.SetParent(f)
	
}

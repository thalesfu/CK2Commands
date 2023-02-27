package dotawo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DotawoCounty interface {
    feud.County
    BBallana巴拉那() 	feud.Barony
    BBuhen布亨() 	feud.Barony
    BDotawo多塔沃() 	feud.Barony
    BFaras法拉斯() 	feud.Barony
    BJebeladda杰贝腊达() 	feud.Barony
    BPremnis普雷姆尼斯() 	feud.Barony
    BQasribrim卡斯利布林() 	feud.Barony
}

type 多塔沃DotawoCounty struct {
	feud.BaseCounty
	巴拉那Ballana 	feud.Barony
	布亨Buhen 	feud.Barony
	多塔沃Dotawo 	feud.Barony
	法拉斯Faras 	feud.Barony
	杰贝腊达Jebeladda 	feud.Barony
	普雷姆尼斯Premnis 	feud.Barony
	卡斯利布林Qasribrim 	feud.Barony
}

func (c *多塔沃DotawoCounty) BBallana巴拉那() feud.Barony {
	return c.巴拉那Ballana
}
    
func (c *多塔沃DotawoCounty) BBuhen布亨() feud.Barony {
	return c.布亨Buhen
}
    
func (c *多塔沃DotawoCounty) BDotawo多塔沃() feud.Barony {
	return c.多塔沃Dotawo
}
    
func (c *多塔沃DotawoCounty) BFaras法拉斯() feud.Barony {
	return c.法拉斯Faras
}
    
func (c *多塔沃DotawoCounty) BJebeladda杰贝腊达() feud.Barony {
	return c.杰贝腊达Jebeladda
}
    
func (c *多塔沃DotawoCounty) BPremnis普雷姆尼斯() feud.Barony {
	return c.普雷姆尼斯Premnis
}
    
func (c *多塔沃DotawoCounty) BQasribrim卡斯利布林() feud.Barony {
	return c.卡斯利布林Qasribrim
}
    
var CDotawo多塔沃 DotawoCounty = &多塔沃DotawoCounty{}

func init() {
	f := CDotawo多塔沃.(*多塔沃DotawoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1282",
		Title:     "dotawo",
		TitleName: "多塔沃",
		TitleCode: "c_dotawo",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴拉那Ballana = BBallana巴拉那
	f.巴拉那Ballana.SetParent(f)
	
	f.布亨Buhen = BBuhen布亨
	f.布亨Buhen.SetParent(f)
	
	f.多塔沃Dotawo = BDotawo多塔沃
	f.多塔沃Dotawo.SetParent(f)
	
	f.法拉斯Faras = BFaras法拉斯
	f.法拉斯Faras.SetParent(f)
	
	f.杰贝腊达Jebeladda = BJebeladda杰贝腊达
	f.杰贝腊达Jebeladda.SetParent(f)
	
	f.普雷姆尼斯Premnis = BPremnis普雷姆尼斯
	f.普雷姆尼斯Premnis.SetParent(f)
	
	f.卡斯利布林Qasribrim = BQasribrim卡斯利布林
	f.卡斯利布林Qasribrim.SetParent(f)
	
}

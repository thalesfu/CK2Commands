package nevers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NeversCounty interface {
    feud.County
    BChateauchinon希农堡() 	feud.Barony
    BClamecy克拉姆西() 	feud.Barony
    BCourtenay库尔特奈() 	feud.Barony
    BDonzy栋济() 	feud.Barony
    BLacharite拉沙里泰() 	feud.Barony
    BNevers讷韦尔() 	feud.Barony
    BVandenesse旺德内斯() 	feud.Barony
}

type 讷韦尔NeversCounty struct {
	feud.BaseCounty
	希农堡Chateauchinon 	feud.Barony
	克拉姆西Clamecy 	feud.Barony
	库尔特奈Courtenay 	feud.Barony
	栋济Donzy 	feud.Barony
	拉沙里泰Lacharite 	feud.Barony
	讷韦尔Nevers 	feud.Barony
	旺德内斯Vandenesse 	feud.Barony
}

func (c *讷韦尔NeversCounty) BChateauchinon希农堡() feud.Barony {
	return c.希农堡Chateauchinon
}
    
func (c *讷韦尔NeversCounty) BClamecy克拉姆西() feud.Barony {
	return c.克拉姆西Clamecy
}
    
func (c *讷韦尔NeversCounty) BCourtenay库尔特奈() feud.Barony {
	return c.库尔特奈Courtenay
}
    
func (c *讷韦尔NeversCounty) BDonzy栋济() feud.Barony {
	return c.栋济Donzy
}
    
func (c *讷韦尔NeversCounty) BLacharite拉沙里泰() feud.Barony {
	return c.拉沙里泰Lacharite
}
    
func (c *讷韦尔NeversCounty) BNevers讷韦尔() feud.Barony {
	return c.讷韦尔Nevers
}
    
func (c *讷韦尔NeversCounty) BVandenesse旺德内斯() feud.Barony {
	return c.旺德内斯Vandenesse
}
    
var CNevers讷韦尔 NeversCounty = &讷韦尔NeversCounty{}

func init() {
	f := CNevers讷韦尔.(*讷韦尔NeversCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "137",
		Title:     "nevers",
		TitleName: "讷韦尔",
		TitleCode: "c_nevers",
		Baronies:  map[string]feud.Barony{},
	}

	f.希农堡Chateauchinon = BChateauchinon希农堡
	f.希农堡Chateauchinon.SetParent(f)
	
	f.克拉姆西Clamecy = BClamecy克拉姆西
	f.克拉姆西Clamecy.SetParent(f)
	
	f.库尔特奈Courtenay = BCourtenay库尔特奈
	f.库尔特奈Courtenay.SetParent(f)
	
	f.栋济Donzy = BDonzy栋济
	f.栋济Donzy.SetParent(f)
	
	f.拉沙里泰Lacharite = BLacharite拉沙里泰
	f.拉沙里泰Lacharite.SetParent(f)
	
	f.讷韦尔Nevers = BNevers讷韦尔
	f.讷韦尔Nevers.SetParent(f)
	
	f.旺德内斯Vandenesse = BVandenesse旺德内斯
	f.旺德内斯Vandenesse.SetParent(f)
	
}

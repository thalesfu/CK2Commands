package mustang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MustangCounty interface {
    feud.County
    BChhonhup冲哈() 	feud.Barony
    BChhoser冲瑟尔() 	feud.Barony
    BChhusang曲桑() 	feud.Barony
    BDrak_gyawu扎甲吾() 	feud.Barony
    BKagbeni卡贝尼() 	feud.Barony
    BLo珞() 	feud.Barony
    BTetang谛当() 	feud.Barony
}

type 木斯塘MustangCounty struct {
	feud.BaseCounty
	冲哈Chhonhup 	feud.Barony
	冲瑟尔Chhoser 	feud.Barony
	曲桑Chhusang 	feud.Barony
	扎甲吾Drak_gyawu 	feud.Barony
	卡贝尼Kagbeni 	feud.Barony
	珞Lo 	feud.Barony
	谛当Tetang 	feud.Barony
}

func (c *木斯塘MustangCounty) BChhonhup冲哈() feud.Barony {
	return c.冲哈Chhonhup
}
    
func (c *木斯塘MustangCounty) BChhoser冲瑟尔() feud.Barony {
	return c.冲瑟尔Chhoser
}
    
func (c *木斯塘MustangCounty) BChhusang曲桑() feud.Barony {
	return c.曲桑Chhusang
}
    
func (c *木斯塘MustangCounty) BDrak_gyawu扎甲吾() feud.Barony {
	return c.扎甲吾Drak_gyawu
}
    
func (c *木斯塘MustangCounty) BKagbeni卡贝尼() feud.Barony {
	return c.卡贝尼Kagbeni
}
    
func (c *木斯塘MustangCounty) BLo珞() feud.Barony {
	return c.珞Lo
}
    
func (c *木斯塘MustangCounty) BTetang谛当() feud.Barony {
	return c.谛当Tetang
}
    
var CMustang木斯塘 MustangCounty = &木斯塘MustangCounty{}

func init() {
	f := CMustang木斯塘.(*木斯塘MustangCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1551",
		Title:     "mustang",
		TitleName: "木斯塘",
		TitleCode: "c_mustang",
		Baronies:  map[string]feud.Barony{},
	}

	f.冲哈Chhonhup = BChhonhup冲哈
	f.冲哈Chhonhup.SetParent(f)
	
	f.冲瑟尔Chhoser = BChhoser冲瑟尔
	f.冲瑟尔Chhoser.SetParent(f)
	
	f.曲桑Chhusang = BChhusang曲桑
	f.曲桑Chhusang.SetParent(f)
	
	f.扎甲吾Drak_gyawu = BDrak_gyawu扎甲吾
	f.扎甲吾Drak_gyawu.SetParent(f)
	
	f.卡贝尼Kagbeni = BKagbeni卡贝尼
	f.卡贝尼Kagbeni.SetParent(f)
	
	f.珞Lo = BLo珞
	f.珞Lo.SetParent(f)
	
	f.谛当Tetang = BTetang谛当
	f.谛当Tetang.SetParent(f)
	
}

package lahur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LahurCounty interface {
    feud.County
    BChunian楚里安() 	feud.Barony
    BData_darbar达塔达巴尔() 	feud.Barony
    BKasur卡苏尔() 	feud.Barony
    BLahur拉合尔() 	feud.Barony
    BMehdiabad迈赫迪耶巴德() 	feud.Barony
    BSamnabad萨姆那巴德() 	feud.Barony
    BSangla桑格拉() 	feud.Barony
}

type 拉合尔LahurCounty struct {
	feud.BaseCounty
	楚里安Chunian 	feud.Barony
	达塔达巴尔Data_darbar 	feud.Barony
	卡苏尔Kasur 	feud.Barony
	拉合尔Lahur 	feud.Barony
	迈赫迪耶巴德Mehdiabad 	feud.Barony
	萨姆那巴德Samnabad 	feud.Barony
	桑格拉Sangla 	feud.Barony
}

func (c *拉合尔LahurCounty) BChunian楚里安() feud.Barony {
	return c.楚里安Chunian
}
    
func (c *拉合尔LahurCounty) BData_darbar达塔达巴尔() feud.Barony {
	return c.达塔达巴尔Data_darbar
}
    
func (c *拉合尔LahurCounty) BKasur卡苏尔() feud.Barony {
	return c.卡苏尔Kasur
}
    
func (c *拉合尔LahurCounty) BLahur拉合尔() feud.Barony {
	return c.拉合尔Lahur
}
    
func (c *拉合尔LahurCounty) BMehdiabad迈赫迪耶巴德() feud.Barony {
	return c.迈赫迪耶巴德Mehdiabad
}
    
func (c *拉合尔LahurCounty) BSamnabad萨姆那巴德() feud.Barony {
	return c.萨姆那巴德Samnabad
}
    
func (c *拉合尔LahurCounty) BSangla桑格拉() feud.Barony {
	return c.桑格拉Sangla
}
    
var CLahur拉合尔 LahurCounty = &拉合尔LahurCounty{}

func init() {
	f := CLahur拉合尔.(*拉合尔LahurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1362",
		Title:     "lahur",
		TitleName: "拉合尔",
		TitleCode: "c_lahur",
		Baronies:  map[string]feud.Barony{},
	}

	f.楚里安Chunian = BChunian楚里安
	f.楚里安Chunian.SetParent(f)
	
	f.达塔达巴尔Data_darbar = BData_darbar达塔达巴尔
	f.达塔达巴尔Data_darbar.SetParent(f)
	
	f.卡苏尔Kasur = BKasur卡苏尔
	f.卡苏尔Kasur.SetParent(f)
	
	f.拉合尔Lahur = BLahur拉合尔
	f.拉合尔Lahur.SetParent(f)
	
	f.迈赫迪耶巴德Mehdiabad = BMehdiabad迈赫迪耶巴德
	f.迈赫迪耶巴德Mehdiabad.SetParent(f)
	
	f.萨姆那巴德Samnabad = BSamnabad萨姆那巴德
	f.萨姆那巴德Samnabad.SetParent(f)
	
	f.桑格拉Sangla = BSangla桑格拉
	f.桑格拉Sangla.SetParent(f)
	
}

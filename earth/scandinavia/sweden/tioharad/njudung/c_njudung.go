package njudung

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NjudungCounty interface {
    feud.County
    BFemtinge费姆廷厄() 	feud.Barony
    BHultaby胡尔塔比() 	feud.Barony
    BKomstad科姆斯塔德() 	feud.Barony
    BKornstad科恩斯塔德() 	feud.Barony
    BLjunga永阿() 	feud.Barony
    BSavsjo赛夫舍() 	feud.Barony
    BVitala维塔拉() 	feud.Barony
}

type 纽东NjudungCounty struct {
	feud.BaseCounty
	费姆廷厄Femtinge 	feud.Barony
	胡尔塔比Hultaby 	feud.Barony
	科姆斯塔德Komstad 	feud.Barony
	科恩斯塔德Kornstad 	feud.Barony
	永阿Ljunga 	feud.Barony
	赛夫舍Savsjo 	feud.Barony
	维塔拉Vitala 	feud.Barony
}

func (c *纽东NjudungCounty) BFemtinge费姆廷厄() feud.Barony {
	return c.费姆廷厄Femtinge
}
    
func (c *纽东NjudungCounty) BHultaby胡尔塔比() feud.Barony {
	return c.胡尔塔比Hultaby
}
    
func (c *纽东NjudungCounty) BKomstad科姆斯塔德() feud.Barony {
	return c.科姆斯塔德Komstad
}
    
func (c *纽东NjudungCounty) BKornstad科恩斯塔德() feud.Barony {
	return c.科恩斯塔德Kornstad
}
    
func (c *纽东NjudungCounty) BLjunga永阿() feud.Barony {
	return c.永阿Ljunga
}
    
func (c *纽东NjudungCounty) BSavsjo赛夫舍() feud.Barony {
	return c.赛夫舍Savsjo
}
    
func (c *纽东NjudungCounty) BVitala维塔拉() feud.Barony {
	return c.维塔拉Vitala
}
    
var CNjudung纽东 NjudungCounty = &纽东NjudungCounty{}

func init() {
	f := CNjudung纽东.(*纽东NjudungCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1703",
		Title:     "njudung",
		TitleName: "纽东",
		TitleCode: "c_njudung",
		Baronies:  map[string]feud.Barony{},
	}

	f.费姆廷厄Femtinge = BFemtinge费姆廷厄
	f.费姆廷厄Femtinge.SetParent(f)
	
	f.胡尔塔比Hultaby = BHultaby胡尔塔比
	f.胡尔塔比Hultaby.SetParent(f)
	
	f.科姆斯塔德Komstad = BKomstad科姆斯塔德
	f.科姆斯塔德Komstad.SetParent(f)
	
	f.科恩斯塔德Kornstad = BKornstad科恩斯塔德
	f.科恩斯塔德Kornstad.SetParent(f)
	
	f.永阿Ljunga = BLjunga永阿
	f.永阿Ljunga.SetParent(f)
	
	f.赛夫舍Savsjo = BSavsjo赛夫舍
	f.赛夫舍Savsjo.SetParent(f)
	
	f.维塔拉Vitala = BVitala维塔拉
	f.维塔拉Vitala.SetParent(f)
	
}

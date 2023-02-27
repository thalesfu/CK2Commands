package amida

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmidaCounty interface {
    feud.County
    BAmida阿米达() 	feud.Barony
    BEgil埃伊尔() 	feud.Barony
    BHazretisuleymancamii哈兹拉特苏莱曼寺() 	feud.Barony
    BIdtodyolataloho埃塔尤拉特拉霍() 	feud.Barony
    BKiaburc奇亚布斯() 	feud.Barony
    BMayyafarikin迈亚法里金() 	feud.Barony
    BSuraamede苏拉阿梅德() 	feud.Barony
    BUlucamii乌鲁贾米() 	feud.Barony
}

type 阿米达AmidaCounty struct {
	feud.BaseCounty
	阿米达Amida 	feud.Barony
	埃伊尔Egil 	feud.Barony
	哈兹拉特苏莱曼寺Hazretisuleymancamii 	feud.Barony
	埃塔尤拉特拉霍Idtodyolataloho 	feud.Barony
	奇亚布斯Kiaburc 	feud.Barony
	迈亚法里金Mayyafarikin 	feud.Barony
	苏拉阿梅德Suraamede 	feud.Barony
	乌鲁贾米Ulucamii 	feud.Barony
}

func (c *阿米达AmidaCounty) BAmida阿米达() feud.Barony {
	return c.阿米达Amida
}
    
func (c *阿米达AmidaCounty) BEgil埃伊尔() feud.Barony {
	return c.埃伊尔Egil
}
    
func (c *阿米达AmidaCounty) BHazretisuleymancamii哈兹拉特苏莱曼寺() feud.Barony {
	return c.哈兹拉特苏莱曼寺Hazretisuleymancamii
}
    
func (c *阿米达AmidaCounty) BIdtodyolataloho埃塔尤拉特拉霍() feud.Barony {
	return c.埃塔尤拉特拉霍Idtodyolataloho
}
    
func (c *阿米达AmidaCounty) BKiaburc奇亚布斯() feud.Barony {
	return c.奇亚布斯Kiaburc
}
    
func (c *阿米达AmidaCounty) BMayyafarikin迈亚法里金() feud.Barony {
	return c.迈亚法里金Mayyafarikin
}
    
func (c *阿米达AmidaCounty) BSuraamede苏拉阿梅德() feud.Barony {
	return c.苏拉阿梅德Suraamede
}
    
func (c *阿米达AmidaCounty) BUlucamii乌鲁贾米() feud.Barony {
	return c.乌鲁贾米Ulucamii
}
    
var CAmida阿米达 AmidaCounty = &阿米达AmidaCounty{}

func init() {
	f := CAmida阿米达.(*阿米达AmidaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "683",
		Title:     "amida",
		TitleName: "阿米达",
		TitleCode: "c_amida",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿米达Amida = BAmida阿米达
	f.阿米达Amida.SetParent(f)
	
	f.埃伊尔Egil = BEgil埃伊尔
	f.埃伊尔Egil.SetParent(f)
	
	f.哈兹拉特苏莱曼寺Hazretisuleymancamii = BHazretisuleymancamii哈兹拉特苏莱曼寺
	f.哈兹拉特苏莱曼寺Hazretisuleymancamii.SetParent(f)
	
	f.埃塔尤拉特拉霍Idtodyolataloho = BIdtodyolataloho埃塔尤拉特拉霍
	f.埃塔尤拉特拉霍Idtodyolataloho.SetParent(f)
	
	f.奇亚布斯Kiaburc = BKiaburc奇亚布斯
	f.奇亚布斯Kiaburc.SetParent(f)
	
	f.迈亚法里金Mayyafarikin = BMayyafarikin迈亚法里金
	f.迈亚法里金Mayyafarikin.SetParent(f)
	
	f.苏拉阿梅德Suraamede = BSuraamede苏拉阿梅德
	f.苏拉阿梅德Suraamede.SetParent(f)
	
	f.乌鲁贾米Ulucamii = BUlucamii乌鲁贾米
	f.乌鲁贾米Ulucamii.SetParent(f)
	
}

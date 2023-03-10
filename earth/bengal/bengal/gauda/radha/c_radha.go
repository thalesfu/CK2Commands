package radha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RadhaCounty interface {
    feud.County
    BGopbhum瞿波部摩() 	feud.Barony
    BHirapur醯罗补罗() 	feud.Barony
    BIchhai_ghoshgarh伊车耶瞿沙姞利呬() 	feud.Barony
    BKalyaneshwari迦梨耶泥湿伐罗() 	feud.Barony
    BKatwa揭婆() 	feud.Barony
    BLakhnor落诃奴罗() 	feud.Barony
    BRampur_boalia罗摩补罗菩阿梨阿() 	feud.Barony
}

type 罗陀RadhaCounty struct {
	feud.BaseCounty
	瞿波部摩Gopbhum 	feud.Barony
	醯罗补罗Hirapur 	feud.Barony
	伊车耶瞿沙姞利呬Ichhai_ghoshgarh 	feud.Barony
	迦梨耶泥湿伐罗Kalyaneshwari 	feud.Barony
	揭婆Katwa 	feud.Barony
	落诃奴罗Lakhnor 	feud.Barony
	罗摩补罗菩阿梨阿Rampur_boalia 	feud.Barony
}

func (c *罗陀RadhaCounty) BGopbhum瞿波部摩() feud.Barony {
	return c.瞿波部摩Gopbhum
}
    
func (c *罗陀RadhaCounty) BHirapur醯罗补罗() feud.Barony {
	return c.醯罗补罗Hirapur
}
    
func (c *罗陀RadhaCounty) BIchhai_ghoshgarh伊车耶瞿沙姞利呬() feud.Barony {
	return c.伊车耶瞿沙姞利呬Ichhai_ghoshgarh
}
    
func (c *罗陀RadhaCounty) BKalyaneshwari迦梨耶泥湿伐罗() feud.Barony {
	return c.迦梨耶泥湿伐罗Kalyaneshwari
}
    
func (c *罗陀RadhaCounty) BKatwa揭婆() feud.Barony {
	return c.揭婆Katwa
}
    
func (c *罗陀RadhaCounty) BLakhnor落诃奴罗() feud.Barony {
	return c.落诃奴罗Lakhnor
}
    
func (c *罗陀RadhaCounty) BRampur_boalia罗摩补罗菩阿梨阿() feud.Barony {
	return c.罗摩补罗菩阿梨阿Rampur_boalia
}
    
var CRadha罗陀 RadhaCounty = &罗陀RadhaCounty{}

func init() {
	f := CRadha罗陀.(*罗陀RadhaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1242",
		Title:     "radha",
		TitleName: "罗陀",
		TitleCode: "c_radha",
		Baronies:  map[string]feud.Barony{},
	}

	f.瞿波部摩Gopbhum = BGopbhum瞿波部摩
	f.瞿波部摩Gopbhum.SetParent(f)
	
	f.醯罗补罗Hirapur = BHirapur醯罗补罗
	f.醯罗补罗Hirapur.SetParent(f)
	
	f.伊车耶瞿沙姞利呬Ichhai_ghoshgarh = BIchhai_ghoshgarh伊车耶瞿沙姞利呬
	f.伊车耶瞿沙姞利呬Ichhai_ghoshgarh.SetParent(f)
	
	f.迦梨耶泥湿伐罗Kalyaneshwari = BKalyaneshwari迦梨耶泥湿伐罗
	f.迦梨耶泥湿伐罗Kalyaneshwari.SetParent(f)
	
	f.揭婆Katwa = BKatwa揭婆
	f.揭婆Katwa.SetParent(f)
	
	f.落诃奴罗Lakhnor = BLakhnor落诃奴罗
	f.落诃奴罗Lakhnor.SetParent(f)
	
	f.罗摩补罗菩阿梨阿Rampur_boalia = BRampur_boalia罗摩补罗菩阿梨阿
	f.罗摩补罗菩阿梨阿Rampur_boalia.SetParent(f)
	
}

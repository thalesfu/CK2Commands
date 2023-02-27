package sambalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SambalpurCounty interface {
    feud.County
    BBimaleswar毗摩丽湿伐罗() 	feud.Barony
    BManeswar摩泥湿伐罗() 	feud.Barony
    BRajgangpur罗阇恒伽补罗() 	feud.Barony
    BRanipur罗尼补罗() 	feud.Barony
    BSamaleswari娑摩丽湿伐利() 	feud.Barony
    BSambalpur三婆罗补罗() 	feud.Barony
    BSarugani娑楼迦尼() 	feud.Barony
}

type 三婆罗补罗SambalpurCounty struct {
	feud.BaseCounty
	毗摩丽湿伐罗Bimaleswar 	feud.Barony
	摩泥湿伐罗Maneswar 	feud.Barony
	罗阇恒伽补罗Rajgangpur 	feud.Barony
	罗尼补罗Ranipur 	feud.Barony
	娑摩丽湿伐利Samaleswari 	feud.Barony
	三婆罗补罗Sambalpur 	feud.Barony
	娑楼迦尼Sarugani 	feud.Barony
}

func (c *三婆罗补罗SambalpurCounty) BBimaleswar毗摩丽湿伐罗() feud.Barony {
	return c.毗摩丽湿伐罗Bimaleswar
}
    
func (c *三婆罗补罗SambalpurCounty) BManeswar摩泥湿伐罗() feud.Barony {
	return c.摩泥湿伐罗Maneswar
}
    
func (c *三婆罗补罗SambalpurCounty) BRajgangpur罗阇恒伽补罗() feud.Barony {
	return c.罗阇恒伽补罗Rajgangpur
}
    
func (c *三婆罗补罗SambalpurCounty) BRanipur罗尼补罗() feud.Barony {
	return c.罗尼补罗Ranipur
}
    
func (c *三婆罗补罗SambalpurCounty) BSamaleswari娑摩丽湿伐利() feud.Barony {
	return c.娑摩丽湿伐利Samaleswari
}
    
func (c *三婆罗补罗SambalpurCounty) BSambalpur三婆罗补罗() feud.Barony {
	return c.三婆罗补罗Sambalpur
}
    
func (c *三婆罗补罗SambalpurCounty) BSarugani娑楼迦尼() feud.Barony {
	return c.娑楼迦尼Sarugani
}
    
var CSambalpur三婆罗补罗 SambalpurCounty = &三婆罗补罗SambalpurCounty{}

func init() {
	f := CSambalpur三婆罗补罗.(*三婆罗补罗SambalpurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1249",
		Title:     "sambalpur",
		TitleName: "三婆罗补罗",
		TitleCode: "c_sambalpur",
		Baronies:  map[string]feud.Barony{},
	}

	f.毗摩丽湿伐罗Bimaleswar = BBimaleswar毗摩丽湿伐罗
	f.毗摩丽湿伐罗Bimaleswar.SetParent(f)
	
	f.摩泥湿伐罗Maneswar = BManeswar摩泥湿伐罗
	f.摩泥湿伐罗Maneswar.SetParent(f)
	
	f.罗阇恒伽补罗Rajgangpur = BRajgangpur罗阇恒伽补罗
	f.罗阇恒伽补罗Rajgangpur.SetParent(f)
	
	f.罗尼补罗Ranipur = BRanipur罗尼补罗
	f.罗尼补罗Ranipur.SetParent(f)
	
	f.娑摩丽湿伐利Samaleswari = BSamaleswari娑摩丽湿伐利
	f.娑摩丽湿伐利Samaleswari.SetParent(f)
	
	f.三婆罗补罗Sambalpur = BSambalpur三婆罗补罗
	f.三婆罗补罗Sambalpur.SetParent(f)
	
	f.娑楼迦尼Sarugani = BSarugani娑楼迦尼
	f.娑楼迦尼Sarugani.SetParent(f)
	
}

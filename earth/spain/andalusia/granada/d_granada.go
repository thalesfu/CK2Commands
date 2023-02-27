package granada

import (
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/granada/almeria"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/granada/granada"
	"github.com/thalesfu/CK2Commands/earth/spain/andalusia/granada/malaga"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GranadaDuke interface {
    feud.Duke
    CAlmeria阿尔梅里亚() 	almeria.AlmeriaCounty
    CGranada格拉纳达() 	granada.GranadaCounty
    CMalaga马拉加() 	malaga.MalagaCounty
}

type 格拉纳达GranadaDuke struct {
	feud.BaseDuke
	阿尔梅里亚Almeria 	almeria.AlmeriaCounty
	格拉纳达Granada 	granada.GranadaCounty
	马拉加Malaga 	malaga.MalagaCounty
}

func (d *格拉纳达GranadaDuke) CAlmeria阿尔梅里亚() almeria.AlmeriaCounty {
	return d.阿尔梅里亚Almeria
}
    
func (d *格拉纳达GranadaDuke) CGranada格拉纳达() granada.GranadaCounty {
	return d.格拉纳达Granada
}
    
func (d *格拉纳达GranadaDuke) CMalaga马拉加() malaga.MalagaCounty {
	return d.马拉加Malaga
}
    
var DGranada格拉纳达 GranadaDuke = &格拉纳达GranadaDuke{}

func init() {
	f := DGranada格拉纳达.(*格拉纳达GranadaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "granada",
		TitleName: "格拉纳达",
		TitleCode: "d_granada",
		Counties:  map[string]feud.County{},
	}

	f.阿尔梅里亚Almeria = almeria.CAlmeria阿尔梅里亚
	f.阿尔梅里亚Almeria.SetParent(f)
	
	f.格拉纳达Granada = granada.CGranada格拉纳达
	f.格拉纳达Granada.SetParent(f)
	
	f.马拉加Malaga = malaga.CMalaga马拉加
	f.马拉加Malaga.SetParent(f)
	
}

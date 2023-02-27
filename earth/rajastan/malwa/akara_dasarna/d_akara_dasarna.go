package akara_dasarna

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/akara_dasarna/candhoba"
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/akara_dasarna/chanderi"
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/akara_dasarna/sarangpur"
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/akara_dasarna/vidisa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Akara_dasarnaDuke interface {
    feud.Duke
    CCandhoba旃度婆() 	candhoba.CandhobaCounty
    CChanderi旃提梨() 	chanderi.ChanderiCounty
    CSarangpur萨浪伽补罗() 	sarangpur.SarangpurCounty
    CVidisa卑提写() 	vidisa.VidisaCounty
}

type 阿迦罗陀设那Akara_dasarnaDuke struct {
	feud.BaseDuke
	旃度婆Candhoba 	candhoba.CandhobaCounty
	旃提梨Chanderi 	chanderi.ChanderiCounty
	萨浪伽补罗Sarangpur 	sarangpur.SarangpurCounty
	卑提写Vidisa 	vidisa.VidisaCounty
}

func (d *阿迦罗陀设那Akara_dasarnaDuke) CCandhoba旃度婆() candhoba.CandhobaCounty {
	return d.旃度婆Candhoba
}
    
func (d *阿迦罗陀设那Akara_dasarnaDuke) CChanderi旃提梨() chanderi.ChanderiCounty {
	return d.旃提梨Chanderi
}
    
func (d *阿迦罗陀设那Akara_dasarnaDuke) CSarangpur萨浪伽补罗() sarangpur.SarangpurCounty {
	return d.萨浪伽补罗Sarangpur
}
    
func (d *阿迦罗陀设那Akara_dasarnaDuke) CVidisa卑提写() vidisa.VidisaCounty {
	return d.卑提写Vidisa
}
    
var DAkara_dasarna阿迦罗陀设那 Akara_dasarnaDuke = &阿迦罗陀设那Akara_dasarnaDuke{}

func init() {
	f := DAkara_dasarna阿迦罗陀设那.(*阿迦罗陀设那Akara_dasarnaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "akara_dasarna",
		TitleName: "阿迦罗陀设那",
		TitleCode: "d_akara_dasarna",
		Counties:  map[string]feud.County{},
	}

	f.旃度婆Candhoba = candhoba.CCandhoba旃度婆
	f.旃度婆Candhoba.SetParent(f)
	
	f.旃提梨Chanderi = chanderi.CChanderi旃提梨
	f.旃提梨Chanderi.SetParent(f)
	
	f.萨浪伽补罗Sarangpur = sarangpur.CSarangpur萨浪伽补罗
	f.萨浪伽补罗Sarangpur.SetParent(f)
	
	f.卑提写Vidisa = vidisa.CVidisa卑提写
	f.卑提写Vidisa.SetParent(f)
	
}

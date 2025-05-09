package culture

import (
	"fmt"
	"github.com/thalesfu/ck2nebula"
)

func BuildCulture() {
	cultureList := make([]*ck2nebula.Culture, 0)
	for _, culture := range CultureMap {
		cultureList = append(cultureList, culture)
	}

	cultureResult := ck2nebula.InsertCultures(ck2nebula.SPACE, cultureList...)

	if !cultureResult.Ok {
		fmt.Println(cultureResult.Err.Error())
	} else {
		fmt.Println("Insert cultures success")
	}

	cgs, _, cg_c_s := ck2nebula.GenerateCultureData("/Users/thalesfu/Windows/steam/steamapps/common/Crusader Kings II")

	cultureGroupResult := ck2nebula.InsertCultureGroups(ck2nebula.SPACE, cgs...)

	if !cultureGroupResult.Ok {
		fmt.Println(cultureGroupResult.Err.Error())
	} else {
		fmt.Println("Insert culture groups success")
	}

	for _, cg_c := range cg_c_s {
		cg_c.Culture = CultureMap[cg_c.Culture.Code]
	}

	cultrueGroupCultureEdgeResult := ck2nebula.InsertCultureGroup_Cultures(ck2nebula.SPACE, cg_c_s...)

	if !cultrueGroupCultureEdgeResult.Ok {
		fmt.Println(cultrueGroupCultureEdgeResult.Err.Error())
	} else {
		fmt.Println("Insert culture groups culture edge success")
	}

}

package concord

import (
	"fmt"
	"strings"
	"testing"
)

func TestSortedKeys(t *testing.T) {
	data := `2477 |                     en_us_net~ps~psa118.part3.text |    t possession for they give me joy i am determined to obey you
2478 |                      en_us_net~ps~psa118.v111.text |    t possession for they give me joy                              
2479 |                           en_us_net~ps~psa125.text |    aughed loudly and shouted for joy at that time the nations sa
2480 |                        en_us_net~ps~psa125.v2.text |    aughed loudly and shouted for joy at that time the nations sa
2481 |                       en_us_net~ps~psa125.v2a.text |    aughed loudly and shouted for joy                              
2482 |                        en_us_net~ps~psa125.v5.text |     as they plant will shout for joy when they reap the harvest
2483 |                        en_us_net~ps~psa125.v6.text |    ainly come in with a shout of joy carrying his sheaves of gra
2484 |                       en_us_net~ps~psa125.v6b.text |    ainly come in with a shout of joy carrying his sheaves of gra
2485 |                           en_us_net~ps~psa131.text |    our loyal followers shout for joy for the sake of david your 
2486 |                        en_us_net~ps~psa131.v9.text |    our loyal followers shout for joy                              
2487 |                           en_us_net~ps~psa136.text |    er whatever gives me the most joy remember o lord what the ed
2488 |                        en_us_net~ps~psa136.v6.text |    er whatever gives me the most joy                              
2489 |                           en_us_net~ps~psa149.text |    indication let them shout for joy upon their beds may they pr
2490 |                        en_us_net~ps~psa149.v5.text |    indication let them shout for joy upon their beds
2491 |                        en_us_net~ps~bode4.v16.text |    to scatter us they shout with joy as if they were plundering 
2492 |                        en_us_net~ps~bode5.v14.text |    ill rise up wake up and shout joyfully you who live in the gr
`
	lines := strings.Split(data,"\n")
	key := "joy"
	cmap := make(map[string]ConcordanceLine)
	for _, l := range lines {
		parts := strings.Split(l,"|")
		if len(parts) == 3 {
			var cl ConcordanceLine
			cl.ID = parts[1]
			moreParts := strings.Split(parts[2],key)
			cl.Left = moreParts[0]
			cl.Key = key
			cl.Right = moreParts[1]
			cmap[cl.ID] = cl
		}
	}
	fmt.Println("Sort ID")
	for _, res := range SortedKeys(cmap, SortId) {
		fmt.Println(cmap[res].ID, cmap[res].Left, cmap[res].Key, cmap[res].Right)
	}
	fmt.Println("Sort Left")
	for _, res := range SortedKeys(cmap, SortLeft) {
		fmt.Println(cmap[res].ID, cmap[res].Left, cmap[res].Key, cmap[res].Right)
	}
	fmt.Println("Sort Right")
	for _, res := range SortedKeys(cmap, SortRight) {
		fmt.Println(cmap[res].ID, cmap[res].Left, cmap[res].Key, cmap[res].Right)
	}

}

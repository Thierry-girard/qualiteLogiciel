package bowling

import "fmt"

type Frame struct {
	firstThrow  int
	secondThrow int
}

func GetScore(game []Frame) (int, error) {
	score := 0
	onze := 7
	douze := 2
    if len(game) == 0  || game == nil {
        return 0, fmt.Errorf("Game is empty")
    }
	if len(game) < 10{
		return 0, fmt.Errorf("Game Size < 10")
	}
	if len(game) > 10{
		return 0, fmt.Errorf("Game Size > 10")
	}		
	if game[9].firstThrow == 10{
		fmt.Scan(&onze)
        fmt.Scan(&douze)
	}else{
		if game[9].firstThrow + game[9].secondThrow == 10{
			fmt.Scan(&onze)
		}
	}

	for i := 0; i < len(game); i++ {
		if game[i].firstThrow < 0 || game[i].secondThrow < 0{
			return 0, fmt.Errorf("Negative value")
		}
		if game[i].firstThrow + game[i].secondThrow > 10{
			return 0, fmt.Errorf("Value Error : Frame > 10")
		}
		if game[i].firstThrow == 10 {   //Calcul des strikes 
			if i < 9 && game[i+1].firstThrow < 10 { 
				score = score + game[i].firstThrow + game[i+1].firstThrow + game[i+1].secondThrow //cas d'un simple strike 
			}
			if i < 8 && game[i+1].firstThrow == 10 { 
				score = score + game[i].firstThrow + game[i+1].firstThrow +  game[i+2].firstThrow // cas de deux strike consécutif
			}
			if i == 9 { 
				score = score + game[9].firstThrow + onze + douze //cas d'un strike à dizième frame
			}
			if i < 8 && game[9].firstThrow == 10 { 
				score = score + game[8].firstThrow + game[9].firstThrow + onze // cas d'un strike à la neuvième et dizième frame
			}		
		}else if game[i].firstThrow + game[i].secondThrow == 10 {   //Calcul des spares
				if i < 9 { 
					score = score + game[i].firstThrow  + game[i].secondThrow + game[i+1].firstThrow
				} else if i == 9 { 
					score = score + game[9].firstThrow + game[9].secondThrow + onze //cas d'un spare à la dizième frame
				}		
		}else{
        	score = score + game[i].firstThrow + game[i].secondThrow	
        }       
        if(onze > 10 || douze > 10){
            return 0, fmt.Errorf("Value Error : Frame > 10")
        }
        if(onze < 0 || douze < 0){
            return 0, fmt.Errorf("Negative value")
        }
	}
	return score, nil
}
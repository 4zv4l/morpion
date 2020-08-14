package main

import (
	"fmt"
)

func depart(tab [3][3]string) [3][3]string {
	tab[0][0] = "1"
	tab[0][1] = "2"
	tab[0][2] = "3"
	tab[1][0] = "4"
	tab[1][1] = "5"
	tab[1][2] = "6"
	tab[2][0] = "7"
	tab[2][1] = "8"
	tab[2][2] = "9"
	return tab
}

func menu() {
	println("MENU : ")
	println("Bienvenue dans ce jeu de morpion")
	println("Il faut alligner 3 X ou 3 O pour gagner")
	println("Bonne chance")
}

func affichetab(tab [3][3]string) {
	fmt.Println(tab[:][0])
	fmt.Println(tab[:][1])
	fmt.Println(tab[:][2])
}

func changetabX(tab [3][3]string, j1 string) [3][3]string {
	var choix int
	for choix != -33 {
		print("choix pour ",j1,": ")
		fmt.Scan(&choix)
		switch choix {
		case 1:
			if tab[0][0] != "X" && tab[0][0] != "O" {
				tab[0][0] = "X"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 2:
			if tab[0][1] != "X" && tab[0][1] != "O" {
				tab[0][1] = "X"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 3:
			if tab[0][2] != "X" && tab[0][2] != "O" {
				tab[0][2] = "X"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 4:
			if tab[1][0] != "X" && tab[1][0] != "O" {
				tab[1][0] = "X"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 5:
			if tab[1][1] != "X" && tab[1][1] != "O" {
				tab[1][1] = "X"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 6:
			if tab[1][2] != "X" && tab[1][2] != "O" {
				tab[1][2] = "X"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 7:
			if tab[2][0] != "X" && tab[2][0] != "O" {
				tab[2][0] = "X"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 8:
			if tab[2][1] != "X" && tab[2][1] != "O" {
				tab[2][1] = "X"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 9:
			if tab[2][2] != "X" && tab[2][2] != "O" {
				tab[2][2] = "X"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		default:
			println("vous n'avez pas taper un nombre valide !")
		}
		//fmt.Println(tab[:][0])
		//fmt.Println(tab[:][1])
		//fmt.Println(tab[:][2])
	}
	return tab
}

func changetab0(tab [3][3]string, j2 string) [3][3]string {
	var choix int
	for choix != -33 {
		print("choix pour ",j2,": ")
		fmt.Scan(&choix)
		switch choix {
		case 1:
			if tab[0][0] != "X" && tab[0][0] != "O" {
				tab[0][0] = "O"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 2:
			if tab[0][1] != "X" && tab[0][1] != "O" {
				tab[0][1] = "O"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 3:
			if tab[0][2] != "X" && tab[0][2] != "O" {
				tab[0][2] = "O"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 4:
			if tab[1][0] != "X" && tab[1][0] != "O" {
				tab[1][0] = "O"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 5:
			if tab[1][1] != "X" && tab[1][1] != "O" {
				tab[1][1] = "O"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 6:
			if tab[1][2] != "X" && tab[1][2] != "O" {
				tab[1][2] = "O"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 7:
			if tab[2][0] != "X" && tab[2][0] != "O" {
				tab[2][0] = "O"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 8:
			if tab[2][1] != "X" && tab[2][1] != "O" {
				tab[2][1] = "O"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		case 9:
			if tab[2][2] != "X" && tab[2][2] != "O" {
				tab[2][2] = "O"
				choix = -33
			} else {
				println("case deja utilisee")
			}
		default:
			print("vous n'avez pas taper un nombre valide !")
		}
		//fmt.Println(tab[:][0])
		//fmt.Println(tab[:][1])
		//fmt.Println(tab[:][2])
	}
	return tab
}

func verification(tab [3][3]string) int {
	var point int
	//verif horizontale X
	if tab[0][0] == "X" && tab[0][1] == "X" && tab[0][2] == "X" {
		point = 1
	} else if tab[1][0] == "X" && tab[1][1] == "X" && tab[1][2] == "X" {
		point = 1
	} else if tab[2][0] == "X" && tab[2][1] == "X" && tab[2][2] == "X" {
		point = 1
	} else if tab[0][0] == "O" && tab[0][1] == "O" && tab[0][2] == "O" {
		point = 2
	} else if tab[1][0] == "O" && tab[1][1] == "O" && tab[1][2] == "O" {
		point = 2
	} else if tab[2][0] == "O" && tab[2][1] == "O" && tab[2][2] == "O" {
		point = 2
	}
	//verif vertical
	if tab[0][0] == "X" && tab[1][0] == "X" && tab[2][0] == "X" {
		point = 1
	} else if tab[0][1] == "X" && tab[1][1] == "X" && tab[2][1] == "X" {
		point = 1
	} else if tab[0][1] == "X" && tab[1][1] == "X" && tab[2][1] == "X" {
		point = 1
	} else if tab[0][0] == "O" && tab[1][0] == "O" && tab[2][0] == "O" {
		point = 2
	} else if tab[0][1] == "O" && tab[1][1] == "O" && tab[2][1] == "O" {
		point = 2
	} else if tab[0][2] == "O" && tab[1][2] == "O" && tab[2][2] == "O" {
		point = 2
	}
	//verif horizontal
	if tab[0][0] == "X" && tab[1][1] == "X" && tab[2][2] == "X" {
		point = 1
	} else if tab[0][2] == "X" && tab[1][1] == "X" && tab[2][0] == "X" {
		point = 1
	} else if tab[0][0] == "O" && tab[1][1] == "O" && tab[2][2] == "O" {
		point = 2
	} else if tab[0][2] == "O" && tab[1][1] == "O" && tab[2][0] == "O" {
		point = 2
	}
	return point
}

func main() {
	var j1,j2 string
	print("j1 pour les X : ")
	fmt.Scanln(&j1)
	print("j2 pour les O : ")
	fmt.Scanln(&j2)
	menu()
	var reussi int = 0
	var plateau [3][3]string
	plateau = depart(plateau)
	//affiche le tableau
	for reussi != 1 {
		affichetab(plateau)
		plateau = changetabX(plateau, j1)
		if verification(plateau) == 1 {
			println(j1,"gagnant")
			affichetab(plateau)
			break
		}
		affichetab(plateau)
		plateau = changetab0(plateau, j2)
		if verification(plateau) == 2 {
			println(j2,"gagnant")
			affichetab(plateau)
			break
		}
	}
}

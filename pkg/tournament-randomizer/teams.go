package main

type GodTear struct {
	Teams []Teams
}

type Teams []string

func NewGodTear(teams []Teams) GodTear {
	return GodTear{Teams: teams}
}

/*
Manchester City, Liverpool, Manchester United, Arsenal

Bayern Munich, Borussia Dortmund,

Real Madrid, Barcelona, Athletic Madrid

Napoli, Milan, Inter



*/

func ReceiveGodTearTeams() []Teams {
	england := Teams{"England:Manchester-City,", "England:Liverpool", "England:Manchester-United", "England:Chelsea"}
	germany := Teams{"Germany:Bayern-Munich"}
	italy := Teams{"Italy:Juventus(Piemonte Calcio)", "Italy:Inter", "Italy:Napoli"}
	spain := Teams{"Spain:Real-Madrid", "Spain:Barcelona", "Spain:Athletic-Madrid"}

	return []Teams{
		england,
		germany,
		italy,
		spain,
	}
}

type SecondTier struct {
	Teams []Teams
}

func NewSecondTier(teams []Teams) SecondTier {
	return SecondTier{Teams: teams}
}

func ReceiveSecondTierTeams() []Teams {
	england := Teams{"England:Tottenham", "England:Leicester-City", "England:West-Ham"}
	germany := Teams{"Germany:Borussia-Dortmund", "Germany:Leverkusen", "Germany:RB-Leipzig"}
	italy := Teams{"Italy:Lazio(Latuim)", "Italy:Roma"}
	spain := Teams{"Spain:Athletic-Club", "Spain:Real-Sociedad", "Spain:Villarreal", "Spain:Sevilla"}

	others := Teams{"Netherlands:Ajax", "Portugal:Benfica", "France:Monaco", "France:Marcel-OM"}

	return []Teams{
		england,
		germany,
		italy,
		spain,
		others,
	}
}

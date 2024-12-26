package main

type Teams []string

type GodTear struct {
	Teams []Teams
}

func (gt *GodTear) ReceiveGodTearTeams() []Teams {
	england := Teams{"England:Manchester-City", "England:Liverpool", "England:Arsenal", "England:Man United"}
	germany := Teams{"Germany:Bayern-Munich"}
	france := Teams{"France:PSG"}
	spain := Teams{"Spain:Real-Madrid", "Spain:Barcelona", "Spain:Athletic-Madrid"}

	return []Teams{
		england,
		germany,
		spain,
		france,
	}
}

func NewGodTear(teams []Teams) GodTear {
	godTear := GodTear{}

	godTear.Teams = godTear.ReceiveGodTearTeams()

	return godTear
}

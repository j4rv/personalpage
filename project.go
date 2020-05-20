package main

type project struct {
	Name  string
	URL   string
	Image string
}

var projects = []project{
	{
		Name:  "Cards against humanity!",
		URL:   "https://j4rvcah.herokuapp.com",
		Image: "cah_background.png",
	},
	{
		Name:  "PC Remote Controller",
		URL:   "https://github.com/j4rv/remotecontrol",
		Image: "remotecontrol_background.png",
	},
	{
		Name:  "Risk Dice Roller",
		URL:   "https://j4rv.github.io/risk-dice-roller",
		Image: "riskdiceroller_background.png",
	},
	{
		Name:  "Nonograms~",
		URL:   "https://j4rv.github.io/nonograms",
		Image: "nonograms_background.png",
	},
}

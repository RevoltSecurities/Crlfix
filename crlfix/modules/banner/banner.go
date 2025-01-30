package banner

import (
	"fmt"
	"math/rand"

	"github.com/RevoltSecurities/Crlfix/crlfix/modules/logger"
	"github.com/common-nighthawk/go-figure"
	"github.com/logrusorgru/aurora/v4"
)

func Randomchoice(choices []string) string {
	choosed := rand.Intn(len(choices))
	return choices[choosed]
}

func BannerGenerator(banner_name string) aurora.Value {
	choices := []string{"big", "ogre", "shadow", "script", "graffiti", "slant"}
	colors := []string{"green", "cyan", "blue", "white", "magenta"}
	banner := figure.NewFigure(banner_name, Randomchoice(choices), true)
	banners := logger.Bannerizer(fmt.Sprintf(`%s`, banner.String()), Randomchoice(colors))
	return banners
}

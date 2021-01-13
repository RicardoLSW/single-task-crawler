package parser

import (
	"regexp"

	"single-task-crawler/model"

	"single-task-crawler/engine"
)

var userInfoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)

func parseProfile(contents []byte) engine.ParseResult {
	match := userInfoRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	profile := model.Profile{}
	for i, user := range match {
		switch i {
		case 0:
			profile.Marriage = string(user[1])
			break
		case 1:
			profile.Age = string(user[1])
			break
		case 2:
			profile.Xinzuo = string(user[1])
			break
		case 3:
			profile.Height = string(user[1])
			break
		case 4:
			break
		case 5:
			profile.Income = string(user[1])
			break
		case 6:
			profile.Education = string(user[1])
			break

		}
		result.Items = append(result.Items, profile)
		result.Requests = append(result.Requests, engine.Request{Url: "", ParserFunc: engine.NilParser})
	}
	return result
}

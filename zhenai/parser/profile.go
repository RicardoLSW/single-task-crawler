package parser

import (
	"regexp"

	"single-task-crawler/model"

	"single-task-crawler/engine"
)

var userInfoRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)

func parseProfile(contents []byte, name string) engine.ParseResult {
	match := userInfoRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	profile := model.Profile{}
	profile.Name = name
	profile.Marriage = string(match[0][1])
	profile.Age = string(match[1][1])
	profile.Xinzuo = string(match[2][1])
	profile.Height = string(match[3][1])
	profile.Income = string(match[5][1])
	profile.Education = string(match[6][1])
	result.Items = append(result.Items, profile)
	return result
}

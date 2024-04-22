package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name   string
	win    int
	loss   int
	draw   int
	match  int
	points int
}

var teams map[string]*team

func processLine(line string) error {
	line = strings.Trim(line, " ")
	if line == "" || line[0] == '#' {
		return nil
	}
	words := strings.Split(line, ";")
	if len(words) < 3 || !(words[2] == "win" || words[2] == "loss" || words[2] == "draw") {
		return errors.New("error")
	}
	var team1 *team
	var team2 *team
	var ok bool
	if team1, ok = teams[words[0]]; !ok {
		team1 = &team{name: words[0]}
		teams[words[0]] = team1
	}
	if team2, ok = teams[words[1]]; !ok {
		team2 = &team{name: words[1]}
		teams[words[1]] = team2
	}
	if words[2] == "win" {
		team1.win++
		team2.loss++
		team1.match++
		team2.match++
		team1.points += 3
	} else if words[2] == "loss" {
		team2.win++
		team1.loss++
		team2.match++
		team1.match++
		team2.points += 3
	} else {
		team2.draw++
		team1.draw++
		team2.match++
		team1.match++
		team2.points++
		team1.points++
	}
	return nil
}

func read(reader io.Reader) string {
	var b strings.Builder
	p := make([]byte, 100)
	for {
		l, err := reader.Read(p)
		if l == 0 && err != nil {
			break
		}
		for i := 0; i < l; i++ {
			b.WriteByte(p[i])
		}
	}
	return b.String()
}

// Tally results of a small football competition.
func Tally(reader io.Reader, buffer io.Writer) error {
	var teamsArr []*team
	s := read(reader)
	lines := strings.Split(s, "\n")
	teams = make(map[string]*team)
	for _, line := range lines {
		err := processLine(line)
		if err != nil {
			return err
		}
	}
	l := 0
	for _, t := range teams {
		if l < len(t.name) {
			l = len(t.name)
		}
		teamsArr = append(teamsArr, t)
	}
	sort.Slice(teamsArr, func(i, j int) bool {
		if teamsArr[i].points == teamsArr[j].points {
			return teamsArr[i].name < teamsArr[j].name
		}
		return teamsArr[i].points > teamsArr[j].points
	})

	buffer.Write([]byte(fmt.Sprintf("%-*s| MP |  W |  D |  L |  P\n", l+8, "Team")))
	for _, team := range teamsArr {
		buffer.Write([]byte(fmt.Sprintf("%-*v|  %v |  %v |  %v |  %v |  %v\n", l+8, team.name, team.match, team.win, team.draw, team.loss, team.points)))
	}
	return nil
}

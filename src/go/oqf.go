package oqf

import (
	"fmt"
	"strings"
	"strconv"
)

type Question struct {
	Title    string
	Choices  []string
	Answers  []bool
	Points   int
	Comment  string
	Time     int
	Required bool
}

var DefaultQuestion = Question{
	Required: true,
}

func Parse(raw []byte) []Question {
	data := make([]Question, 0)
	q := -1
	a := 0
	t := ""
	flag := ""
	for _, chr := range raw {
		if len(flag) == 0 {
			if chr == ':' {
				data = append(data, DefaultQuestion)
				q += 1
				flag = "T"
				continue
			}
			continue
		}
		switch flag[0] {
		case 'T':
			if chr != 0x0A {
				data[q].Title += string(chr)
			} else {
				flag = "QD"
			}
		case 'Q':
			switch flag[1] {
			case 'D':
				switch chr {
				case '=':
					data[q].Choices = append(data[q].Choices, "")
					data[q].Answers = append(data[q].Answers, true)
					flag = "QN"
				case '!':
					data[q].Choices = append(data[q].Choices, "")
					data[q].Answers = append(data[q].Answers, false)
					flag = "QN"
				case '#':
					data[q].Comment = ""
					flag = "QK"
				case ';':
					a = 0
					q += 1
					flag = ""
				case '?':
					data[q].Required = false
				case '"':
					flag = "QT"
				case '/':
					flag = "QP"
				}
			case 'N':
				switch chr {
				case 0x0A:
					a += 1
					flag = "QD"
				case ';':
					a = 0
					flag = ""
				default:
					data[q].Choices[a] += string(chr)
				}
			case 'K':
				if chr == 0x0A {
					flag = "QD"
				} else {
					data[q].Comment += string(chr)
				}
			case 'T':
				if chr == 0x0A {
					time, err := strconv.Atoi(t)
					if err != nil {
						time = 0
					}
					data[q].Time = time
					flag = "QD"
					t = ""
				} else {
					t += string(chr)
				}
			case 'P':
				if chr == 0x0A {
					pts, err := strconv.Atoi(t)
					if err != nil {
						pts = 0
					}
					data[q].Points = pts
					flag = "QD"
					t = ""
				} else {
					t += string(chr)
				}
			}
		}
	}
	return data
}

func (q Question) ToBytes() []byte {
	var s strings.Builder
	s.Write(fmt.Appendf([]byte{}, ":%s", q.Title))
	if q.Comment != "" {
		s.Write(fmt.Appendf([]byte{}, "\n#%s", q.Comment))
	}
	if q.Time != 0 {
		s.Write(fmt.Appendf([]byte{}, "\n\"%d", q.Time))
	}
	if q.Points != 0 {
		s.Write(fmt.Appendf([]byte{}, "\n/%d", q.Points))
	}
	if !q.Required {
		s.Write([]byte{0x0A,'?'})
	}
	for i := range(q.Choices) {
		if q.Answers[i] {
			s.Write(fmt.Appendf([]byte{}, "\n=%s", q.Choices[i]))
		} else {
			s.Write(fmt.Appendf([]byte{}, "\n!%s", q.Choices[i]))
		}
	}
	s.Write([]byte{';',0x0A,0x0A})
	return []byte(s.String())
}

func Compose(q []Question) []byte {
	var b []byte
	for _, i := range(q) {
		b = append(b, i.ToBytes()...)
	}
	return b[0:len(b)-1]
}

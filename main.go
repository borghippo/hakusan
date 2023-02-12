package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/borghippo/hakusan/generate"
	tea "github.com/charmbracelet/bubbletea"
)

const numOptions = 5

var ans, choices = generate.GenerateHiragana(numOptions)

type model struct {
	cursor int
	choice generate.HiraganaCharacter
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString(ans.Hiragana + "\n\n")

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i].Latin)
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func main() {
	p := tea.NewProgram(model{})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.choice.Latin != "" {
		if m.choice.Latin == ans.Latin {
			fmt.Printf("\n---\nCorrect! %s is %s\n", m.choice.Hiragana, m.choice.Latin)
		} else {
			fmt.Printf("\n---\nWrong... %s is %s\n", ans.Hiragana, ans.Latin)
		}
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/borghippo/hakusan/generate"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const numOptions = 5

var (
	ans, choices   = generate.GenerateHiragana(numOptions)
	choiceStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	incorrectStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#fb4934")).Bold(true)
	correctStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#b8bb26")).Bold(true)
)

type model struct {
	cursor int
	choice generate.HiraganaCharacter
}

func checkbox(label string, checked bool) string {
	if checked {
		return fmt.Sprintf("%s", choiceStyle.Render("[x] "+label))
	}
	return fmt.Sprintf("[ ] %s", label)
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
	s := fmt.Sprintf("%s\n\n", ans.Hiragana)

	for i := 0; i < len(choices); i++ {
		if m.cursor == i {
			s += checkbox(choices[i].Latin, true)
		} else {
			s += checkbox(choices[i].Latin, false)
		}
		s += "\n"
	}
	s += "\n(press enter to choose)\n"

	return fmt.Sprintf(s)
}

func main() {
	p := tea.NewProgram(model{})

	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if m, ok := m.(model); ok && m.choice.Latin != "" {
		if m.choice.Latin == ans.Latin {
			fmt.Printf("\n---\n%s %s is %s\n", correctStyle.Render("Correct!"), m.choice.Hiragana, m.choice.Latin)
		} else {
			fmt.Printf("\n---\n%s %s is %s\n", incorrectStyle.Render("Wrong..."), ans.Hiragana, ans.Latin)
		}
	}
}

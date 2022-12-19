package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices []string
	cursor int
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		choices: []string{"Buy carrots", "Buy celery", "Buy coffee"},

		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is a key pressed?
	case tea.KeyMsg:

		// Which key?
		switch msg.String() {

		// The keys to exit the program
		case "ctrl+c", "q":
			return m, tea.Quit

		// Move the cursor up
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}

		// Move the cursor down
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The enter key toggles the selected state
		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	// The header
	s := "What do I need from the grocery store?\n\n"

	for i, choice := range m.choices {

		// Is the cursor pointing at the choice?
		cursor := " " // No cursor
		if m.cursor == i {
			cursor = ">"
		}

		// Is the choice selected?
		checked := " " // Not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // Selected
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// Footer
	s += "\nPress q to quit\n"

	// Render the UI
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}